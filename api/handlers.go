package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/project"
	"github.com/gorilla/mux"
)

func getProject(composeData []byte, appID string) (*project.Project, error){
    return docker.NewProject(&docker.Context{
			Context: project.Context{
				ComposeBytes: [][]byte{composeData},
				ProjectName:  appID,
			},
            ClientFactory: ClientFactory,
		})
}

//ProjectList list all projects
func ProjectList(w http.ResponseWriter, r *http.Request) {
	projects, err := GetListComposeProject()

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

	composeData, err := GetComposeProject(appID)

	if err == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(struct{ Compose string  `json:compose`}{string(composeData)}); err != nil {
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
			err = SetComposeProject(appID, body)
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

	composeData, err := GetComposeProject(appID)

	if err == nil {
        proj, err := getProject(composeData, appID);

		if err == nil {
			var projInfo projectInfo
			var svcInfo []serviceInfo

			projInfo.ProjectId = appID

			for name, config := range proj.Configs {
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

	composeData, err := GetComposeProject(appID)

	if err == nil {
		proj, err := getProject(composeData, appID)

		if err == nil {
			log.Println("services : ", svcs.ServicesName)
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

	composeData, err := GetComposeProject(appID)

	if err == nil {
		proj, err := getProject(composeData, appID)

		if err == nil {
			log.Println("services : ", svcs.ServicesName)
			if len(svcs.ServicesName) == 0 {
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

	composeData, err := GetComposeProject(appID)
	if err == nil {
		proj, err := getProject(composeData, appID)

		if err == nil {
			log.Println("services : ", svcs.ServicesName)
			if len(svcs.ServicesName) == 0 {
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

	composeData, err := GetComposeProject(appID)
	if err == nil {
		proj, err := getProject(composeData, appID)

		if err == nil {
			log.Println("services : ", svcs.ServicesName)
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

	var body []byte
	var svcs services
	body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err == nil {
		if err = r.Body.Close(); err == nil {
			_ = json.Unmarshal(body, &svcs)
		}
	}
	err = nil

	composeData, err := GetComposeProject(appID)
	if err == nil {
		proj, err := getProject(composeData, appID)

		if err == nil {
			log.Println("services : ", svcs.ServicesName)
            proj.Down()
		    proj.Delete()
            err = RemoveComposeProject(appID)
               
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

// ServiceScale scale all services
func ServiceScale(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var err error

	appID := vars["appId"]

	composeData, err := GetComposeProject(appID)
	if err == nil {
		var body []byte
		body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err == nil {
			if err = r.Body.Close(); err == nil {
				var scaleServices []scaleService
				if err = json.Unmarshal(body, &scaleServices); err == nil {
					var proj *project.Project
					proj, err = getProject(composeData, appID)

					if err == nil {
						log.Println("services scale : ", scaleServices)
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
