package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ChannelService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewChannelService(sling *sling.Sling) *ChannelService {
	if sling == nil {
		return nil
	}

	return &ChannelService{
		sling: sling,
		path:  "channels",
	}
}

func (s *ChannelService) Get(id string) (*model.Channel, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("ChannelService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// GetAll returns all Channels.
func (s *ChannelService) GetAll() (*[]model.Channel, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Channel), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Channel), nil
}

// Add creates a new Channel.
func (s *ChannelService) Add(resource *model.Channel) (*model.Channel, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("ChannelService: invalid parameter, resource")
	}

	err = resource.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Channel), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// Delete removes the Channel that matches the input ID.
func (s *ChannelService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return errors.New("ChannelService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Channel based on the one provided as input.
func (s *ChannelService) Update(channel *model.Channel) (*model.Channel, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	err = channel.Validate()
	if err != nil {
		return nil, err
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
		return fmt.Errorf("ChannelService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ChannelService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ChannelService{}
