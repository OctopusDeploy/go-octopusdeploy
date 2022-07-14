package dashboard

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type DashboardConfigurationService struct {
	services.Service
}

func NewDashboardConfigurationService(sling *sling.Sling, uriTemplate string) *DashboardConfigurationService {
	return &DashboardConfigurationService{
		Service: services.NewService(constants.ServiceDashboardConfigurationService, sling, uriTemplate),
	}
}
