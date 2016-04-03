package utils

import (
	"fmt"
	"strings"
    "os"
    "errors"
    "time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
	"github.com/docker/libkv/store/consul"
)


//KvStore the kv store
var KvStore store.Store

//InitStore the KV store
func InitStore() error {
    
    var backend store.Backend
	switch os.Getenv("SWARMUI_KV_STORE") {
	case "etcd":
		backend = store.ETCD
	    etcd.Register()
	case "consul":
		backend = store.CONSUL
        consul.Register()
	case "zk":
		backend = store.ZK
	}

	if backend == "" {
		return errors.New("No KV store selected")
	}

	host := os.Getenv("SWARMUI_KV_HOST")

	if host == "" {
		return errors.New("No KV host")
	}

    var err error
	KvStore, err = libkv.NewStore(
		backend,
		strings.Split(host, ","),
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
    return err
}

//GetSwarmLeader get the swarm leader
func GetSwarmLeader() (string, error) {
    pair, err := KvStore.Get("/docker/swarm/leader")
	if err != nil {
		return "", fmt.Errorf("Error trying accessing the swarm leader")
	}
	return string(pair.Value), nil
}

//GetComposeProject get a project definition
func GetComposeProject(projectID string) ([]byte, error) {
	pair, err := KvStore.Get("/swarmui/project/" + projectID)
	if err != nil {
		return nil, fmt.Errorf("Error trying accessing project at ID: %v", projectID)
	}
	return pair.Value, nil
}

//GetListComposeProject get the list of project
func GetListComposeProject() ([]string, error) {
	pairs, err := KvStore.List("/swarmui/project/")
	if err != nil {
		return nil, fmt.Errorf("Error trying accessing the list of projects")
	}
	var ret []string
	for _, pair := range pairs {
		ret = append(ret, strings.Replace(string(pair.Key), "/swarmui/project/", "", 1))
	}
	return ret, nil
}

//SetComposeProject set a project definition
func SetComposeProject(projectID string, value []byte) error {
	err := KvStore.Put("/swarmui/project/"+projectID, value, nil)
	if err != nil {
		return fmt.Errorf("Error trying setting project at ID: %v", projectID)
	}
	return nil
}

//RemoveComposeProject get a project definition
func RemoveComposeProject(projectID string) error {
	return KvStore.Delete("/swarmui/project/" + projectID)
}

//WatchAutoscalingProjects watch autoscaling projects from kv store
func WatchAutoscalingProjects(stop <-chan struct{}) (<- chan []*store.KVPair, error) {
    return KvStore.WatchTree("/swarmui/autoscaling/", stop)
}

//SetAutoscalingProject set autoscaling projects
func SetAutoscalingProject(projectID string, jsonAutoscaling []byte) error {
    return KvStore.Put("/swarmui/autoscaling/"+projectID, jsonAutoscaling, nil)
}

//GetAutoscalingProject set autoscaling projects
func GetAutoscalingProject(projectID string) (*store.KVPair, error) {
    return KvStore.Get("/swarmui/autoscaling/"+projectID)
}

//GetListAutoscalingProjects get autoscaling projects from kv store
func GetListAutoscalingProjects() ([] *store.KVPair, error) {
    return KvStore.List("/swarmui/autoscaling/")
}