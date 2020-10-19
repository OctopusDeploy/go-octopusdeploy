package model

// CloudRegionEndpoint represents a cloud service endpoint.
type CloudRegionEndpoint struct {
	DefaultWorkerPoolID string `json:"DefaultWorkerPoolId"`

	endpoint
}

// NewCloudRegionEndpoint creates and initializes a new cloud service endpoint.
func NewCloudRegionEndpoint() *CloudRegionEndpoint {
	cloudRegionEndpoint := &CloudRegionEndpoint{
		endpoint: *newEndpoint("None"),
	}

	return cloudRegionEndpoint
}

// GetDefaultWorkerPoolID returns the default worker pool ID of this endpoint.
func (e CloudRegionEndpoint) GetDefaultWorkerPoolID() string {
	return e.DefaultWorkerPoolID
}

// SetDefaultWorkerPoolID sets the default worker pool ID of this endpoint.
func (e CloudRegionEndpoint) SetDefaultWorkerPoolID(defaultWorkerPoolID string) {
	e.DefaultWorkerPoolID = defaultWorkerPoolID
}

var _ IRunsOnAWorker = &CloudRegionEndpoint{}
