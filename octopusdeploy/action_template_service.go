package octopusdeploy

import (
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
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
	actionTemplateService.service = newService(ServiceActionTemplateService, sling, uriTemplate)

	return actionTemplateService
}

// Add creates a new action template.
func (s actionTemplateService) Add(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if actionTemplate == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterActionTemplate)
	}

	path, err := getAddPath(s, actionTemplate)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), actionTemplate, new(ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return response.(*ActionTemplate), nil
}

// Get returns a collection of action templates based on the criteria defined
// by its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s actionTemplateService) Get(actionTemplatesQuery ActionTemplatesQuery) (*ActionTemplates, error) {
	v, _ := query.Values(actionTemplatesQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := apiGet(s.getClient(), new(ActionTemplates), path)
	if err != nil {
		return &ActionTemplates{}, err
	}

	return resp.(*ActionTemplates), nil
}

// GetAll returns all action templates. If none can be found or an error
// occurs, it returns an empty collection.
func (s actionTemplateService) GetAll() ([]*ActionTemplate, error) {
	items := []*ActionTemplate{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
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
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

// Search lists all available action templates including built-in, custom, and
// community-contributed step templates.
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

// Update modifies an ActionTemplate based on the one provided as input.
func (s actionTemplateService) Update(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if actionTemplate == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterActionTemplate)
	}

	path, err := getUpdatePath(s, actionTemplate)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), actionTemplate, new(ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}
