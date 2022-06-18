package externalsecuritygroupproviders

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ExternalSecurityGroupProviderService struct {
	services.Service
}

func NewExternalSecurityGroupProviderService(sling *sling.Sling, uriTemplate string) *ExternalSecurityGroupProviderService {
	return &ExternalSecurityGroupProviderService{
		Service: services.NewService(constants.ServiceExternalSecurityGroupProviderService, sling, uriTemplate),
	}
}
