package utils

import(
    "os"
	"path/filepath"
    "bufio"
    "encoding/json"
    
    "golang.org/x/net/context"
    
    "github.com/docker/libcompose/docker"
    "github.com/docker/libcompose/project"
    "github.com/docker/libcompose/logger"
	"github.com/docker/engine-api/types"
)

//ClientFactory the docker client factory
var ClientFactory docker.ClientFactory

//InitDocker init the docker client factory
func InitDocker() error{
    var err error
    
    /*
        If the SWARMUI_SWARM_REPLICATION get the leader host with the key "/docker/swarm/leader" from the KV store
        Else use default env vars to connect to the docker daemon (DOCKER_HOST, ...)
    */
    if os.Getenv("SWARMUI_SWARM_REPLICATION") == "true" {
        var swarmLearder string
        swarmLearder, err = GetSwarmLeader()
        
        if err != nil {
            return err
        }
        
        opts := docker.ClientOpts{
            Host: swarmLearder,
            TLSVerify: os.Getenv("DOCKER_TLS_VERIFY") == "1",
        }
        
        if dockerCertPath := os.Getenv("DOCKER_CERT_PATH"); dockerCertPath != "" {
             opts.TLSOptions.CAFile = filepath.Join(dockerCertPath, "ca.pem")
             opts.TLSOptions.CertFile = filepath.Join(dockerCertPath, "cert.pem")
             opts.TLSOptions.KeyFile = filepath.Join(dockerCertPath, "key.pem")
             opts.TLSOptions.InsecureSkipVerify = os.Getenv("DOCKER_TLS_VERIFY") == ""
        }
        
        ClientFactory, err = docker.NewDefaultClientFactory(opts)    
    } else {
        ClientFactory, err = docker.NewDefaultClientFactory(docker.ClientOpts{})
    }
    return err
}

//GetProject get the docker compose project
func GetProject(composeData []byte, appID string) (*project.Project, error){
    return docker.NewProject(&docker.Context{
			Context: project.Context{
				ComposeBytes: [][]byte{composeData},
				ProjectName:  appID,
                LoggerFactory: &logger.NullLogger{},
			},
            ClientFactory: ClientFactory,
		})
}

//Stat container stat
type Stat struct {
    Stat                types.Stats
    PreviousStat        *types.Stats
    ContainerID         string
}

//GetContainersStats get the stats of a container
func GetContainersStats(containerID string, quit <-chan struct{}, statChannel chan<- Stat) ( error ){
    cli := ClientFactory.Create(nil)
    ctx, cancelContext := context.WithCancel(context.Background())
    reader, err := cli.ContainerStats(ctx, containerID, true)
    
    var previousStat *types.Stats
    var stat types.Stats
    
    if err == nil {
        scan := bufio.NewReader(reader)
        go func(){
            for {
                select {
                    case <- quit:
                        cancelContext()
                        return
                    default:
                        line, _, _ := scan.ReadLine()
                        json.Unmarshal(line, &stat) 
                        statChannel <- Stat{
                            Stat : stat,
                            PreviousStat : previousStat,
                            ContainerID : containerID,
                        }
                        previousStat = new(types.Stats)
                        *previousStat = stat
                }
            }
        }()
        return nil
    }
    return err
}