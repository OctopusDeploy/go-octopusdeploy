package octopusservernodes

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
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
