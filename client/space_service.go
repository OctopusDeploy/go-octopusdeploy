package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type SpaceService struct {
	sling *sling.Sling
	path  string
}

func NewSpaceService(sling *sling.Sling) *SpaceService {
	return &SpaceService{
		sling: sling,
		path:  "spaces",
	}
}

func (s *SpaceService) Get(id string) (*model.Space, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

func (s *SpaceService) GetAll() (*[]model.Space, error) {
	var p []model.Space
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Spaces), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Spaces)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *SpaceService) GetByName(name string) (*model.Space, error) {
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

func (s *SpaceService) Add(resource *model.Space) (*model.Space, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Space), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}

func (s *SpaceService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *SpaceService) Update(resource *model.Space) (*model.Space, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Space), nil
}
