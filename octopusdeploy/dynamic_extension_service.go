package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type dynamicExtensionService struct {
	featuresMetadataPath string
	featuresValuesPath   string
	scriptsPath          string

	services.service
}

func newDynamicExtensionService(sling *sling.Sling, uriTemplate string, featuresMetadataPath string, featuresValuesPath string, scriptsPath string) *dynamicExtensionService {
	return &dynamicExtensionService{
		featuresMetadataPath: featuresMetadataPath,
		featuresValuesPath:   featuresValuesPath,
		scriptsPath:          scriptsPath,
		service:              services.newService(ServiceDynamicExtensionService, sling, uriTemplate),
	}
}
