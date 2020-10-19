package client

import "github.com/dghubble/sling"

type schedulerService struct {
	service
}

func newSchedulerService(sling *sling.Sling, uriTemplate string) *schedulerService {
	return &schedulerService{
		service: newService(serviceSchedulerService, sling, uriTemplate, nil),
	}
}
