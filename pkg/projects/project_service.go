package projects

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
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

const (
	projectsTemplate     = "/api/{spaceId}/projects{/id}{?name,skip,ids,clone,take,partialName,clonedFromProjectId}"
	convertToVcsTemplate = "/api/{spaceId}/projects/{id}/git/convert"
)

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
//
// Deprecated: Use projects.Add
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

func (s *ProjectService) Clone(sourceProject *Project, request ProjectCloneRequest) (*Project, error) {
	path, err := s.GetURITemplate().Expand(&ProjectCloneQuery{CloneProjectID: sourceProject.GetID()})
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiPost(s.GetClient(), request, new(Project), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Project), nil
}

// ConvertToVcs converts an input project to use a version-control system (VCS) for its persistence. initialCommitBranch is ignored unless
// the default branch in the gitPersistenceSettings appears in the protected branch patterns, and will default to "octopus-vcs-conversion"
// if not explicitly specified.
//
// Deprecated: Use projects.ConvertToVCS
func (s *ProjectService) ConvertToVcs(project *Project, commitMessage string, initialCommitBranch string, gitPersistenceSettings GitPersistenceSettings) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("ConvertToVcs", "project")
	}

	if gitPersistenceSettings == nil {
		return nil, fmt.Errorf("input parameter (versionControlSettings) is nil")
	}

	if project.Links == nil {
		return nil, fmt.Errorf("the state of the input project is not valid; links collection is empty")
	}

	if len(project.Links["ConvertToVcs"]) == 0 {
		return nil, fmt.Errorf("the state of the input project is not valid; cannot resolve ConvertToVcs link")
	}

	convertToVcs := NewConvertToVcs(commitMessage, initialCommitBranch, gitPersistenceSettings)

	_, err := services.ApiAddWithResponseStatus(s.GetClient(), convertToVcs, new(ConvertToVcsResponse), project.Links["ConvertToVcs"], http.StatusOK)
	if err != nil {
		return nil, err
	}

	return s.GetByID(project.GetID())
}

// Get returns a collection of projects based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
//
// Deprecated: Use projects.Get
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
//
// Deprecates: use projects.GetAll
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
//
// Deprecated: Use projects.GetByID
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

// Deprecated: Use projects.GetByName
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

// Deprecated: Use project.GetByIdentifier
func (p *ProjectService) GetByIdentifier(identifier string) (*Project, error) {
	project, err := p.GetByID(identifier)
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

	return p.GetByName(identifier)
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
//
// Deprecated: Use projects.Update
func (s *ProjectService) Update(project *Project) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterProject)
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

	if project.PersistenceSettings == nil || project.PersistenceSettings.Type() != PersistenceSettingsTypeVersionControlled {
		return s.Update(project)
	}

	if len(gitRef) == 0 {
		gitRef = project.PersistenceSettings.(GitPersistenceSettings).DefaultBranch()
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

func (s *ProjectService) GetProgression(project *Project) (*Progression, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("GetProgression", "project")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	url, err := url.Parse(project.Links[constants.LinkProgression])

	if err != nil {
		return nil, err
	}

	path := strings.Split(url.Path, "{")[0]
	resp, err := api.ApiGet(s.GetClient(), new(Progression), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Progression), nil
}

// ----- new -----

// Add creates a new project.
func Add(client newclient.Client, project *Project) (*Project, error) {
	if IsNil(project) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterProject)
	}

	// Remove persistence settings if specified; this will generate an error from
	// the endpoint if specified. Persistence settings are available AFTER the project
	// has been created or converted.
	project.PersistenceSettings = nil

	spaceID, err := internal.GetSpaceID(project.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(projectsTemplate, map[string]any{"spaceId": spaceID})
	if err != nil {
		return nil, err
	}

	projectResponse, err := newclient.Post[Project](client.HttpSession(), expandedUri, project)
	if err != nil {
		return nil, err
	}

	return projectResponse, nil
}

// Get returns a collection of projects based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func Get(client newclient.Client, spaceID string, projectsQuery ProjectsQuery) (*resources.Resources[*Project], error) {
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	values, _ := uritemplates.Struct2map(projectsQuery)
	if values == nil {
		values = map[string]any{}
	}
	values["spaceId"] = spaceID

	expandedUri, err := client.URITemplateCache().Expand(projectsTemplate, values)
	if err != nil {
		return nil, err
	}

	resp, err := newclient.Get[resources.Resources[*Project]](client.HttpSession(), expandedUri)
	if err != nil {
		return &resources.Resources[*Project]{}, err
	}

	return resp, nil
}

// GetByID returns the project that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*Project, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(projectsTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})

	resp, err := newclient.Get[Project](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ConvertToVCS converts an input project to use a version-control system (VCS) for its persistence. initialCommitBranch is ignored unless
// the default branch in the gitPersistenceSettings appears in the protected branch patterns, and will default to "octopus-vcs-conversion"
// if not explicitly specified.
func ConvertToVCS(client newclient.Client, project *Project, commitMessage string, initialCommitBranch string, gitPersistenceSettings GitPersistenceSettings) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError("ConvertToVcs", "project")
	}

	if gitPersistenceSettings == nil {
		return nil, fmt.Errorf("input parameter (versionControlSettings) is nil")
	}

	if project.Links == nil {
		return nil, fmt.Errorf("the state of the input project is not valid; links collection is empty")
	}

	if len(project.Links["ConvertToVcs"]) == 0 {
		return nil, fmt.Errorf("the state of the input project is not valid; cannot resolve ConvertToVcs link")
	}

	spaceID, err := internal.GetSpaceID(project.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	convertToVcs := NewConvertToVcs(commitMessage, initialCommitBranch, gitPersistenceSettings)

	expandedUri, err := client.URITemplateCache().Expand(convertToVcsTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      project.ID,
	})

	_, err = newclient.Post[ConvertToVcsResponse](client.HttpSession(), expandedUri, convertToVcs)
	if err != nil {
		return nil, err
	}

	return GetByID(client, spaceID, project.GetID())
}

// Update modifies a project based on the one provided as input.
func Update(client newclient.Client, project *Project) (*Project, error) {
	if project == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterProject)
	}

	spaceID, err := internal.GetSpaceID(project.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(projectsTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      project.ID,
	})
	if err != nil {
		return nil, err
	}

	project.Links = nil

	resp, err := newclient.Put[Project](client.HttpSession(), expandedUri, project)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteByID deletes the resource that matches the space ID and input ID.
func DeleteByID(client newclient.Client, spaceID string, id string) error {
	if internal.IsEmpty(id) {
		return internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID)
	}

	expandedUri, err := client.URITemplateCache().Expand(projectsTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), expandedUri)
}

// GetAll returns all projects. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*Project, error) {
	return newclient.GetAll[Project](client, projectsTemplate, spaceID)
}

func GetByName(client newclient.Client, spaceId string, name string) (*Project, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	projects, err := Get(client, client.GetSpaceID(), ProjectsQuery{
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

func GetByIdentifier(client newclient.Client, spaceId string, identifier string) (*Project, error) {
	project, err := GetByID(client, client.GetSpaceID(), identifier)
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

	return GetByName(client, client.GetSpaceID(), identifier)
}
