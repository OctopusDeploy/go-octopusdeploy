package scheduler

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type SchedulerService struct {
	services.Service
}

func NewSchedulerService(sling *sling.Sling, uriTemplate string) *SchedulerService {
	return &SchedulerService{
		Service: services.NewService(constants.ServiceSchedulerService, sling, uriTemplate),
	}
}
