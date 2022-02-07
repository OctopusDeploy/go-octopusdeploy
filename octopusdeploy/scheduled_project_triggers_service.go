package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type scheduledProjectTriggerService struct {
	services.service
}

func newScheduledProjectTriggerService(sling *sling.Sling, uriTemplate string) *scheduledProjectTriggerService {
	return &scheduledProjectTriggerService{
		service: services.newService(ServiceScheduledProjectTriggerService, sling, uriTemplate),
	}
}
