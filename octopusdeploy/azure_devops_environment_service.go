package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type azureEnvironmentService struct {
	services.service
}

func newAzureEnvironmentService(sling *sling.Sling, uriTemplate string) *azureEnvironmentService {
	return &azureEnvironmentService{
		service: services.newService(ServiceAzureEnvironmentService, sling, uriTemplate),
	}
}
