package buildinformation

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type BuildInformationService struct {
	bulkPath string

	services.CanDeleteService
}

func NewBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *BuildInformationService {
	return &BuildInformationService{
		bulkPath: bulkPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceBuildInformationService, sling, uriTemplate),
		},
	}
}
