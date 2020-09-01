package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

type MachineConnectivityPolicy struct {
	MachineConnectivityBehavior enum.MachineConnectivityBehavior `json:"MachineConnectivityBehavior,omitempty"`
}
