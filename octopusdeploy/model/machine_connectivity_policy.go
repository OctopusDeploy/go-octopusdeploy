package model

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"

type MachineConnectivityPolicy struct {
	MachineConnectivityBehavior enum.MachineConnectivityBehavior `json:"MachineConnectivityBehavior,omitempty"`
}
