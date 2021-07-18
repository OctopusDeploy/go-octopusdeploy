package octopusdeploy

import (
	"net/url"
	"strings"

	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type projectService struct {
	experimentalSummariesPath string
	pulsePath                 string

	canDeleteService
}

func newProjectService(sling *sling.Sling, uriTemplate string, pulsePath string, experimentalSummariesPath string) *projectService {
	projectService := &projectService{
		experimentalSummariesPath: experimentalSummariesPath,
		pulsePath:                 pulsePath,
	}
	projectService.service = newService(ServiceProjectService, sling, uriTemplate)

	return projectService
}

// Get returns a collection of projects based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s projectService) Get(projectsQuery ProjectsQuery) (*Projects, error) {
	v, _ := query.Values(projectsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := apiGet(s.getClient(), new(Projects), path)
	if err != nil {
		return &Projects{}, err
	}

	return resp.(*Projects), nil
}

// GetAll returns all projects. If none can be found or an error occurs, it
// returns an empty collection.
func (s projectService) GetAll() ([]*Project, error) {
	items := []*Project{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the project that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s projectService) GetByID(id string) (*Project, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// GetByName performs a lookup and returns the Project with a matching name.
func (s projectService) GetByName(name string) (*Project, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(OperationGetByName, ParameterName)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return item, nil
		}
	}

	return nil, createItemNotFoundError(s.getName(), OperationGetByName, name)
}

func (s projectService) GetChannels(project *Project) ([]*Channel, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationGetChannels, ParameterProject)
	}

	channels := []*Channel{}

	err := validateInternalState(s)

	if err != nil {
		return channels, err
	}

	url, err := url.Parse(project.Links[linkChannels])

	if err != nil {
		return channels, err
	}

	path := strings.Split(url.Path, "{")[0]
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Channels), path)

		if err != nil {
			return channels, err
		}

		r := resp.(*Channels)
		channels = append(channels, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return channels, nil
}

func (s projectService) GetSummary(project *Project) (*ProjectSummary, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationGetSummary, ParameterProject)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	path := project.Links[linkSummary]
	resp, err := apiGet(s.getClient(), new(ProjectSummary), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectSummary), nil
}

func (s projectService) GetReleases(project *Project) ([]*Release, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationGetReleases, ParameterProject)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	url, err := url.Parse(project.Links[linkReleases])

	if err != nil {
		return nil, err
	}

	path := strings.Split(url.Path, "{")[0]

	p := []*Release{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Releases), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*Releases)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// Add creates a new project.
func (s projectService) Add(resource *Project) (*Project, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// Update modifies a project based on the one provided as input.
func (s projectService) Update(project *Project) (*Project, error) {
	path, err := getUpdatePath(s, project)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}
