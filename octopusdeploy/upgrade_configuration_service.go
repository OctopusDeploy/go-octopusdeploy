package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type upgradeConfigurationService struct {
	services.service
}

func newUpgradeConfigurationService(sling *sling.Sling, uriTemplate string) *upgradeConfigurationService {
	return &upgradeConfigurationService{
		service: services.newService(ServiceUpgradeConfigurationService, sling, uriTemplate),
	}
}
