package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type EnvironmentService struct {
	sling *sling.Sling
	path  string
}

func NewEnvironmentService(sling *sling.Sling) *EnvironmentService {
	return &EnvironmentService{
		sling: sling,
		path:  "environments",
	}
}

func (s *EnvironmentService) Get(id string) (*model.Environment, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) GetAll() (*[]model.Environment, error) {
	var p []model.Environment
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Environments), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Environments)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *EnvironmentService) GetByName(name string) (*model.Environment, error) {
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

func (s *EnvironmentService) Add(resource *model.Environment) (*model.Environment, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Environment), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s *EnvironmentService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *EnvironmentService) Update(resource *model.Environment) (*model.Environment, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Environment), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}
