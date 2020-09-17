package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ConfigurationService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewConfigurationService(sling *sling.Sling, uriTemplate string) *ConfigurationService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ConfigurationService{
		name:  "ConfigurationService",
		path:  path,
		sling: sling,
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
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
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
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ConfigurationService{}
