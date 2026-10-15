package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type SmtpConfigurationService struct {
	isConfiguredPath string

	services.Service
}

func NewSmtpConfigurationService(sling *sling.Sling, uriTemplate string, isConfiguredPath string) *SmtpConfigurationService {
	return &SmtpConfigurationService{
		isConfiguredPath: isConfiguredPath,
		Service:          services.NewService(constants.ServiceSMTPConfigurationService, sling, uriTemplate),
	}
}
