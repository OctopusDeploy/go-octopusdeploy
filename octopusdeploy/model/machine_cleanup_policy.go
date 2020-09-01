package model

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"

type MachineCleanupPolicy struct {
	DeleteMachinesBehavior        enum.DeleteMachinesBehavior `json:"DeleteMachinesBehavior,omitempty"`
	DeleteMachinesElapsedTimeSpan string                      `json:"DeleteMachinesElapsedTimeSpan,omitempty"`
}
