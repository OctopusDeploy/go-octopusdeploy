package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type PerformanceConfigurationService struct {
	services.Service
}

func NewPerformanceConfigurationService(sling *sling.Sling, uriTemplate string) *PerformanceConfigurationService {
	return &PerformanceConfigurationService{
		Service: services.NewService(constants.ServicePerformanceConfigurationService, sling, uriTemplate),
	}
}
