package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type FeedService struct {
	sling *sling.Sling
	path  string
}

func NewFeedService(sling *sling.Sling) *FeedService {
	return &FeedService{
		sling: sling,
		path:  "feeds",
	}
}

func (s *FeedService) Get(id string) (*model.Feed, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("FeedService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Feed), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s *FeedService) GetAll() (*[]model.Feed, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Feed), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Feed), nil
}

func (s *FeedService) GetByName(name string) (*model.Feed, error) {
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

func (s *FeedService) Add(resource *model.Feed) (*model.Feed, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Feed), "feeds")

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s *FeedService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *FeedService) Update(resource *model.Feed) (*model.Feed, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Feed), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s *FeedService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("FeedService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("FeedService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &FeedService{}
