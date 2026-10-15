package reporting

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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
