package projects

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"net/http"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type ProjectService struct {
	experimentalSummariesPath string
	exportProjectsPath        string
	importProjectsPath        string
	pulsePath                 string

	services.CanDeleteService
}

func NewProjectService(sling *sling.Sling, uriTemplate string, pulsePath string, experimentalSummariesPath string, importProjectsPath string, exportProjectsPath string) *ProjectService {
	return &ProjectService{
		experimentalSummariesPath: experimentalSummariesPath,
		exportProjectsPath:        exportProjectsPath,
		importProjectsPath:        importProjectsPath,
		pulsePath:                 pulsePath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceProjectService, sling, uriTemplate),
		},
	}
}

// Add creates a new project.
func (s *ProjectService) Add(project *Project) (*Project, error) {
	if IsNil(project) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterProject)
	}

	path, err := services.GetAddPath(s, project)
	if err != nil {
		return nil, err
	}

	// Remove persistence settings if specified; this will generate an error from
	// the endpoint if specified. Persistence settings are available AFTER the project
	// has been created or converted.
	project.PersistenceSettings = nil

	resp, err := services.ApiAdd(s.GetClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// ConvertToVcs converts an input project to use a version-control system (VCS) for its persistence.
func (s *ProjectService) ConvertToVcs(project *Project, commitMessage string, versionControlSettings *VersionControlSettings) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("ConvertToVcs", "project")
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

	convertToVcs := NewConvertToVcs(commitMessage, versionControlSettings)
	_, err := services.ApiAddWithResponseStatus(s.GetClient(), convertToVcs, new(ConvertToVcsResponse), project.Links["ConvertToVcs"], http.StatusOK)
	if err != nil {
		return nil, err
	}

	return s.GetByID(project.GetID())
}

// Get returns a collection of projects based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *ProjectService) Get(projectsQuery ProjectsQuery) (*resources.Resources[*Project], error) {
	v, _ := query.Values(projectsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Project]), path)
	if err != nil {
		return &resources.Resources[*Project]{}, err
	}

	return resp.(*resources.Resources[*Project]), nil
}

// GetAll returns all projects. If none can be found or an error occurs, it
// returns an empty collection.
func (s *ProjectService) GetAll() ([]*Project, error) {
	items := []*Project{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the project that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *ProjectService) GetByID(id string) (*Project, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

func (p *ProjectService) GetByName(name string) (*Project, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	projects, err := p.Get(ProjectsQuery{
		PartialName: name,
	})
	if err != nil {
		return nil, err
	}

	for _, project := range projects.Items {
		if strings.EqualFold(project.Name, name) {
			return project, nil
		}
	}

	return nil, services.ErrItemNotFound
}

func (p *ProjectService) GetByIdOrName(idOrName string) (*Project, error) {
	project, err := p.GetByID(idOrName)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if project != nil {
			return project, nil
		}
	}

	return p.GetByName(idOrName)
}

func (s *ProjectService) GetProject(channel *channels.Channel) (*Project, error) {
	if channel == nil {
		return nil, internal.CreateInvalidParameterError("GetProject", "channel")
	}

	path := channel.GetLinks()[constants.LinkProjects]
	resp, err := api.ApiGet(s.GetClient(), new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

func (s *ProjectService) GetChannels(project *Project) ([]*channels.Channel, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetChannels", "project")
	}

	projectChannels := []*channels.Channel{}

	if err := services.ValidateInternalState(s); err != nil {
		return projectChannels, err
	}

	channelsUrl, err := url.Parse(project.Links[constants.LinkChannels])

	if err != nil {
		return projectChannels, err
	}

	path := strings.Split(channelsUrl.Path, "{")[0]
	loadNextPage := true

	for loadNextPage {
		resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*channels.Channel]), path)

		if err != nil {
			return projectChannels, err
		}

		r := resp.(*resources.Resources[*channels.Channel])
		projectChannels = append(projectChannels, r.Items...)
		path, loadNextPage = services.LoadNextPage(r.PagedResults)
	}

	return projectChannels, nil
}

func (s *ProjectService) GetSummary(project *Project) (*ProjectSummary, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetSummary, constants.ParameterProject)
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	path := project.Links[constants.LinkSummary]
	resp, err := api.ApiGet(s.GetClient(), new(ProjectSummary), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectSummary), nil
}

func (s *ProjectService) GetReleases(project *Project) ([]*releases.Release, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetReleases", "project")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	releasesUrl, err := url.Parse(project.Links[constants.LinkReleases])

	if err != nil {
		return nil, err
	}

	path := strings.Split(releasesUrl.Path, "{")[0]

	result := make([]*releases.Release, 0, 4)
	loadNextPage := true

	for loadNextPage {
		resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*releases.Release]), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*resources.Resources[*releases.Release])
		result = append(result, r.Items...)
		path, loadNextPage = services.LoadNextPage(r.PagedResults)
	}

	return result, nil
}

// Update modifies a project based on the one provided as input.
func (s *ProjectService) Update(project *Project) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterProject)
	}

	if project.PersistenceSettings != nil && project.PersistenceSettings.GetType() == "VersionControlled" {
		defaultBranch := project.PersistenceSettings.(*GitPersistenceSettings).DefaultBranch
		return s.UpdateWithGitRef(project, defaultBranch)
	}

	path, err := services.GetUpdatePath(s, project)
	if err != nil {
		return nil, err
	}

	project.Links = nil
	resp, err := services.ApiUpdate(s.GetClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// Update modifies a Git-based project based on the one provided as input.
func (s *ProjectService) UpdateWithGitRef(project *Project, gitRef string) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("UpdateWithGitRef", "project")
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
	resp, err := services.ApiUpdate(s.GetClient(), project, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}
