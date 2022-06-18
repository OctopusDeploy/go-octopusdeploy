package machines

type CloudServiceEndpoint struct {
	CloudServiceName        string `json:"CloudServiceName,omitempty"`
	Slot                    string `json:"Slot,omitempty"`
	StorageAccountName      string `json:"StorageAccountName,omitempty"`
	SwapIfPossible          bool   `json:"SwapIfPossible"`
	UseCurrentInstanceCount bool   `json:"UseCurrentInstanceCount"`

	endpoint
}

func NewCloudServiceEndpoint() *CloudServiceEndpoint {
	cloudServiceEndpoint := &CloudServiceEndpoint{
		endpoint: *newEndpoint("None"),
	}
	return cloudServiceEndpoint
}
