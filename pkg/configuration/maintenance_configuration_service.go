package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type MaintenanceConfigurationService struct {
	services.Service
}

func NewMaintenanceConfigurationService(sling *sling.Sling, uriTemplate string) *MaintenanceConfigurationService {
	return &MaintenanceConfigurationService{
		Service: services.NewService(constants.ServiceMaintenanceConfigurationService, sling, uriTemplate),
	}
}
