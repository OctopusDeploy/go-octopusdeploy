package service

import (
	"github.com/dghubble/sling"
)

type octopusServerNodeService struct {
	clusterSummaryPath string

	canDeleteService
}

func newOctopusServerNodeService(sling *sling.Sling, uriTemplate string, clusterSummaryPath string) *octopusServerNodeService {
	octopusServerNodeService := &octopusServerNodeService{
		clusterSummaryPath: clusterSummaryPath,
	}
	octopusServerNodeService.service = newService(ServiceOctopusServerNodeService, sling, uriTemplate)

	return octopusServerNodeService
}
