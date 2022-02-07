package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type serverConfigurationService struct {
	settingsPath string

	services.service
}

func newServerConfigurationService(sling *sling.Sling, uriTemplate string, settingsPath string) *serverConfigurationService {
	return &serverConfigurationService{
		settingsPath: settingsPath,
		service:      services.newService(ServiceServerConfigurationService, sling, uriTemplate),
	}
}
