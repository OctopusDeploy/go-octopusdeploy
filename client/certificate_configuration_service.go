package client

import "github.com/dghubble/sling"

type certificateConfigurationService struct {
	service
}

func newCertificateConfigurationService(sling *sling.Sling, uriTemplate string) *certificateConfigurationService {
	return &certificateConfigurationService{
		service: newService(serviceCertificateConfigurationService, sling, uriTemplate, nil),
	}
}
