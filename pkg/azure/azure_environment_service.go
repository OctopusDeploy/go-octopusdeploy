package azure

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type AzureEnvironmentService struct {
	services.Service
}

func NewAzureEnvironmentService(sling *sling.Sling, uriTemplate string) *AzureEnvironmentService {
	return &AzureEnvironmentService{
		Service: services.NewService(constants.ServiceAzureEnvironmentService, sling, uriTemplate),
	}
}
