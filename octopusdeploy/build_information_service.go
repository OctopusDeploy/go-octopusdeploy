package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type buildInformationService struct {
	bulkPath string

	services.canDeleteService
}

func newBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *buildInformationService {
	buildInformationService := &buildInformationService{
		bulkPath: bulkPath,
	}
	buildInformationService.service = services.newService(ServiceBuildInformationService, sling, uriTemplate)

	return buildInformationService
}
