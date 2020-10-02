package model

type CloudRegionEndpoint struct {
	DefaultWorkerPoolID string `json:"DefaultWorkerPoolId" validate:"required"`

	endpoint
}

func NewCloudRegionEndpoint() *CloudRegionEndpoint {
	resource := &CloudRegionEndpoint{}
	resource.CommunicationStyle = "None"

	return resource
}
