package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type FeedService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewFeedService(sling *sling.Sling, uriTemplate string) *FeedService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &FeedService{
		name:  "FeedService",
		path:  path,
		sling: sling,
	}
}

func (s *FeedService) Get(id string) (*model.Feed, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Feed), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

// GetAll returns all instances of a Feed.
func (s *FeedService) GetAll() ([]model.Feed, error) {

	feeds := new([]model.Feed)
	err := s.validateInternalState()

	if err != nil {
		return *feeds, err
	}

	_, err = apiGet(s.sling, feeds, s.path+"/all")

	return *feeds, err
}

// GetByName performs a lookup and returns the Feed with a matching name.
func (s *FeedService) GetByName(name string) (*model.Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("GetByName", "name")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new Feed.
func (s *FeedService) Add(feed model.Feed) (*model.Feed, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = feed.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, &feed, new(model.Feed), "feeds")

	if err != nil {
		return nil, err
	}

	return resp.(*model.Feed), nil
}

func (s *FeedService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
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
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &FeedService{}
