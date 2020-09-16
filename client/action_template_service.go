package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/go-playground/validator"
)

// ActionTemplateService handles communication with ActionTemplate-related
// methods of the Octopus API.
type ActionTemplateService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewActionTemplateService returns an ActionTemplateService with a
// preconfigured client.
func NewActionTemplateService(sling *sling.Sling, uriTemplate string) *ActionTemplateService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ActionTemplateService{
		sling: sling,
		path:  path,
	}
}

// Get returns an ActionTemplate that matches the input ID.
func (s *ActionTemplateService) Get(id string) (*model.ActionTemplate, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("ActionTemplateService", "id")
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

	actionTemplates := new([]model.ActionTemplate)

	if err != nil {
		return *actionTemplates, err
	}

	_, err = apiGet(s.sling, actionTemplates, s.path+"/all")

	return *actionTemplates, err
}

// GetByName performs a lookup and returns the ActionTemplate with a matching name.
func (s *ActionTemplateService) GetByName(name string) (*model.ActionTemplate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("ActionTemplateService", "name")
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

	return nil, errors.New("client: item not found")
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
func (s *ActionTemplateService) Delete(actionTemplateID string) error {
	if isEmpty(actionTemplateID) {
		return createInvalidParameterError("ActionTemplateService", "actionTemplateID")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", actionTemplateID))
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
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ServiceInterface = &ActionTemplateService{}
