package tasks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type TaskService struct {
	taskTypesPath string

	services.Service
}

func NewTaskService(sling *sling.Sling, uriTemplate string, taskTypesPath string) *TaskService {
	return &TaskService{
		taskTypesPath: taskTypesPath,
		Service:       services.NewService(constants.ServiceTaskService, sling, uriTemplate),
	}
}

// Add creates a new task.
func (s *TaskService) Add(task *Task) (*Task, error) {
	if IsNil(task) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTask)
	}

	path, err := services.GetAddPath(s, task)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), task, new(Task), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Task), nil
}

// Get returns a collection of tasks based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s *TaskService) Get(tasksQuery TasksQuery) (*resources.Resources[*Task], error) {
	path, err := s.GetURITemplate().Expand(tasksQuery)
	if err != nil {
		return &resources.Resources[*Task]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Task]), path)
	if err != nil {
		return &resources.Resources[*Task]{}, err
	}

	return response.(*resources.Resources[*Task]), nil
}

// Get task detals by Id
func GetDetails(client newclient.Client, spaceID string, taskID string) (*TaskDetailsResource, error) {
	const detailTemplate = "/api/tasks{/id}/details"
	return newclient.GetByID[TaskDetailsResource](client, detailTemplate, spaceID, taskID)
}
