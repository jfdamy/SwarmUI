package api

import (
	"fmt"
	"strings"
)

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