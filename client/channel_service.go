package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ChannelService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewChannelService(sling *sling.Sling, uriTemplate string) *ChannelService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ChannelService{
		name:  "ChannelService",
		path:  path,
		sling: sling,
	}
}

func (s *ChannelService) Get(id string) (*model.Channel, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// GetAll returns all instances of a Channel.
func (s *ChannelService) GetAll() ([]model.Channel, error) {
	err := s.validateInternalState()

	items := new([]model.Channel)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

func (s *ChannelService) GetProject(channel model.Channel) (*model.Project, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := channel.Links["Project"]
	resp, err := apiGet(s.sling, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

func (s *ChannelService) GetReleases(channel model.Channel) ([]model.Release, error) {
	releases := []model.Release{}

	err := s.validateInternalState()

	if err != nil {
		return releases, err
	}

	url, err := url.Parse(channel.Links["Releases"])

	if err != nil {
		return releases, err
	}

	path := strings.Split(url.Path, "{")[0]
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Releases), path)

		if err != nil {
			return releases, err
		}

		r := resp.(*model.Releases)
		releases = append(releases, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return releases, nil
}

// Add creates a new Channel.
func (s *ChannelService) Add(channel *model.Channel) (*model.Channel, error) {
	if channel == nil {
		return nil, createInvalidParameterError("Add", "channel")
	}

	err := channel.Validate()

	if err != nil {
		return nil, createValidationFailureError("Add", err)
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, channel, new(model.Channel), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// Delete removes the Channel that matches the input ID.
func (s *ChannelService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Channel based on the one provided as input.
func (s *ChannelService) Update(channel model.Channel) (*model.Channel, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = channel.Validate()

	if err != nil {
		return nil, createValidationFailureError("Update", err)
	}

	path := fmt.Sprintf(s.path+"/%s", channel.ID)
	resp, err := apiUpdate(s.sling, channel, new(model.Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

func (s *ChannelService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ChannelService{}
