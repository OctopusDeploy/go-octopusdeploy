package octopusdeploy

import "github.com/dghubble/sling"

type certificateConfigurationService struct {
	service
}

func newCertificateConfigurationService(sling *sling.Sling, uriTemplate string) *certificateConfigurationService {
	return &certificateConfigurationService{
		service: newService(ServiceCertificateConfigurationService, sling, uriTemplate),
	}
}
