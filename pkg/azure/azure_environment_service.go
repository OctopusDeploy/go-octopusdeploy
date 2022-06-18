package azure

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
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
