package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type packageMetadataService struct {
	services.service
}

func newPackageMetadataService(sling *sling.Sling, uriTemplate string) *packageMetadataService {
	return &packageMetadataService{
		service: services.newService(ServicePackageMetadataService, sling, uriTemplate),
	}
}
