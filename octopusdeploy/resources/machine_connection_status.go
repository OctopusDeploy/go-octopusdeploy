package resources

import (
	"time"
)

type MachineConnectionStatus struct {
	CurrentTentacleVersion string                `json:"CurrentTentacleVersion,omitempty"`
	LastChecked            time.Time             `json:"LastChecked,omitempty"`
	Logs                   []*ActivityLogElement `json:"Logs"`
	MachineID              string                `json:"MachineId,omitempty"`
	Status                 string                `json:"Status,omitempty"`

	Resource
}

func NewMachineConnectionStatus() *MachineConnectionStatus {
	return &MachineConnectionStatus{
		Resource: *NewResource(),
	}
}
