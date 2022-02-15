package service

import (
	"github.com/dghubble/sling"
)

type scheduledProjectTriggerService struct {
	service
}

func newScheduledProjectTriggerService(sling *sling.Sling, uriTemplate string) *scheduledProjectTriggerService {
	return &scheduledProjectTriggerService{
		service: newService(ServiceScheduledProjectTriggerService, sling, uriTemplate),
	}
}
