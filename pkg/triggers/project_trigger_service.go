package triggers

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
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
//
// Deprecated: use triggers.GetProjectTriggerByID
func (s *ProjectTriggerService) GetByID(id string) (*ProjectTrigger, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)

	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(ProjectTrigger), path)
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
	if projectTrigger == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationDelete, constants.ParameterProjectTrigger)
	}

	path := fmt.Sprintf("/api/%s/projecttriggers", projectTrigger.SpaceID)

	// TODO: use this updated path (below) once new path resides in production
	// path := fmt.Sprintf("/api/%s/projects/%s/triggers", projectTrigger.SpaceID, projectTrigger.ProjectID)

	resp, err := services.ApiAdd(s.GetClient(), projectTrigger, new(ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

// Delete will delete a project trigger.
func (s *ProjectTriggerService) Delete(projectTrigger *ProjectTrigger) error {
	if projectTrigger == nil {
		return internal.CreateInvalidParameterError(constants.OperationDelete, constants.ParameterProjectTrigger)
	}

	path := fmt.Sprintf("/api/%s/projects/%s/triggers/%s", projectTrigger.SpaceID, projectTrigger.ProjectID, projectTrigger.GetID())
	return services.ApiDelete(s.GetClient(), path)
}

// Update modifies a project trigger based on the one provided as input.
func (s *ProjectTriggerService) Update(projectTrigger *ProjectTrigger) (*ProjectTrigger, error) {
	if projectTrigger == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterProjectTrigger)
	}

	path := fmt.Sprintf("/api/%s/projects/%s/triggers/%s", projectTrigger.SpaceID, projectTrigger.ProjectID, projectTrigger.GetID())
	resp, err := services.ApiUpdate(s.GetClient(), projectTrigger, new(ProjectTrigger), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectTrigger), nil
}

// ----- Experimental --------

const (
	projectTriggersTemplate = "/api/{spaceId}/projecttriggers/{id}"
)

// GetProjectTriggerByID returns the project trigger that matches the input ID.
func (s *ProjectTriggerService) GetProjectTriggerByID(client newclient.Client, spaceId string, id string) (*ProjectTrigger, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	return newclient.GetByID[ProjectTrigger](client, projectTriggersTemplate, spaceId, id)
}
