package extensions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type DynamicExtensionService struct {
	featuresMetadataPath string
	featuresValuesPath   string
	scriptsPath          string

	services.Service
}

func NewDynamicExtensionService(sling *sling.Sling, uriTemplate string, featuresMetadataPath string, featuresValuesPath string, scriptsPath string) *DynamicExtensionService {
	return &DynamicExtensionService{
		featuresMetadataPath: featuresMetadataPath,
		featuresValuesPath:   featuresValuesPath,
		scriptsPath:          scriptsPath,
		Service:              services.NewService(constants.ServiceDynamicExtensionService, sling, uriTemplate),
	}
}
