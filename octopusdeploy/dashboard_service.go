package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type dashboardService struct {
	dashboardDynamicPath string

	services.service
}

func newDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicPath string) *dashboardService {
	return &dashboardService{
		dashboardDynamicPath: dashboardDynamicPath,
		service:              services.newService(ServiceDashboardService, sling, uriTemplate),
	}
}
