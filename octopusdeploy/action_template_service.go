package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// actionTemplateService handles communication for any operations in the
// Octopus API that pertain to action templates.
type actionTemplateService struct {
	categoriesPath    string
	logoPath          string
	searchPath        string
	versionedLogoPath string

	canDeleteService
}

// newActionTemplateService returns an actionTemplateService with a
// preconfigured client.
func newActionTemplateService(sling *sling.Sling, uriTemplate string, categoriesPath string, logoPath string, searchPath string, versionedLogoPath string) *actionTemplateService {
	actionTemplateService := &actionTemplateService{
		categoriesPath:    categoriesPath,
		logoPath:          logoPath,
		searchPath:        searchPath,
		versionedLogoPath: versionedLogoPath,
	}
	actionTemplateService.service = newService(serviceActionTemplateService, sling, uriTemplate, new(ActionTemplate))

	return actionTemplateService
}

func (s actionTemplateService) getPagedResponse(path string) ([]*ActionTemplate, error) {
	resources := []*ActionTemplate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(ActionTemplates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ActionTemplates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new action template.
func (s actionTemplateService) Add(resource *ActionTemplate) (*ActionTemplate, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

// GetAll returns all action templates. If none can be found or an error
// occurs, it returns an empty collection.
func (s actionTemplateService) GetAll() ([]*ActionTemplate, error) {
	items := []*ActionTemplate{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetCategories returns all action template categories.
func (s actionTemplateService) GetCategories() ([]ActionTemplateCategory, error) {
	err := validateInternalState(s)

	items := new([]ActionTemplateCategory)
	if err != nil {
		return *items, err
	}

	path := s.categoriesPath

	_, err = apiGet(s.getClient(), items, path)

	return *items, err
}

// GetByID returns the action template that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s actionTemplateService) GetByID(id string) (*ActionTemplate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ActionTemplate), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*ActionTemplate), nil
}

// Search lists all available action templates including built-in, custom, and community-contributed step templates.
func (s actionTemplateService) Search() ([]ActionTemplateSearch, error) {
	items := new([]ActionTemplateSearch)

	err := validateInternalState(s)
	if err != nil {
		return *items, err
	}

	path := s.searchPath

	_, err = apiGet(s.getClient(), items, path)

	return *items, err
}

// GetByName returns the action templates with a matching partial name.
func (s actionTemplateService) GetByName(name string) ([]*ActionTemplate, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*ActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an ActionTemplate based on the one provided as input.
func (s actionTemplateService) Update(resource ActionTemplate) (*ActionTemplate, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}
