package autoscaling

import (
    "github.com/docker/libcompose/project"
)

const (
    autoscaling = "auto"
    nodescaling = "node"
)

//AutoScalingConfig auto scaling configuration
type AutoScalingConfig struct {
    //MinNumber min number of container (default 1)
    MinNumber   int
    //MinNumber max number of container (default 5)
    MaxNumber   int
    //MinCPUUsage min cpu usage to scale down (default 0.15)
    MinCPUUsage float64
    //MaxCPUUsage max cpu usage to scale up (default 0.70)
    MaxCPUUsage float64
}

//ScalingService config for autoscaling of a service
type ScalingService struct {
    ScalingType         string              `json:"scalingType"`
    ServiceName         string              `json:"serviceName"`
    AutoScalingConf   *AutoScalingConfig  `json:"AutoScalingConfig"`
    containerNumber     int
    service             project.Service
}