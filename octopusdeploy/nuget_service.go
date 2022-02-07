package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type nuGetService struct {
	services.service
}

func newNuGetService(sling *sling.Sling, uriTemplate string) *nuGetService {
	return &nuGetService{
		service: services.newService(ServiceNuGetService, sling, uriTemplate),
	}
}
