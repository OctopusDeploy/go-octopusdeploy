package client

import (
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type projectService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newProjectService(sling *sling.Sling, uriTemplate string) *projectService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &projectService{
		name:        serviceProjectService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s projectService) getClient() *sling.Sling {
	return s.sling
}

func (s projectService) getName() string {
	return s.name
}

func (s projectService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns the project that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s projectService) GetByID(id string) (*model.Project, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Project), path)
	if err != nil {
		return nil, createResourceNotFoundError("project", "ID", id)
	}

	return resp.(*model.Project), nil
}

// GetAll returns all projects. If none can be found or an error occurs, it
// returns an empty collection.
func (s projectService) GetAll() ([]model.Project, error) {
	items := []model.Project{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Project with a matching name.
func (s projectService) GetByName(name string) (*model.Project, error) {
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

func (s projectService) GetChannels(project model.Project) ([]model.Channel, error) {
	channels := []model.Channel{}

	err := validateInternalState(s)

	if err != nil {
		return channels, err
	}

	url, err := url.Parse(project.Links["Channels"])

	if err != nil {
		return channels, err
	}

	path := strings.Split(url.Path, "{")[0]
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Channels), path)

		if err != nil {
			return channels, err
		}

		r := resp.(*model.Channels)
		channels = append(channels, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return channels, nil
}

func (s projectService) GetSummary(project model.Project) (*model.ProjectSummary, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := project.Links["Summary"]
	resp, err := apiGet(s.getClient(), new(model.ProjectSummary), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectSummary), nil
}

func (s projectService) GetReleases(project model.Project) ([]model.Release, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(project.Links["Releases"])

	if err != nil {
		return nil, err
	}

	path := strings.Split(url.Path, "{")[0]

	p := []model.Release{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Releases), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*model.Releases)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// Add creates a new project.
func (s projectService) Add(resource *model.Project) (*model.Project, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// DeleteByID deletes the project that matches the input ID.
func (s projectService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("project", "ID", id)
	}

	return err
}

// Update modifies a project based on the one provided as input.
func (s projectService) Update(resource model.Project) (*model.Project, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

var _ ServiceInterface = &projectService{}
