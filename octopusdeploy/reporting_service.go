package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type reportingService struct {
	deploymentsCountedByWeekPath string

	services.service
}

func newReportingService(sling *sling.Sling, uriTemplate string, deploymentsCountedByWeekPath string) *reportingService {
	return &reportingService{
		deploymentsCountedByWeekPath: deploymentsCountedByWeekPath,

		service: services.newService(ServiceReportingService, sling, uriTemplate),
	}
}
