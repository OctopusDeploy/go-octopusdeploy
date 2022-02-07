package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type letsEncryptConfigurationService struct {
	services.service
}

func newLetsEncryptConfigurationService(sling *sling.Sling, uriTemplate string) *letsEncryptConfigurationService {
	return &letsEncryptConfigurationService{
		service: services.newService(ServiceLetsEncryptConfigurationService, sling, uriTemplate),
	}
}
