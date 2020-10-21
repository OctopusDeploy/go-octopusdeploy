package octopusdeploy

type OctopusServerNodeResource struct {
	IsInMaintenanceMode bool   `json:"IsInMaintenanceMode,omitempty"`
	MaxConcurrentTasks  int32  `json:"MaxConcurrentTasks,omitempty"`
	Name                string `json:"Name,omitempty"`

	resource
}

func NewOctopusServerNodeResource() *OctopusServerNodeResource {
	return &OctopusServerNodeResource{
		resource: *newResource(),
	}
}
