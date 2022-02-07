package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type issueTrackerService struct {
	services.service
}

func newIssueTrackerService(sling *sling.Sling, uriTemplate string) *issueTrackerService {
	return &issueTrackerService{
		service: services.newService(ServiceIssueTrackerService, sling, uriTemplate),
	}
}
