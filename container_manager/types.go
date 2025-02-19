package containermanger

import (
	"context"

	"github.com/docker/docker/client"
)

type Manager struct {
	ctx    context.Context
	client *client.Client
}

type Service struct {
	Name         string            `json:"name"`
	Image        string            `json:"image"`
	Command      []string          `json:"command,omitempty"`
	Env          map[string]string `json:"env,omitempty"`
	VolumeMounts []VolumeMount     `json:"volumemounts,omitempty"`
	Networks     []string          `json:"networks,omitempty"`
	Replicas     uint64            `json:"replicas"`
}

type ServiceRealtimeInfo struct {
	Name              string                     `json:"name"`
	PlacementInfos    []ServiceTaskPlacementInfo `json:"placementinfos"`
	DesiredReplicas   int                        `json:"desiredreplicas"`
	RunningReplicas   int                        `json:"runningreplicas"`
	ReplicatedService bool                       `json:"replicatedservice"`
}

type ServiceTaskPlacementInfo struct {
	NodeID          string `json:"nodeid"`
	NodeName        string `json:"nodename"`
	IsManagerNode   bool   `json:"ismanagernode"`
	RunningReplicas int    `json:"runningreplicas"`
}

// Volume Mount
type VolumeMount struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	ReadOnly bool   `json:"readonly"`
}
