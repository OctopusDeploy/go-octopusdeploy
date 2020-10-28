package octopusdeploy

import "github.com/dghubble/sling"

type taskService struct {
	taskTypesPath string

	service
}

func newTaskService(sling *sling.Sling, uriTemplate string, taskTypesPath string) *taskService {
	return &taskService{
		taskTypesPath: taskTypesPath,
		service:       newService(ServiceTaskService, sling, uriTemplate),
	}
}
