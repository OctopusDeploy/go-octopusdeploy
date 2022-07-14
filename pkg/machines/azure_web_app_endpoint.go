package machines

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

// AzureWebAppEndpoint represents the an Azure web app-based endpoint.
type AzureWebAppEndpoint struct {
	AccountID         string `json:"AccountId"`
	ResourceGroupName string `json:"ResourceGroupName,omitempty"`
	WebAppName        string `json:"WebAppName,omitempty"`
	WebAppSlotName    string `json:"WebAppSlotName"`

	endpoint
}

// NewAzureWebAppEndpoint creates a new endpoint for Azure web apps.
func NewAzureWebAppEndpoint() *AzureWebAppEndpoint {
	azureWebAppEndpoint := &AzureWebAppEndpoint{
		endpoint: *newEndpoint("AzureWebApp"),
	}

	return azureWebAppEndpoint
}

var _ resources.IResource = &AzureWebAppEndpoint{}
var _ IEndpoint = &AzureWebAppEndpoint{}
