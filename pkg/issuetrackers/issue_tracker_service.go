package issuetrackers

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type IssueTrackerService struct {
	services.Service
}

func NewIssueTrackerService(sling *sling.Sling, uriTemplate string) *IssueTrackerService {
	return &IssueTrackerService{
		Service: services.NewService(constants.ServiceIssueTrackerService, sling, uriTemplate),
	}
}
