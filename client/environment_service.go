package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type environmentService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newEnvironmentService(sling *sling.Sling, uriTemplate string) *environmentService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &environmentService{
		name:        serviceEnvironmentService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s environmentService) getClient() *sling.Sling {
	return s.sling
}

func (s environmentService) getName() string {
	return s.name
}

func (s environmentService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

func (s environmentService) GetByID(id string) (*model.Environment, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

// GetAll returns all instances of an Environment. If none can be found or an error occurs, it returns an empty collection.
func (s environmentService) GetAll() ([]model.Environment, error) {
	items := new([]model.Environment)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// GetByName performs a lookup and returns the Environment with a matching name.
func (s environmentService) GetByName(name string) (*model.Environment, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
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

	return nil, createItemNotFoundError(s.name, operationGetByName, name)
}

// Add creates a new Environment.
func (s environmentService) Add(environment *model.Environment) (*model.Environment, error) {
	if environment == nil {
		return nil, createInvalidParameterError(operationAdd, parameterEnvironment)
	}

	err := environment.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), environment, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

func (s environmentService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

func (s environmentService) Update(environment *model.Environment) (*model.Environment, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = environment.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", environment.ID)

	resp, err := apiUpdate(s.getClient(), environment, new(model.Environment), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Environment), nil
}

var _ ServiceInterface = &environmentService{}
