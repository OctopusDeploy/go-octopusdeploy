package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectService struct {
	sling *sling.Sling
	path  string
}

func NewProjectService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: sling,
		path:  "projects",
	}
}

// Get returns a single project by its ID in Octopus Deploy
func (s *ProjectService) Get(id string) (*model.Project, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// GetAll returns all projects in Octopus Deploy
func (s *ProjectService) GetAll() (*[]model.Project, error) {
	var p []model.Project
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Projects), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Projects)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName gets an existing project by its project name in Octopus Deploy
func (s *ProjectService) GetByName(name string) (*model.Project, error) {
	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

// Add adds an new project in Octopus Deploy
func (s *ProjectService) Add(resource *model.Project) (*model.Project, error) {
	err := model.ValidateProjectValues(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Project), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// Delete deletes an existing project in Octopus Deploy
func (s *ProjectService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing project in Octopus Deploy
func (s *ProjectService) Update(resource *model.Project) (*model.Project, error) {
	err := model.ValidateProjectValues(resource)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}
