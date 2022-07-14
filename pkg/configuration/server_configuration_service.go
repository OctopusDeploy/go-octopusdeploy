package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type ServerConfigurationService struct {
	settingsPath string

	services.Service
}

func NewServerConfigurationService(sling *sling.Sling, uriTemplate string, settingsPath string) *ServerConfigurationService {
	return &ServerConfigurationService{
		settingsPath: settingsPath,
		Service:      services.NewService(constants.ServiceServerConfigurationService, sling, uriTemplate),
	}
}
