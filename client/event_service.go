package client

import "github.com/dghubble/sling"

type eventService struct {
	agentsPath        string
	categoriesPath    string
	documentTypesPath string
	groupsPath        string

	service
}

func newEventService(sling *sling.Sling, uriTemplate string, agentsPath string, categoriesPath string, documentTypesPath string, groupsPath string) *eventService {
	return &eventService{
		service: newService(serviceEventService, sling, uriTemplate, nil),
	}
}
