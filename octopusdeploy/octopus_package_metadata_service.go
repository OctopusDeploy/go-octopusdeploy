package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type octopusPackageMetadataService struct {
	services.service
}

func newOctopusPackageMetadataService(sling *sling.Sling, uriTemplate string) *octopusPackageMetadataService {
	return &octopusPackageMetadataService{
		service: services.newService(ServiceOctopusPackageMetadataService, sling, uriTemplate),
	}
}
