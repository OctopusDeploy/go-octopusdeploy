package machines

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type MachineConnectionStatus struct {
	CurrentTentacleVersion string                `json:"CurrentTentacleVersion,omitempty"`
	LastChecked            time.Time             `json:"LastChecked,omitempty"`
	Logs                   []*ActivityLogElement `json:"Logs"`
	MachineID              string                `json:"MachineId,omitempty"`
	Status                 string                `json:"Status,omitempty"`

	resources.Resource
}

func NewMachineConnectionStatus() *MachineConnectionStatus {
	return &MachineConnectionStatus{
		Resource: *resources.NewResource(),
	}
}
