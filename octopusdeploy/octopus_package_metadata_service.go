package octopusdeploy

import "github.com/dghubble/sling"

type octopusPackageMetadataService struct {
	service
}

func newOctopusPackageMetadataService(sling *sling.Sling, uriTemplate string) *octopusPackageMetadataService {
	return &octopusPackageMetadataService{
		service: newService(ServiceOctopusPackageMetadataService, sling, uriTemplate),
	}
}
