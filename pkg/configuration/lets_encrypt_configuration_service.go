package configuration

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type LetsEncryptConfigurationService struct {
	services.Service
}

func NewLetsEncryptConfigurationService(sling *sling.Sling, uriTemplate string) *LetsEncryptConfigurationService {
	return &LetsEncryptConfigurationService{
		Service: services.NewService(constants.ServiceLetsEncryptConfigurationService, sling, uriTemplate),
	}
}
