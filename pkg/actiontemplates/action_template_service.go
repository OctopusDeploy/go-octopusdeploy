package actiontemplates

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// ActionTemplateService handles communication for any operations in the
// Octopus API that pertain to action templates.
type ActionTemplateService struct {
	categoriesPath    string
	logoPath          string
	searchPath        string
	versionedLogoPath string

	services.CanDeleteService
}

// NewActionTemplateService returns an actionTemplateService with a
// preconfigured client.
func NewActionTemplateService(sling *sling.Sling, uriTemplate string, categoriesPath string, logoPath string, searchPath string, versionedLogoPath string) *ActionTemplateService {
	return &ActionTemplateService{
		categoriesPath:    categoriesPath,
		logoPath:          logoPath,
		searchPath:        searchPath,
		versionedLogoPath: versionedLogoPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceActionTemplateService, sling, uriTemplate),
		},
	}
}

// Add creates a new action template.
//
// Deprecated: Use actiontemplates.Add
func (s *ActionTemplateService) Add(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if IsNil(actionTemplate) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterActionTemplate)
	}

	if err := actionTemplate.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, actionTemplate)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), actionTemplate, new(ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return response.(*ActionTemplate), nil
}

// Get returns a collection of action templates based on the criteria defined
// by its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
//
// Deprecated: Use actiontemplates.Get
func (s *ActionTemplateService) Get(actionTemplatesQuery Query) (*resources.Resources[*ActionTemplate], error) {
	v, _ := query.Values(actionTemplatesQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*ActionTemplate]), path)
	if err != nil {
		return &resources.Resources[*ActionTemplate]{}, err
	}

	return resp.(*resources.Resources[*ActionTemplate]), nil
}

// GetAll returns all action templates. If none can be found or an error
// occurs, it returns an empty collection.
//
// Deprecated: Use actiontemplates.GetAll
func (s *ActionTemplateService) GetAll() ([]*ActionTemplate, error) {
	items := []*ActionTemplate{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetCategories returns all action template categories.
//
// Deprecated: Use actiontemplates.GetCategories
func (s *ActionTemplateService) GetCategories() ([]ActionTemplateCategory, error) {
	items := new([]ActionTemplateCategory)
	if err := services.ValidateInternalState(s); err != nil {
		return *items, err
	}

	path := s.categoriesPath

	_, err := api.ApiGet(s.GetClient(), items, path)

	return *items, err
}

// GetByID returns the action template that matches the input ID. If one cannot
// be found, it returns nil and an error.
//
// Deprecated: Use actiontemplates.GetByID
func (s *ActionTemplateService) GetByID(id string) (*ActionTemplate, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

// Search lists all available action templates including built-in, custom, and
// community-contributed step templates.
//
// Deprecated: Use actiontemplates.Search
func (s *ActionTemplateService) Search(searchQuery string) ([]ActionTemplateSearch, error) {
	searchResults := []ActionTemplateSearch{}
	if err := services.ValidateInternalState(s); err != nil {
		return searchResults, err
	}

	template, err := uritemplates.Parse(s.searchPath)
	if err != nil {
		return searchResults, err
	}

	path, err := template.Expand(map[string]interface{}{"type": searchQuery})
	if err != nil {
		return searchResults, err
	}

	if len(searchQuery) <= 0 {
		path = strings.Split(path, "?")[0]
	}

	_, err = api.ApiGet(s.GetClient(), &searchResults, path)

	return searchResults, err
}

// Update modifies an ActionTemplate based on the one provided as input.
//
// Deprecated: Use actiontemplates.Update
func (s *ActionTemplateService) Update(actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if actionTemplate == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "actionTemplate")
	}

	path, err := services.GetUpdatePath(s, actionTemplate)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), actionTemplate, new(ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ActionTemplate), nil
}

// ----- new -----
const (
	template           = "/api/{spaceId}/actiontemplates{/id}"
	categoriesTemplate = "/api/{spaceId}/actiontemplates/categories"
	searchTemplate     = "/api/{spaceId}/actiontemplates/search"
)

// Add creates a new action template.
func Add(client newclient.Client, actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if IsNil(actionTemplate) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterActionTemplate)
	}

	if err := actionTemplate.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	return newclient.Add[ActionTemplate](client, template, actionTemplate.SpaceID, actionTemplate)
}

// Get returns a collection of action templates based on the criteria defined
// by its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, actionTemplatesQuery Query) (*resources.Resources[*ActionTemplate], error) {
	return newclient.GetByQuery[ActionTemplate](client, template, spaceID, actionTemplatesQuery)
}

// GetAll returns all action templates. If none can be found or an error
// occurs, it returns an empty collection.
func GetAll(client newclient.Client, spaceID string) ([]*ActionTemplate, error) {
	return newclient.GetAll[ActionTemplate](client, template, spaceID)
}

// GetCategories returns all action template categories.
func GetCategories(client newclient.Client, spaceID string) ([]*ActionTemplateCategory, error) {
	return newclient.GetAll[ActionTemplateCategory](client, categoriesTemplate, spaceID)
}

// GetByID returns the action template that matches the input ID. If one cannot
// be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*ActionTemplate, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	return newclient.GetByID[ActionTemplate](client, template, spaceID, id)
}

// Search lists all available action templates including built-in, custom, and
// community-contributed step templates.
func Search(client newclient.Client, spaceID string, searchQuery string) (*[]ActionTemplateSearch, error) {
	path, err := client.URITemplateCache().Expand(searchTemplate, map[string]any{
		"spaceId": spaceID,
		"type":    searchQuery,
	})
	if err != nil {
		return nil, err
	}

	if len(searchQuery) <= 0 {
		path = strings.Split(path, "?")[0]
	}

	return newclient.Get[[]ActionTemplateSearch](client.HttpSession(), path)
}

// Update modifies an ActionTemplate based on the one provided as input.
func Update(client newclient.Client, actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	if actionTemplate == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "actionTemplate")
	}

	return newclient.Update[ActionTemplate](client, template, actionTemplate.SpaceID, actionTemplate.ID, actionTemplate)
}

func DeleteByID(client newclient.Client, spaceID string, id string) error {
	return newclient.DeleteByID(client, template, spaceID, id)
}
