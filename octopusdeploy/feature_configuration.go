package octopusdeploy

import "github.com/dghubble/sling"

type featuresConfigurationService struct {
	service
}

func newFeaturesConfigurationService(sling *sling.Sling, uriTemplate string) *featuresConfigurationService {
	return &featuresConfigurationService{
		service: newService(serviceFeaturesConfigurationService, sling, uriTemplate, nil),
	}
}
