package services

import (
	"github.com/dghubble/sling"
)

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

// Add creates a new task.
func (s taskService) Add(task *Task) (*Task, error) {
	path, err := getAddPath(s, task)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), task, new(Task), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Task), nil
}

// Get returns a collection of tasks based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s taskService) Get(tasksQuery TasksQuery) (*Tasks, error) {
	path, err := s.getURITemplate().Expand(tasksQuery)
	if err != nil {
		return &Tasks{}, err
	}

	response, err := apiGet(s.getClient(), new(Tasks), path)
	if err != nil {
		return &Tasks{}, err
	}

	return response.(*Tasks), nil
}
