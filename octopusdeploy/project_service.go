package octopusdeploy

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type projectService struct {
	experimentalSummariesPath string
	exportProjectsPath        string
	importProjectsPath        string
	pulsePath                 string

	canDeleteService
}

func newProjectService(sling *sling.Sling, uriTemplate string, pulsePath string, experimentalSummariesPath string, importProjectsPath string, exportProjectsPath string) *projectService {
	projectService := &projectService{
		experimentalSummariesPath: experimentalSummariesPath,
		exportProjectsPath:        exportProjectsPath,
		importProjectsPath:        importProjectsPath,
		pulsePath:                 pulsePath,
	}
	projectService.service = newService(ServiceProjectService, sling, uriTemplate)

	return projectService
}

// Add creates a new project.
func (s projectService) Add(project *Project) (*Project, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterProject)
	}

	path, err := getAddPath(s, project)
	if err != nil {
		return nil, err
	}

	// Remove persistence settings if specified; this will generate an error from
	// the endpoint if specified. Persistence settings are available AFTER the project
	// has been created or converted.
	project.PersistenceSettings = nil

	resp, err := apiAdd(s.getClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// ConvertToVcs converts an input project to use a version-control system (VCS) for its persistence.
func (s projectService) ConvertToVcs(project *Project, versionControlSettings *VersionControlSettings) (*Project, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterProject)
	}

	if versionControlSettings == nil {
		return nil, fmt.Errorf("input parameter (versionControlSettings) is nil")
	}

	if project.Links == nil {
		return nil, fmt.Errorf("the state of the input project is not valid; links collection is empty")
	}

	if len(project.Links["ConvertToVcs"]) == 0 {
		return nil, fmt.Errorf("the state of the input project is not valid; cannot resolve ConvertToVcs link")
	}

	convertToVcs := NewConvertToVcs(project.Name, versionControlSettings)
	_, err := apiAddWithResponseStatus(s.getClient(), convertToVcs, new(ConvertToVcsResponse), project.Links["ConvertToVcs"], http.StatusOK)
	if err != nil {
		return nil, err
	}

	return s.GetByID(project.GetID())
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

// Update modifies a project based on the one provided as input.
func (s projectService) Update(project *Project) (*Project, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterProject)
	}

	if project.PersistenceSettings != nil && project.PersistenceSettings.GetType() == "VersionControlled" {
		defaultBranch := project.PersistenceSettings.(*GitPersistenceSettings).DefaultBranch
		return s.UpdateWithGitRef(project, defaultBranch)
	}

	path, err := getUpdatePath(s, project)
	if err != nil {
		return nil, err
	}

	project.Links = nil
	resp, err := apiUpdate(s.getClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// Update modifies a Git-based project based on the one provided as input.
func (s projectService) UpdateWithGitRef(project *Project, gitRef string) (*Project, error) {
	if project == nil {
		return nil, createInvalidParameterError(OperationGet, ParameterProject)
	}

	if project.PersistenceSettings == nil || project.PersistenceSettings.GetType() != "VersionControlled" {
		return s.Update(project)
	}

	if len(gitRef) == 0 {
		gitRef = project.PersistenceSettings.(*GitPersistenceSettings).DefaultBranch
	}

	if len(gitRef) == 0 {
		return nil, fmt.Errorf("the gitRef is empty")
	}

	template, _ := uritemplates.Parse(project.Links["Self"])
	path, _ := template.Expand(map[string]interface{}{"gitRef": gitRef})

	project.Links = nil
	resp, err := apiUpdate(s.getClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}
