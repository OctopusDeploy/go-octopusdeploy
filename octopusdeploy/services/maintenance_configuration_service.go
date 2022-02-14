package services

import (
	"github.com/dghubble/sling"
)

type maintenanceConfigurationService struct {
	service
}

func newMaintenanceConfigurationService(sling *sling.Sling, uriTemplate string) *maintenanceConfigurationService {
	return &maintenanceConfigurationService{
		service: newService(ServiceMaintenanceConfigurationService, sling, uriTemplate),
	}
}
