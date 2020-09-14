package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ConfigurationService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewConfigurationService(sling *sling.Sling) *ConfigurationService {
	if sling == nil {
		return nil
	}

	return &ConfigurationService{
		sling: sling,
		path:  "configuration",
	}
}

// GetAll returns all instances of a ConfigurationSections.
func (s *ConfigurationService) GetAll() (*model.ConfigurationSections, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new(model.ConfigurationSections), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ConfigurationSections), nil
}

// Get returns a ConfigurationSection that matches the input ID.
func (s *ConfigurationService) Get(id string) (*model.ConfigurationSection, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("ConfigurationService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ConfigurationSection), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ConfigurationSection), nil
}

func (s *ConfigurationService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("ConfigurationService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ConfigurationService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ConfigurationService{}
