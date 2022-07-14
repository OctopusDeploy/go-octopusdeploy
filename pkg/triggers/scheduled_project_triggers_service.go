package triggers

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type ScheduledProjectTriggerService struct {
	services.Service
}

func NewScheduledProjectTriggerService(sling *sling.Sling, uriTemplate string) *ScheduledProjectTriggerService {
	return &ScheduledProjectTriggerService{
		Service: services.NewService(constants.ServiceScheduledProjectTriggerService, sling, uriTemplate),
	}
}
