package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type featuresConfigurationService struct {
	services.service
}

func newFeaturesConfigurationService(sling *sling.Sling, uriTemplate string) *featuresConfigurationService {
	return &featuresConfigurationService{
		service: services.newService(ServiceFeaturesConfigurationService, sling, uriTemplate),
	}
}
