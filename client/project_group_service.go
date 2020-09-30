package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// projectGroupService handles communication with ProjectGroup-related methods of the Octopus API.
type projectGroupService struct {
	name        string                    `validate:"required"`
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

// Add creates a new project group.
func (s projectGroupService) Add(resource *model.ProjectGroup) (*model.ProjectGroup, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

// DeleteByID deletes the project group that matches the input ID.
func (s projectGroupService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("project group", "ID", id)
	}

	return err
}

// GetByID returns the project group that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s projectGroupService) GetByID(id string) (*model.ProjectGroup, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ProjectGroup), path)
	if err != nil {
		return nil, createResourceNotFoundError("project group", "ID", id)
	}

	return resp.(*model.ProjectGroup), nil
}

// GetAll returns all project groups. If none can be found or an error occurs,
// it returns an empty collection.
func (s projectGroupService) GetAll() ([]model.ProjectGroup, error) {
	items := []model.ProjectGroup{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// Update modifies a project group based on the one provided as input.
func (s projectGroupService) Update(resource model.ProjectGroup) (*model.ProjectGroup, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

var _ ServiceInterface = &projectGroupService{}
