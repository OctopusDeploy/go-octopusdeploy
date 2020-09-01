package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectTriggerService struct {
	sling *sling.Sling
	path  string
}

func NewProjectTriggerService(sling *sling.Sling) *ProjectTriggerService {
	return &ProjectTriggerService{
		sling: sling,
		path:  "projecttriggers",
	}
}

func (s *ProjectTriggerService) Get(id string) (*model.ProjectTrigger, error) {
	path := fmt.Sprintf(s.path+"/%s", id)

	resp, err := apiGet(s.sling, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) GetByProjectID(id string) (*[]model.ProjectTrigger, error) {
	var triggersByProject []model.ProjectTrigger

	triggers, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	triggersByProject = append(triggersByProject, *triggers...)

	return &triggersByProject, nil
}

func (s *ProjectTriggerService) GetAll() (*[]model.ProjectTrigger, error) {
	var p []model.ProjectTrigger
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.ProjectTriggers), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.ProjectTriggers)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *ProjectTriggerService) Add(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	resp, err := apiAdd(s.sling, resource, new(model.ProjectTrigger), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}

func (s *ProjectTriggerService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectTriggerService) Update(resource *model.ProjectTrigger) (*model.ProjectTrigger, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.ProjectTrigger), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectTrigger), nil
}
