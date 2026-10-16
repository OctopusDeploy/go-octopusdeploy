package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type MaintenanceConfigurationService struct {
	services.Service
}

func NewMaintenanceConfigurationService(sling *sling.Sling, uriTemplate string) *MaintenanceConfigurationService {
	return &MaintenanceConfigurationService{
		Service: services.NewService(constants.ServiceMaintenanceConfigurationService, sling, uriTemplate),
	}
}
