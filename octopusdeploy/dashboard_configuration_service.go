package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type dashboardConfigurationService struct {
	services.service
}

func newDashboardConfigurationService(sling *sling.Sling, uriTemplate string) *dashboardConfigurationService {
	return &dashboardConfigurationService{
		service: services.newService(ServiceDashboardConfigurationService, sling, uriTemplate),
	}
}
