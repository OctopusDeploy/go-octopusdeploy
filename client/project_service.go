package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewProjectService(sling *sling.Sling, uriTemplate string) *ProjectService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ProjectService{
		name:  "ProjectService",
		path:  path,
		sling: sling,
	}
}

// Get returns a single project by its ID in Octopus Deploy
func (s *ProjectService) Get(id string) (*model.Project, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// GetAll returns all instances of a Project.
func (s *ProjectService) GetAll() ([]model.Project, error) {
	err := s.validateInternalState()

	items := new([]model.Project)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Project with a matching name.
func (s *ProjectService) GetByName(name string) (*model.Project, error) {
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

func (s *ProjectService) GetChannels(project model.Project) ([]model.Channel, error) {
	channels := []model.Channel{}

	err := s.validateInternalState()

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
		resp, err := apiGet(s.sling, new(model.Channels), path)

		if err != nil {
			return channels, err
		}

		r := resp.(*model.Channels)
		channels = append(channels, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return channels, nil
}

func (s *ProjectService) GetSummary(project model.Project) (*model.ProjectSummary, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := project.Links["Summary"]
	resp, err := apiGet(s.sling, new(model.ProjectSummary), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectSummary), nil
}

func (s *ProjectService) GetReleases(project model.Project) ([]model.Release, error) {
	err := s.validateInternalState()

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
		resp, err := apiGet(s.sling, new(model.Releases), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Releases)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// Add creates a new Project.
func (s *ProjectService) Add(project *model.Project) (*model.Project, error) {
	if project == nil {
		return nil, createInvalidParameterError("Add", "project")
	}

	err := project.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, project, new(model.Project), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

// Delete deletes an existing project in Octopus Deploy
func (s *ProjectService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update updates an existing project in Octopus Deploy
func (s *ProjectService) Update(resource *model.Project) (*model.Project, error) {
	if resource == nil {
		return nil, createInvalidParameterError("Update", "resource")
	}

	err := resource.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Project), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Project), nil
}

func (s *ProjectService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ProjectService{}
