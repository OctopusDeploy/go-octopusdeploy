package scheduler

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type SchedulerService struct {
	services.Service
}

func NewSchedulerService(sling *sling.Sling, uriTemplate string) *SchedulerService {
	return &SchedulerService{
		Service: services.NewService(constants.ServiceSchedulerService, sling, uriTemplate),
	}
}
