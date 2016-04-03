//Package api hold the REST API of the Swarm UI project
package api

import (
	"net/http"
    
	log "github.com/Sirupsen/logrus"
    "github.com/jfdamy/swarmui/utils"
)

//Serve serve the api and the webui
func Serve() {

	router := NewRouter()

	err := utils.InitStore()
    if err != nil {
        panic(err)
    }
    
    err = utils.InitDocker()
    if err != nil {
        panic(err)
    }
    listen := ":8080"
    log.Info("Swarmui API Listen to ", listen)
	log.Fatal(http.ListenAndServe(listen, router))
}
