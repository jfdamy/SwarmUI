package api

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"path/filepath"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
    "github.com/docker/libcompose/docker"
)

//KvStore the kv store
var KvStore store.Store
var ClientFactory docker.ClientFactory

func init() {
	etcd.Register()
}

//Serve serve the api and the webui
func Serve() {

	router := NewRouter()

	var backend store.Backend
	switch os.Getenv("SWARMUI_KV_STORE") {
	case "etcd":
		backend = store.ETCD
	case "consul":
		backend = store.CONSUL
	case "zk":
		backend = store.ZK
	}

	if backend == "" {
		panic(errors.New("No KV store selected"))
	}

	host := os.Getenv("SWARMUI_KV_HOST")

	if host == "" {
		panic(errors.New("No KV host"))
	}

	KvStore, _ = libkv.NewStore(
		backend,
		strings.Split(host, ","),
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
    
    /*
        If the SWARMUI_SWARM_REPLICATION get the leader host with the key "/docker/swarm/leader" from the KV store
        Else use default env vars to connect to the docker daemon (DOCKER_HOST, ...)
    */
    if os.Getenv("SWARMUI_SWARM_REPLICATION") == "true" {
        swarmLearder, err := GetSwarmLeader()
        if err != nil {
            panic(err)
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
        
        if err != nil {
            panic(err)
        }
    }
    
	log.Fatal(http.ListenAndServe(":8080", router))
}
