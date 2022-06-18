package configuration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type LetsEncryptConfigurationService struct {
	services.Service
}

func NewLetsEncryptConfigurationService(sling *sling.Sling, uriTemplate string) *LetsEncryptConfigurationService {
	return &LetsEncryptConfigurationService{
		Service: services.NewService(constants.ServiceLetsEncryptConfigurationService, sling, uriTemplate),
	}
}
