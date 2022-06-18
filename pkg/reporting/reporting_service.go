package reporting

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ReportingService struct {
	deploymentsCountedByWeekPath string

	services.Service
}

func NewReportingService(sling *sling.Sling, uriTemplate string, deploymentsCountedByWeekPath string) *ReportingService {
	return &ReportingService{
		deploymentsCountedByWeekPath: deploymentsCountedByWeekPath,
		Service:                      services.NewService(constants.ServiceReportingService, sling, uriTemplate),
	}
}
