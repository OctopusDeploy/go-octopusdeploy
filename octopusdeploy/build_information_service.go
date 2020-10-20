package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type buildInformationService struct {
	bulkPath string

	canDeleteService
}

func newBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *buildInformationService {
	buildInformationService := &buildInformationService{
		bulkPath: bulkPath,
	}
	buildInformationService.service = newService(serviceBuildInformationService, sling, uriTemplate, nil)

	return buildInformationService
}
