package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type UpgradeConfigurationService struct {
	services.Service
}

func NewUpgradeConfigurationService(sling *sling.Sling, uriTemplate string) *UpgradeConfigurationService {
	return &UpgradeConfigurationService{
		Service: services.NewService(constants.ServiceUpgradeConfigurationService, sling, uriTemplate),
	}
}
