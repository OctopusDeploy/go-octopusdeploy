package service

import (
	"github.com/dghubble/sling"
)

type dynamicExtensionService struct {
	featuresMetadataPath string
	featuresValuesPath   string
	scriptsPath          string

	service
}

func newDynamicExtensionService(sling *sling.Sling, uriTemplate string, featuresMetadataPath string, featuresValuesPath string, scriptsPath string) *dynamicExtensionService {
	return &dynamicExtensionService{
		featuresMetadataPath: featuresMetadataPath,
		featuresValuesPath:   featuresValuesPath,
		scriptsPath:          scriptsPath,
		service:              newService(ServiceDynamicExtensionService, sling, uriTemplate),
	}
}
