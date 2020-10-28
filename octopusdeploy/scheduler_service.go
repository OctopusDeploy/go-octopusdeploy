package octopusdeploy

import "github.com/dghubble/sling"

type schedulerService struct {
	service
}

func newSchedulerService(sling *sling.Sling, uriTemplate string) *schedulerService {
	return &schedulerService{
		service: newService(ServiceSchedulerService, sling, uriTemplate),
	}
}
