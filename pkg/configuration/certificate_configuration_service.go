package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type CertificateConfigurationService struct {
	services.Service
}

func NewCertificateConfigurationService(sling *sling.Sling, uriTemplate string) *CertificateConfigurationService {
	return &CertificateConfigurationService{
		Service: services.NewService(constants.ServiceCertificateConfigurationService, sling, uriTemplate),
	}
}
