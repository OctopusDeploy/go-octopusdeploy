package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type performanceConfigurationService struct {
	services.service
}

func newPerformanceConfigurationService(sling *sling.Sling, uriTemplate string) *performanceConfigurationService {
	return &performanceConfigurationService{
		service: services.newService(ServicePerformanceConfigurationService, sling, uriTemplate),
	}
}
