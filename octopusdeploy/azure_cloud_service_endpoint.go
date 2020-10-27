package octopusdeploy

// AzureCloudServiceEndpoint represents an Azure cloud service endpoint.
type AzureCloudServiceEndpoint struct {
	AccountID               string `json:"AccountId"`
	CloudServiceName        string `json:"CloudServiceName"`
	DefaultWorkerPoolID     string `json:"DefaultWorkerPoolId"`
	Slot                    string `json:"Slot"`
	StorageAccountName      string `json:"StorageAccountName"`
	SwapIfPossible          bool   `json:"SwapIfPossible"`
	UseCurrentInstanceCount bool   `json:"UseCurrentInstanceCount"`

	endpoint
}

// NewAzureCloudServiceEndpoint creates and initializes a new Azure cloud
// service endpoint.
func NewAzureCloudServiceEndpoint() *AzureCloudServiceEndpoint {
	ServiceFabricEndpoint := &AzureCloudServiceEndpoint{
		endpoint: *newEndpoint("AzureCloudService"),
	}

	return ServiceFabricEndpoint
}

// GetAccountID returns the account ID associated with this Azure cloud service
// endpoint.
func (a AzureCloudServiceEndpoint) GetAccountID() string {
	return a.AccountID
}

// GetDefaultWorkerPoolID returns the default worker pool ID of this Azure
// cloud service endpoint.
func (a AzureCloudServiceEndpoint) GetDefaultWorkerPoolID() string {
	return a.DefaultWorkerPoolID
}

// SetDefaultWorkerPoolID sets the default worker pool ID of this Azure cloud
// service endpoint.
func (a AzureCloudServiceEndpoint) SetDefaultWorkerPoolID(defaultWorkerPoolID string) {
	a.DefaultWorkerPoolID = defaultWorkerPoolID
}

var _ IEndpointWithAccount = &AzureCloudServiceEndpoint{}
var _ IRunsOnAWorker = &AzureCloudServiceEndpoint{}
