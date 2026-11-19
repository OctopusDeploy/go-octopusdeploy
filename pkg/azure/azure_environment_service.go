package azure

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type AzureEnvironmentService struct {
	services.Service
}

func NewAzureEnvironmentService(sling *sling.Sling, uriTemplate string) *AzureEnvironmentService {
	return &AzureEnvironmentService{
		Service: services.NewService(constants.ServiceAzureEnvironmentService, sling, uriTemplate),
	}
}
