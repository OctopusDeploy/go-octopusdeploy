package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type externalSecurityGroupProviderService struct {
	services.service
}

func newExternalSecurityGroupProviderService(sling *sling.Sling, uriTemplate string) *externalSecurityGroupProviderService {
	return &externalSecurityGroupProviderService{
		service: services.newService(ServiceExternalSecurityGroupProviderService, sling, uriTemplate),
	}
}
