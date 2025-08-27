package actiontemplates

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
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
// Deprecated: use actiontemplates.Add
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
// Deprecated: use actiontemplates.GetByID
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

// GetByPartialName performs a lookup and returns action templates with a matching
// partial name.
func (s *ActionTemplateService) GetByPartialName(partialName string) ([]*ActionTemplateSearch, error) {
	if internal.IsEmpty(partialName) {
		return []*ActionTemplateSearch{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*ActionTemplateSearch{}, err
	}

	return services.GetPagedResponse[ActionTemplateSearch](s, path)
}

// Update modifies an ActionTemplate based on the one provided as input.
//
// Deprecated: use actiontemplates.Update
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

// --- new ---

const template = "/api/{spaceId}/actiontemplates{/id}{?skip,take,ids,partialName}"
const templateVersions = "/api/{spaceId}/actiontemplates/{id}/versions/{version}"

// Add creates a new action template.
func Add(client newclient.Client, actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	return newclient.Add[ActionTemplate](client, template, actionTemplate.SpaceID, actionTemplate)
}

// DeleteByID deletes the action template based on the ID provided as input.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetByID returns the action template that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, actionTemplateID string) (*ActionTemplate, error) {
	return newclient.GetByID[ActionTemplate](client, template, spaceID, actionTemplateID)
}

// Update modifies an action template based on the one provided as input.
func Update(client newclient.Client, actionTemplate *ActionTemplate) (*ActionTemplate, error) {
	return newclient.Update[ActionTemplate](client, template, actionTemplate.SpaceID, actionTemplate.ID, actionTemplate)
}

// GetVersionByID returns the action template that matches the input ID and a Version. If one cannot be
// found, it returns nil and an error.
func GetVersionByID(client newclient.Client, spaceID string, actionTemplateID string, actionTemplateVersion int32) (*ActionTemplate, error) {
	spaceId, spaceIdError := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if spaceIdError != nil {
		return nil, spaceIdError
	}

	params := map[string]any{"spaceId": spaceId, "id": actionTemplateID, "version": actionTemplateVersion}
	uri, uriError := client.URITemplateCache().Expand(templateVersions, params)
	if uriError != nil {
		return nil, uriError
	}

	return newclient.Get[ActionTemplate](client.HttpSession(), uri)
}

// Get returns a collection of action templates based on the criteria defined by its
// input query parameter.
func Get(client newclient.Client, spaceID string, actionsQuery ActionTemplateSearch) (*resources.Resources[*ActionTemplate], error) {
	return newclient.GetByQuery[ActionTemplate](client, template, spaceID, actionsQuery)
}
