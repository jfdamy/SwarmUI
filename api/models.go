package api

import (
	"github.com/docker/libcompose/project"
)

type projectInfo struct {
	ProjectID string    `json:"projectId"`
	Services  []serviceInfo `json:"services"`
}

type serviceInfo struct {
	ServiceName string                 `json:"serviceName"`
	Containers  []containerInfo        `json:"containers"`
	Config      *project.ServiceConfig `json:"config"`
}

//containerInfo is a data model for the info of a project
type containerInfo struct {
	ContainerName string `json:"containerName"`
	IsRunning     bool   `json:"isRunning"`
	Port          []port `json:"port"`
}

type port struct {
	PortHost string `json:"portHost"`
	PortCont string `json:"portCont"`
}

//scaleService is a data model for settings the scaling of services
type scaleService struct {
	ServiceName string `json:"serviceName"`
	Number      int    `json:"number"`
}

//services is a data model for seletings multiple service by name
type services struct {
	ServicesName []string `json:"servicesName"`
}
