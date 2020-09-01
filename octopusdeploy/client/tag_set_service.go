package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type TagSetService struct {
	sling *sling.Sling
	path  string
}

func NewTagSetService(sling *sling.Sling) *TagSetService {
	return &TagSetService{
		sling: sling,
		path:  "tagsets",
	}
}

func (s *TagSetService) Get(id string) (*model.TagSet, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) GetAll() (*[]model.TagSet, error) {
	var p []model.TagSet
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.TagSets), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.TagSets)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *TagSetService) GetByName(name string) (*model.TagSet, error) {
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

func (s *TagSetService) Add(resource *model.TagSet) (*model.TagSet, error) {
	resp, err := apiAdd(s.sling, resource, new(model.TagSet), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}

func (s *TagSetService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *TagSetService) Update(resource *model.TagSet) (*model.TagSet, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.TagSet), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.TagSet), nil
}
