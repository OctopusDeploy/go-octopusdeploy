package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type FeedService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewFeedService(sling *sling.Sling) *FeedService {
	if sling == nil {
		return nil
	}

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
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(name, " ")) == 0 {
		return nil, errors.New("FeedService: invalid parameter, name")
	}

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

func (s *FeedService) Add(feed *model.Feed) (*model.Feed, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if feed == nil {
		return nil, errors.New("FeedService: invalid parameter, feed")
	}

	resp, err := apiAdd(s.sling, feed, new(model.Feed), "feeds")

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s *FeedService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return errors.New("FeedService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *FeedService) Update(feed model.Feed) (*model.Feed, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	err = feed.Validate()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", feed.ID)
	resp, err := apiUpdate(s.sling, feed, new(model.Feed), path)

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
