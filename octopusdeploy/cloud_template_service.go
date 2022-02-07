package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type cloudTemplateService struct {
	services.service
}

func newCloudTemplateService(sling *sling.Sling, uriTemplate string) *cloudTemplateService {
	return &cloudTemplateService{
		service: services.newService(ServiceCloudTemplateService, sling, uriTemplate),
	}
}
