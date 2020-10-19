package client

import "github.com/dghubble/sling"

type scheduledProjectTriggerService struct {
	service
}

func newScheduledProjectTriggerService(sling *sling.Sling, uriTemplate string) *scheduledProjectTriggerService {
	return &scheduledProjectTriggerService{
		service: newService(serviceScheduledProjectTriggerService, sling, uriTemplate, nil),
	}
}
