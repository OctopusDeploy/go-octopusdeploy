package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type schedulerService struct {
	services.service
}

func newSchedulerService(sling *sling.Sling, uriTemplate string) *schedulerService {
	return &schedulerService{
		service: services.newService(ServiceSchedulerService, sling, uriTemplate),
	}
}
