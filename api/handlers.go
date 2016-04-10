package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
    
    "github.com/jfdamy/swarmui/utils"
    "github.com/jfdamy/swarmui/autoscaling"
    
	"github.com/docker/libcompose/project"
	"github.com/gorilla/mux"
	log "github.com/Sirupsen/logrus"
)

//ProjectList list all projects
func ProjectList(w http.ResponseWriter, r *http.Request) {
	projects, err := utils.GetListComposeProject()

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(projects); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectDefinition get the compose definition
func ProjectDefinition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["appId"]

	composeData, err := utils.GetComposeProject(appID)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(struct{ Compose string  `json:"compose"`}{string(composeData)}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectCreate create a project
func ProjectCreate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["appId"]

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err := r.Body.Close(); err == nil {
			err = utils.SetComposeProject(appID, body)
			if err == nil {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
					panic(err)
				}
				return
			}
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectShow show info of the project
func ProjectShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	appID := vars["appId"]

	composeData, err := utils.GetComposeProject(appID)

	if err == nil {
        proj, err := utils.GetProject(composeData, appID);

		if err == nil {
			var projInfo projectInfo
			var svcInfo []serviceInfo

			projInfo.ProjectID = appID

			for _, name := range proj.Configs.Keys() {
                config, _ := proj.Configs.Get(name)
				service, _ := proj.CreateService(name)
				containers, _ := service.Containers()
			    var containersInfo []containerInfo
                
				for _, cont := range containers {
					isRunning, _ := cont.IsRunning()

					var ports []port
					for _, p := range config.Ports {
						if !strings.HasSuffix(p, "/tcp") && !strings.HasSuffix(p, "/udp") {
							p = p + "/tcp"
						}
						portHost, _ := cont.Port(p)
						ports = append(ports, port{
							PortHost: portHost,
							PortCont: p,
						})
					}

					containersInfo = append(containersInfo, containerInfo{
						ContainerName: cont.Name(),
						Port:          ports,
						IsRunning:     isRunning,
					})
				}
				serviceTemp := serviceInfo{
					ServiceName: name,
					Containers:  containersInfo,
					Config:      config,
				}
				svcInfo = append(svcInfo, serviceTemp)
			}
			projInfo.Services = svcInfo
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(projInfo); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectUp up the project
func ProjectUp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	var body []byte
	var svcs services
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err = r.Body.Close(); err == nil {
			_ = json.Unmarshal(body, &svcs)
		}
	}
	err = nil

	composeData, err := utils.GetComposeProject(appID)

	if err == nil {
		proj, err := utils.GetProject(composeData, appID)

		if err == nil {
			log.Info("Project : ", appID, " up services : ", svcs.ServicesName)
			if len(svcs.ServicesName) == 0 {
				proj.Up(svcs.ServicesName...)
			} else {
				proj.Up()
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectStop stop the project
func ProjectStop(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	var body []byte
	var svcs services
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err = r.Body.Close(); err == nil {
			_ = json.Unmarshal(body, &svcs)
		}
	}
	err = nil

	composeData, err := utils.GetComposeProject(appID)

	if err == nil {
		proj, err := utils.GetProject(composeData, appID)

		if err == nil {
			log.Info("Project : ", appID, " stop services : ", svcs.ServicesName)
			if len(svcs.ServicesName) != 0 {
				proj.Down(svcs.ServicesName...)
			} else {
				proj.Down()
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectKill kill the project
func ProjectKill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	var body []byte
	var svcs services
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err = r.Body.Close(); err == nil {
			_ = json.Unmarshal(body, &svcs)
		}
	}
	err = nil

	composeData, err := utils.GetComposeProject(appID)
	if err == nil {
		proj, err := utils.GetProject(composeData, appID)

		if err == nil {
			log.Info("Project : ", appID, " kill services : ", svcs.ServicesName)
			if len(svcs.ServicesName) != 0 {
				proj.Kill(svcs.ServicesName...)
			} else {
				proj.Kill()
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectRemove remove container(s) of the project
func ProjectRemove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	var body []byte
	var svcs services
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err = r.Body.Close(); err == nil {
			_ = json.Unmarshal(body, &svcs)
		}
	}
	err = nil

	composeData, err := utils.GetComposeProject(appID)
	if err == nil {
		proj, err := utils.GetProject(composeData, appID)

		if err == nil {
			log.Info("Project : ", appID, " remove container of services : ", svcs.ServicesName)
			if len(svcs.ServicesName) == 0 {
				proj.Delete(svcs.ServicesName...)
			} else {
				proj.Delete()
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
				panic(err)
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

//ProjectDelete delete the project
func ProjectDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	composeData, err := utils.GetComposeProject(appID)
	if err == nil {
		proj, err := utils.GetProject(composeData, appID)

		if err == nil {
			log.Info("Delete Project : ", appID)
            proj.Down()
		    proj.Delete()
            err = utils.RemoveComposeProject(appID)
               
            if err == nil {
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusOK)
                if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
                    panic(err)
                }
                return
            }
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

// ServiceScale scale services
func ServiceScale(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	composeData, err := utils.GetComposeProject(appID)
	if err == nil {
		var body []byte
		body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err == nil {
			if err = r.Body.Close(); err == nil {
				var scaleServices []scaleService
				if err = json.Unmarshal(body, &scaleServices); err == nil {
					var proj *project.Project
					proj, err = utils.GetProject(composeData, appID)

					if err == nil {
						log.Info("Project : ", appID, " scale services : ", scaleServices)
						for _, scale := range scaleServices {
							var service project.Service
							service, err = proj.CreateService(scale.ServiceName)
							if err == nil {
								service.Scale(scale.Number)
							}
						}

						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)
						if err = json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
							panic(err)
						}
						return
					}
				}
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

// ServiceAutoScaling set the autoscaling
func ServiceAutoScaling(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]
	var body []byte
    body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err == nil {
        if err = r.Body.Close(); err == nil {
                var scaleService autoscaling.ScalingService
                var scaleServices []autoscaling.ScalingService
				if err = json.Unmarshal(body, &scaleService); err == nil {
                    log.WithField("scaleService", scaleService).Info("Project : ", appID, " autoscaling")
                    scaleServicesJSON, _ := utils.GetAutoscalingProject(appID)
                    if scaleServicesJSON != nil && len(scaleServicesJSON.Value) != 0 {
                        if err = json.Unmarshal(scaleServicesJSON.Value, &scaleServices); err != nil {
                            log.Error(err)
                        } else {
                            found := -1
                            for i, value := range scaleServices{
                                if value.ServiceName == scaleService.ServiceName{
                                    found = i
                                }
                            }
                            if found != -1 {
                                scaleServices = append(scaleServices[:found], scaleServices[found+1:]...)
                            }
                        }
                    }
                    
                    scaleServices = append(scaleServices, scaleService)
                    autoscalingProjectJSON, _ := json.Marshal(scaleServices)
                    err = utils.SetAutoscalingProject(appID, autoscalingProjectJSON)
                    
                    if err == nil {
                        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                        w.WriteHeader(http.StatusOK)
                        if err = json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
                            panic(err)
                        }
                        return
                    }
                }
        }
    }
    
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
		panic(err)
	}
}

// ServiceRemoveAutoScaling remove the autoscaling
func ServiceRemoveAutoScaling(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]
    
    err = utils.RemoveAutoscalingProject(appID)
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    if err == nil{
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "OK"}); err != nil {
            panic(err)
        }
    } else {
        w.WriteHeader(http.StatusNotFound)
        if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: err.Error()}); err != nil {
            panic(err)
        }
    }
}