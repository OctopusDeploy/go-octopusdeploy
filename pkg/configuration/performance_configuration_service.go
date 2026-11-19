package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type PerformanceConfigurationService struct {
	services.Service
}

func NewPerformanceConfigurationService(sling *sling.Sling, uriTemplate string) *PerformanceConfigurationService {
	return &PerformanceConfigurationService{
		Service: services.NewService(constants.ServicePerformanceConfigurationService, sling, uriTemplate),
	}
}
