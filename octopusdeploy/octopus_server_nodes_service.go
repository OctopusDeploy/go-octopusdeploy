package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type octopusServerNodeService struct {
	clusterSummaryPath string

	services.canDeleteService
}

func newOctopusServerNodeService(sling *sling.Sling, uriTemplate string, clusterSummaryPath string) *octopusServerNodeService {
	octopusServerNodeService := &octopusServerNodeService{
		clusterSummaryPath: clusterSummaryPath,
	}
	octopusServerNodeService.service = services.newService(ServiceOctopusServerNodeService, sling, uriTemplate)

	return octopusServerNodeService
}
