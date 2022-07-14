package issuetrackers

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type IssueTrackerService struct {
	services.Service
}

func NewIssueTrackerService(sling *sling.Sling, uriTemplate string) *IssueTrackerService {
	return &IssueTrackerService{
		Service: services.NewService(constants.ServiceIssueTrackerService, sling, uriTemplate),
	}
}
