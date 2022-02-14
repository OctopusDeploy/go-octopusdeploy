package services

import (
	"github.com/dghubble/sling"
)

type cloudTemplateService struct {
	service
}

func newCloudTemplateService(sling *sling.Sling, uriTemplate string) *cloudTemplateService {
	return &cloudTemplateService{
		service: newService(ServiceCloudTemplateService, sling, uriTemplate),
	}
}
