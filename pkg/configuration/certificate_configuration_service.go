package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type CertificateConfigurationService struct {
	services.Service
}

func NewCertificateConfigurationService(sling *sling.Sling, uriTemplate string) *CertificateConfigurationService {
	return &CertificateConfigurationService{
		Service: services.NewService(constants.ServiceCertificateConfigurationService, sling, uriTemplate),
	}
}
