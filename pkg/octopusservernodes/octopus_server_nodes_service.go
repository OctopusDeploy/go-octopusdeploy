package octopusservernodes

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type OctopusServerNodeService struct {
	clusterSummaryPath string

	services.CanDeleteService
}

func NewOctopusServerNodeService(sling *sling.Sling, uriTemplate string, clusterSummaryPath string) *OctopusServerNodeService {
	return &OctopusServerNodeService{
		clusterSummaryPath: clusterSummaryPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceOctopusServerNodeService, sling, uriTemplate),
		},
	}
}
