package client

import "github.com/dghubble/sling"

type issueTrackerService struct {
	service
}

func newIssueTrackerService(sling *sling.Sling, uriTemplate string) *issueTrackerService {
	return &issueTrackerService{
		service: newService(serviceIssueTrackerService, sling, uriTemplate, nil),
	}
}
