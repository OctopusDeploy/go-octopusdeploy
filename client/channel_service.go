package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ChannelService struct {
	sling *sling.Sling
	path  string
}

func NewChannelService(sling *sling.Sling) *ChannelService {
	return &ChannelService{
		sling: sling,
		path:  "channels",
	}
}

func (s *ChannelService) Get(id string) (*model.Channel, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// GetAll returns all Channels.
func (s *ChannelService) GetAll() (*[]model.Channel, error) {
	var p []model.Channel
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Channels), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Channels)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// Add creates a new Channel.
func (s *ChannelService) Add(resource *model.Channel) (*model.Channel, error) {
	err := model.ValidateChannelValues(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Channel), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}

// Delete removes the Account that matches the input ID.
func (s *ChannelService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Channel based on the one provided as input.
func (s *ChannelService) Update(resource *model.Channel) (*model.Channel, error) {
	err := model.ValidateChannelValues(resource)

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Channel), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Channel), nil
}
