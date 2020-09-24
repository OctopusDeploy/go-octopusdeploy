package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// projectGroupService handles communication with ProjectGroup-related methods of the Octopus API.
type projectGroupService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newProjectGroupService returns a projectGroupService with a preconfigured client.
func newProjectGroupService(sling *sling.Sling, uriTemplate string) *projectGroupService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &projectGroupService{
		name:        serviceProjectGroupService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s projectGroupService) getClient() *sling.Sling {
	return s.sling
}

func (s projectGroupService) getName() string {
	return s.name
}

func (s projectGroupService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a ProjectGroup that matches the input ID. If one cannot be found, it returns nil and an error.
func (s projectGroupService) GetByID(id string) (*model.ProjectGroup, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

// GetAll returns all instances of a ProjectGroup. If none can be found or an error occurs, it returns an empty collection.
func (s projectGroupService) GetAll() ([]model.ProjectGroup, error) {
	items := new([]model.ProjectGroup)
	path, err := getAllPath(s)
	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.getClient(), items, path)
	return *items, err
}

// Add creates a new ProjectGroup.
func (s projectGroupService) Add(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	if projectGroup == nil {
		return nil, createInvalidParameterError(operationAdd, "projectGroup")
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = projectGroup.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), projectGroup, new(model.ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s projectGroupService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

func (s projectGroupService) Update(projectGroup *model.ProjectGroup) (*model.ProjectGroup, error) {
	if projectGroup == nil {
		return nil, createInvalidParameterError(operationUpdate, "projectGroup")
	}

	err := projectGroup.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", projectGroup.ID)

	resp, err := apiUpdate(s.getClient(), projectGroup, new(model.ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

var _ ServiceInterface = &projectGroupService{}
