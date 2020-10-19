package client

import "github.com/dghubble/sling"

type smtpConfigurationService struct {
	isConfiguredPath string

	service
}

func newSMTPConfigurationService(sling *sling.Sling, uriTemplate string, isConfiguredPath string) *smtpConfigurationService {
	return &smtpConfigurationService{
		isConfiguredPath: isConfiguredPath,
		service:          newService(serviceSMTPConfigurationService, sling, uriTemplate, nil),
	}
}
