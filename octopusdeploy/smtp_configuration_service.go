package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type smtpConfigurationService struct {
	isConfiguredPath string

	services.service
}

func newSMTPConfigurationService(sling *sling.Sling, uriTemplate string, isConfiguredPath string) *smtpConfigurationService {
	return &smtpConfigurationService{
		isConfiguredPath: isConfiguredPath,
		service:          services.newService(ServiceSMTPConfigurationService, sling, uriTemplate),
	}
}
