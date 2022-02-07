package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type certificateConfigurationService struct {
	services.service
}

func newCertificateConfigurationService(sling *sling.Sling, uriTemplate string) *certificateConfigurationService {
	return &certificateConfigurationService{
		service: services.newService(ServiceCertificateConfigurationService, sling, uriTemplate),
	}
}
