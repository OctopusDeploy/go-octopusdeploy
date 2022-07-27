package projectgroups

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// ProjectGroupService handles communication with ProjectGroup-related methods of the Octopus API.
type ProjectGroupService struct {
	services.CanDeleteService
}

// NewProjectGroupService returns a projectGroupService with a preconfigured client.
func NewProjectGroupService(sling *sling.Sling, uriTemplate string) *ProjectGroupService {
	return &ProjectGroupService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceProjectGroupService, sling, uriTemplate),
		},
	}
}

// Add creates a new project group.
func (s *ProjectGroupService) Add(projectGroup *ProjectGroup) (*ProjectGroup, error) {
	if IsNil(projectGroup) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterProjectGroup)
	}

	path, err := services.GetAddPath(s, projectGroup)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), projectGroup, new(ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}

// Get returns a collection of project groups based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s *ProjectGroupService) Get(projectGroupsQuery ProjectGroupsQuery) (*resources.Resources[*ProjectGroup], error) {
	path, err := s.GetURITemplate().Expand(projectGroupsQuery)
	if err != nil {
		return &resources.Resources[*ProjectGroup]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*ProjectGroup]), path)
	if err != nil {
		return &resources.Resources[*ProjectGroup]{}, err
	}

	return response.(*resources.Resources[*ProjectGroup]), nil
}

// GetAll returns all project groups. If none can be found or an error occurs,
// it returns an empty collection.
func (s *ProjectGroupService) GetAll() ([]*ProjectGroup, error) {
	items := []*ProjectGroup{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the project group that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s *ProjectGroupService) GetByID(id string) (*ProjectGroup, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}

// GetByPartialName performs a lookup and returns a collection of project
// groups with a matching partial name.
func (s *ProjectGroupService) GetByPartialName(partialName string) ([]*ProjectGroup, error) {
	if internal.IsEmpty(partialName) {
		return []*ProjectGroup{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*ProjectGroup{}, err
	}

	return services.GetPagedResponse[ProjectGroup](s, path)
}

func (s *ProjectGroupService) GetProjects(projectGroup *ProjectGroup) ([]*projects.Project, error) {
	projectsToReturn := []*projects.Project{}

	if projectGroup == nil {
		return projectsToReturn, internal.CreateInvalidParameterError(constants.OperationGetProjects, constants.ParameterProjectGroup)
	}

	path := projectGroup.Links[constants.LinkProjects]

	loadNextPage := true

	for loadNextPage {
		resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*projects.Project]), path)
		if err != nil {
			return projectsToReturn, err
		}

		projectList := resp.(*resources.Resources[*projects.Project])
		projectsToReturn = append(projectsToReturn, projectList.Items...)
		path, loadNextPage = services.LoadNextPage(projectList.PagedResults)
	}

	return projectsToReturn, nil
}

// Update modifies a project group based on the one provided as input.
func (s *ProjectGroupService) Update(resource ProjectGroup) (*ProjectGroup, error) {
	path, err := services.GetUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}
