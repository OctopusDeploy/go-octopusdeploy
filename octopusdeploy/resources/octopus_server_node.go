package resources

type OctopusServerNodeResource struct {
	IsInMaintenanceMode bool   `json:"IsInMaintenanceMode,omitempty"`
	MaxConcurrentTasks  int32  `json:"MaxConcurrentTasks,omitempty"`
	Name                string `json:"Name,omitempty"`

	Resource
}

func NewOctopusServerNodeResource() *OctopusServerNodeResource {
	return &OctopusServerNodeResource{
		Resource: *NewResource(),
	}
}
