package octopusdeploy

import (
	"time"
)

type MachineConnectionStatus struct {
	CurrentTentacleVersion string                `json:"CurrentTentacleVersion,omitempty"`
	LastChecked            time.Time             `json:"LastChecked,omitempty"`
	Logs                   []*ActivityLogElement `json:"Logs"`
	MachineID              string                `json:"MachineId,omitempty"`
	Status                 string                `json:"Status,omitempty"`

	resource
}

func NewMachineConnectionStatus() *MachineConnectionStatus {
	return &MachineConnectionStatus{
		resource: *newResource(),
	}
}
