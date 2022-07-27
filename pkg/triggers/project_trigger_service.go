package triggers

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	services.CanDeleteService
}

func NewProjectTriggerService(sling *sling.Sling, uriTemplate string) *ProjectTriggerService {
	return &ProjectTriggerService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceProjectTriggerService, sling, uriTemplate),
		},
	}
}

// GetByID returns the project trigger that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s *ProjectTriggerService) GetByID(id string) (*ProjectTrigger, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(id string) ([]*ProjectTrigger, error) {
	var triggersByProject []*ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, triggers...)

	return triggersByProject, nil
}

// GetAll returns all project triggers. If none can be found or an error
// occurs, it returns an empty collection.
func (s *ProjectTriggerService) GetAll() ([]*ProjectTrigger, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return []*ProjectTrigger{}, err
	}

	return services.GetPagedResponse[ProjectTrigger](s, path)
}

// Add creates a new project trigger.
func (s *ProjectTriggerService) Add(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	if IsNil(projectTrigger) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterProjectTrigger)
	}

	path, err := services.GetAddPath(s, projectTrigger)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), projectTrigger, new(ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

// Update modifies a project trigger based on the one provided as input.
func (s *ProjectTriggerService) Update(resource ProjectTrigger) (*ProjectTrigger, error) {
	path, err := services.GetUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}
