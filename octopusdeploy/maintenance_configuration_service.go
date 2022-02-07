package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type maintenanceConfigurationService struct {
	services.service
}

func newMaintenanceConfigurationService(sling *sling.Sling, uriTemplate string) *maintenanceConfigurationService {
	return &maintenanceConfigurationService{
		service: services.newService(ServiceMaintenanceConfigurationService, sling, uriTemplate),
	}
}
