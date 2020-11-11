package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// projectGroupService handles communication with ProjectGroup-related methods of the Octopus API.
type projectGroupService struct {
	canDeleteService
}

// newProjectGroupService returns a projectGroupService with a preconfigured client.
func newProjectGroupService(sling *sling.Sling, uriTemplate string) *projectGroupService {
	projectGroupService := &projectGroupService{}
	projectGroupService.service = newService(ServiceProjectGroupService, sling, uriTemplate)

	return projectGroupService
}

func (s projectGroupService) getPagedResponse(path string) ([]*ProjectGroup, error) {
	resources := []*ProjectGroup{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(ProjectGroups), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ProjectGroups)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new project group.
func (s projectGroupService) Add(resource *ProjectGroup) (*ProjectGroup, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}

// Get returns a collection of project groups based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s projectGroupService) Get(projectGroupsQuery ProjectGroupsQuery) (*ProjectGroups, error) {
	path, err := s.getURITemplate().Expand(projectGroupsQuery)
	if err != nil {
		return &ProjectGroups{}, err
	}

	response, err := apiGet(s.getClient(), new(ProjectGroups), path)
	if err != nil {
		return &ProjectGroups{}, err
	}

	return response.(*ProjectGroups), nil
}

// GetAll returns all project groups. If none can be found or an error occurs,
// it returns an empty collection.
func (s projectGroupService) GetAll() ([]*ProjectGroup, error) {
	items := []*ProjectGroup{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the project group that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s projectGroupService) GetByID(id string) (*ProjectGroup, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ProjectGroup), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*ProjectGroup), nil
}

// GetByPartialName performs a lookup and returns a collection of project
// groups with a matching partial name.
func (s projectGroupService) GetByPartialName(name string) ([]*ProjectGroup, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*ProjectGroup{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a project group based on the one provided as input.
func (s projectGroupService) Update(resource ProjectGroup) (*ProjectGroup, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(ProjectGroup), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ProjectGroup), nil
}
