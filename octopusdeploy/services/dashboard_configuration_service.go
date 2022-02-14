package services

import (
	"github.com/dghubble/sling"
)

type dashboardConfigurationService struct {
	service
}

func newDashboardConfigurationService(sling *sling.Sling, uriTemplate string) *dashboardConfigurationService {
	return &dashboardConfigurationService{
		service: newService(ServiceDashboardConfigurationService, sling, uriTemplate),
	}
}
