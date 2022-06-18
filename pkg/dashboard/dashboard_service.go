package dashboard

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type DashboardService struct {
	dashboardDynamicPath string

	services.Service
}

func NewDashboardService(sling *sling.Sling, uriTemplate string, dashboardDynamicPath string) *DashboardService {
	return &DashboardService{
		dashboardDynamicPath: dashboardDynamicPath,
		Service:              services.NewService(constants.ServiceDashboardService, sling, uriTemplate),
	}
}
