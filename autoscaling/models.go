package autoscaling

import (
    "github.com/docker/libcompose/project"
)

const (
    autoscaling = "auto"
    nodescaling = "node"
)

//ScalingService config for autoscaling of a service
type ScalingService struct {
    ScalingType         string              `json:"scalingType"`
    ServiceName         string              `json:"serviceName"`
    containerNumber     int
    service             project.Service
}