package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type FeaturesConfigurationService struct {
	services.Service
}

func NewFeaturesConfigurationService(sling *sling.Sling, uriTemplate string) *FeaturesConfigurationService {
	return &FeaturesConfigurationService{
		Service: services.NewService(constants.ServiceFeaturesConfigurationService, sling, uriTemplate),
	}
}
