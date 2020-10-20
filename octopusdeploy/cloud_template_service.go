package octopusdeploy

import "github.com/dghubble/sling"

type cloudTemplateService struct {
	service
}

func newCloudTemplateService(sling *sling.Sling, uriTemplate string) *cloudTemplateService {
	return &cloudTemplateService{
		service: newService(serviceCloudTemplateService, sling, uriTemplate, nil),
	}
}
