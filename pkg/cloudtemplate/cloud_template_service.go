package cloudtemplate

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type CloudTemplateService struct {
	services.Service
}

func NewCloudTemplateService(sling *sling.Sling, uriTemplate string) *CloudTemplateService {
	return &CloudTemplateService{
		Service: services.NewService(constants.ServiceCloudTemplateService, sling, uriTemplate),
	}
}
