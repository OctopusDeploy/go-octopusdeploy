package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ActionTemplateService struct {
	sling *sling.Sling
	path  string
}

func NewActionTemplateService(sling *sling.Sling) *ActionTemplateService {
	if sling == nil {
		fmt.Println(fmt.Errorf("ActionTemplateService: input parameter (sling) is nil"))
		return nil
	}

	return &ActionTemplateService{
		sling: sling,
		path:  "actiontemplates",
	}
}

func (s *ActionTemplateService) Get(id string) (*model.ActionTemplate, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("ActionTemplateService: invalid parameter, ID")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

func (s *ActionTemplateService) GetAll() (*[]model.ActionTemplate, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.ActionTemplate), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.ActionTemplate), nil
}

func (s *ActionTemplateService) GetByName(name string) (*model.ActionTemplate, error) {
	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if *item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

func (s *ActionTemplateService) Add(resource *model.ActionTemplate) (*model.ActionTemplate, error) {
	resp, err := apiAdd(s.sling, resource, new(model.ActionTemplate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

// Delete removes the ActionTemplate that matches the input ID.
func (s *ActionTemplateService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ActionTemplateService) Update(resource model.ActionTemplate) (*model.ActionTemplate, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.ActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

func (s *ActionTemplateService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("ActionTemplateService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("ActionTemplateService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &ActionTemplateService{}
