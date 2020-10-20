package octopusdeploy

import "github.com/dghubble/sling"

type performanceConfigurationService struct {
	service
}

func newPerformanceConfigurationService(sling *sling.Sling, uriTemplate string) *performanceConfigurationService {
	return &performanceConfigurationService{
		service: newService(servicePerformanceConfigurationService, sling, uriTemplate, nil),
	}
}
