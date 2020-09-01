package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type UserService struct {
	sling *sling.Sling
	path  string
}

func NewUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling,
		path:  "users",
	}
}

func (s *UserService) Get(id string) (*model.User, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) GetAll() (*[]model.User, error) {
	var p []model.User
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Users), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Users)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *UserService) GetByName(name string) (*model.User, error) {
	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Username == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

func (s *UserService) Add(resource *model.User) (*model.User, error) {
	resp, err := apiAdd(s.sling, resource, new(model.User), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *UserService) Update(resource *model.User) (*model.User, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}
