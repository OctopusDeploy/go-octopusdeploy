package octopusservernodes

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type OctopusServerNodeResource struct {
	IsInMaintenanceMode bool   `json:"IsInMaintenanceMode"`
	MaxConcurrentTasks  int32  `json:"MaxConcurrentTasks,omitempty"`
	Name                string `json:"Name,omitempty"`

	resources.Resource
}

func NewOctopusServerNodeResource() *OctopusServerNodeResource {
	return &OctopusServerNodeResource{
		Resource: *resources.NewResource(),
	}
}
