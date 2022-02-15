package service

import (
	"github.com/dghubble/sling"
)

type serverConfigurationService struct {
	settingsPath string

	service
}

func newServerConfigurationService(sling *sling.Sling, uriTemplate string, settingsPath string) *serverConfigurationService {
	return &serverConfigurationService{
		settingsPath: settingsPath,
		service:      newService(ServiceServerConfigurationService, sling, uriTemplate),
	}
}
