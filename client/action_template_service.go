package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// ActionTemplateService handles communication with ActionTemplate-related
// methods of the Octopus API.
type ActionTemplateService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewActionTemplateService returns an ActionTemplateService with a
// preconfigured client.
func NewActionTemplateService(sling *sling.Sling, uriTemplate string) *ActionTemplateService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ActionTemplateService{
		name:  "ActionTemplateService",
		path:  path,
		sling: sling,
	}
}

// Get returns an ActionTemplate that matches the input ID.
func (s *ActionTemplateService) Get(id string) (*model.ActionTemplate, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

// GetAll returns all instances of an ActionTemplate.
func (s *ActionTemplateService) GetAll() ([]model.ActionTemplate, error) {
	err := s.validateInternalState()

	items := new([]model.ActionTemplate)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the ActionTemplate with a matching name.
func (s *ActionTemplateService) GetByName(name string) (*model.ActionTemplate, error) {
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

// Add creates a new ActionTemplate.
func (s *ActionTemplateService) Add(actionTemplate *model.ActionTemplate) (*model.ActionTemplate, error) {
	if actionTemplate == nil {
		return nil, createInvalidParameterError("Add", "actionTemplate")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = actionTemplate.Validate()

	if err != nil {
		return nil, createValidationFailureError("Add", err)
	}

	resp, err := apiAdd(s.sling, actionTemplate, new(model.ActionTemplate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

// Delete removes the ActionTemplate that matches the input ID.
func (s *ActionTemplateService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ActionTemplateService) Update(actionTemplate model.ActionTemplate) (*model.ActionTemplate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = actionTemplate.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", actionTemplate.ID)
	resp, err := apiUpdate(s.sling, actionTemplate, new(model.ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

func (s *ActionTemplateService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ActionTemplateService{}
