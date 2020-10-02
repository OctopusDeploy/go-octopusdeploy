package model

// AzureWebAppEndpoint represents the an Azure web app-based endpoint.
type AzureWebAppEndpoint struct {
	ResourceGroupName string `json:"ResourceGroupName,omitempty"`
	WebAppName        string `json:"WebAppName,omitempty"`
	WebAppSlotName    int    `json:"WebAppSlotName"`

	endpoint
}

// NewAzureWebAppEndpoint creates a new endpoint for Azure web apps.
func NewAzureWebAppEndpoint() *AzureWebAppEndpoint {
	resource := &AzureWebAppEndpoint{}
	return resource
}
