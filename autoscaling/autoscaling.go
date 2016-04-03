package autoscaling
import (
    "time"
    "strings"
    "os"
    "os/signal"
    "strconv"
    "encoding/json"
    
	log "github.com/Sirupsen/logrus"
    "github.com/jfdamy/swarmui/utils"
)

//nodeScaling Manage a service with node scaling
func nodeScaling(service ScalingService, quit <-chan struct{}){
    log.Info("Node scaling for ", service.ServiceName)
    
    for {
        select {
            case <- quit:
                log.Debug("Scaling service ", service.ServiceName, " exited")
                return
            default:
                info, _ := utils.ClientFactory.Create(nil).Info()
                if info.DriverStatus[3][0][1:len(info.DriverStatus[3][0])] == "Nodes" {
                    scaleTo, err := strconv.Atoi(info.DriverStatus[3][1])
                    if err == nil {
                        if service.containerNumber != scaleTo {
                            log.Info("Scaling service ", service.ServiceName, " to ", scaleTo)
                            service.containerNumber = scaleTo
                            service.service.Scale(service.containerNumber)
                        }
                    } else {
                        log.Error(err)
                    }
                    
                } else {
                    log.Error("No nodes definition, not docker swarm ?")
                    <- quit
                    return
                }
                time.Sleep(50 * time.Millisecond)
        }
    }
}

//watchContainers Watch containers (for their stats)
func watchContainers(service ScalingService, mapWatchContainer map[string]chan struct{}, 
    statsChannel chan<- utils.Stat){
    containers, err := service.service.Containers()
   
    if err == nil {
        mapContainers := make(map[string]bool)
        for _, container := range containers{
            id,_ := container.ID()
            running ,_ := container.IsRunning()
            mapContainers[id]=true
            
            if running && mapWatchContainer[id] == nil {
                
                //Watch every running containers (not already being watched)
                log.Debug("Watch container id ", id, " to service ", service.ServiceName)
                quitChannel := make(chan struct{}, 2) 
                err = utils.GetContainersStats(id, quitChannel, statsChannel)
                if err != nil {
                     log.Error(err)
                } else {
                    mapWatchContainer[id] = quitChannel
                }
            }  else if !running && mapWatchContainer[id] != nil {
                
                //Unwatch every stopped containers
                mapWatchContainer[id] <- struct{}{} //mapContainer[id] channel has a size of 2 to not block here with GetContainersStats
                delete(mapWatchContainer, id)
                log.Debug("Unwatch container id ", id, " to service", service.ServiceName)  
            }
        }
        
        for key := range mapWatchContainer {
            
            //Unwatch every removed containers
            if !mapContainers[key] {
                mapWatchContainer[key] <- struct{}{} //mapContainer[id] channel has a size of 2 to not block here with GetContainersStats
                delete(mapWatchContainer, key)
                log.Debug("Unwatch container id ", key, " to service", service.ServiceName)  
            }
        }
    } else {
        log.Error(err)
    }
}

//autoScaling Manage a service with auto scaling (scaling by cpu)
func autoScaling(service ScalingService, quit <-chan struct{}){
    log.Info("Auto scaling for ", service.ServiceName)
   
    mapWatchContainer := make(map[string]chan struct{})
    mapCPUUsage := make(map[string]float64)
    statsChannel := make(chan utils.Stat)
    
    //Watch containers (for their stats)
    watchContainers(service, mapWatchContainer, statsChannel)
    
    for {
            select {
            case <- quit:
                log.Debug("Scaling service ", service.ServiceName, " exited")
                for _, channel := range mapWatchContainer{
                    channel <- struct{}{}
                }
                return
            case stats := <- statsChannel:
                if stats.PreviousStat != nil {
                    
                    //Compute the cpu usage of the container in percent
                    diffCPU := (float64)(stats.Stat.CPUStats.CPUUsage.TotalUsage - stats.PreviousStat.CPUStats.CPUUsage.TotalUsage)
                    diffCPUSystem := (float64)(stats.Stat.CPUStats.SystemUsage - stats.PreviousStat.CPUStats.SystemUsage)
                    cpuPercent := (diffCPU / diffCPUSystem)
                    if mapWatchContainer[stats.ContainerID] != nil {
                        mapCPUUsage[stats.ContainerID] = cpuPercent
                    }
                    
                    //Compute the global cpu usage of all containers
                    if len(mapCPUUsage) == service.containerNumber {
                        cpt := 0.0
                        globalCPUUsage := 0.0
                        for _, cpu := range mapCPUUsage {
                            globalCPUUsage += cpu
                            cpt++
                        }
                        globalCPUUsage = globalCPUUsage / cpt
                        if globalCPUUsage >= 0.50 {
                            service.service.Scale(service.containerNumber+1)
                            service.containerNumber = service.containerNumber+1
                            log.Info("Scaling up service ", service.ServiceName)
                        } else if globalCPUUsage <= 0.15 && service.containerNumber > 1{
                            service.service.Scale(service.containerNumber-1)
                            service.containerNumber = service.containerNumber-1
                            log.Info("Scaling down service ", service.ServiceName)
                            
                            //Clean the map
                            mapCPUUsage = make(map[string]float64)
                        }
                    }
                }
                
                //Update watched containers (for their stats)
                watchContainers(service, mapWatchContainer, statsChannel)
        }
    }
}

//manageServices Manage services
func manageServices(project string, services []byte, quit <-chan struct{}){
    var scalingServices []ScalingService;
    err := json.Unmarshal(services, &scalingServices)
    if err != nil {
        log.Error(err)
    }
    compose, err := utils.GetComposeProject(project)
    if err != nil {
        log.Error(err)
    }
    proj, err := utils.GetProject(compose, project)
    if err != nil {
        log.Error(err)
    }
       
    if err == nil {
        for _, scalingDef  := range scalingServices{
            svc, err := proj.CreateService(scalingDef.ServiceName)
            if err != nil {
                log.Error(err)
            } else {
                if svc != nil {
                    containers, _ := svc.Containers()
                    scalingDef.containerNumber = len(containers)
                    scalingDef.service = svc
                    if scalingDef.ScalingType == nodescaling {
                        go nodeScaling(scalingDef, quit)
                    } else if  scalingDef.ScalingType == autoscaling {
                        go autoScaling(scalingDef, quit)
                    } else {
                        log.Fatal(scalingDef.ScalingType, " scaling type for ",scalingDef.ServiceName," is not a scaling type")
                        return
                    }
                }
            }
        }
    }
    <- quit
}

//Run the autoscaling
func Run() {

	err := utils.InitStore()
    if err != nil {
        panic(err)
    }
    
    err = utils.InitDocker()
    if err != nil {
        panic(err)
    }
    
    stopChannel := make(<- chan struct{})
    signalChannel := make(chan os.Signal)
    stopChannels := make(map[string]chan struct{})
    watch, err := utils.WatchAutoscalingProjects(stopChannel)
    
	signal.Notify(signalChannel, os.Interrupt)
    
    
    if err == nil{
        log.Info("Autoscaling started")
        for {
            select {
                case changes := <- watch:
                    for _, project := range changes {
                        if project != nil {
                            id := strings.Split(project.Key, "/")[3]
                            if stopChannels[id] == nil {
                                ch := make(chan struct{})
                                stopChannels[id] = ch
                                go manageServices(id, project.Value, ch)
                            } else {
                                log.Info("Update auto scaling for ", project.Key)
                                stopChannels[id] <- struct{}{}
                                go manageServices(id, project.Value, stopChannels[id])
                            }
                        }
                    }
               case <- signalChannel:
                    log.Info("Autoscaling exiting")
                    for _, stop := range stopChannels {
                        stop <- struct{}{}
                    }
                    return
               case <- stopChannel:
                    log.Info("Autoscaling exiting")
                    for _, stop := range stopChannels {
                        stop <- struct{}{}
                    }
                    return
            }
        }
    }
    log.Error(err)
}
