package client

import (
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// actionTemplateService handles communication for any operations in the
// Octopus API that pertain to action templates.
type actionTemplateService struct {
	categoriesURL    url.URL                   `validate:"required"`
	name             string                    `validate:"required"`
	searchURL        url.URL                   `validate:"required"`
	sling            *sling.Sling              `validate:"required"`
	uriTemplate      *uritemplates.UriTemplate `validate:"required"`
	versionedLogoURL url.URL                   `validate:"required"`
}

// newActionTemplateService returns an actionTemplateService with a
// preconfigured client.
func newActionTemplateService(sling *sling.Sling, uriTemplate string, categoriesURL url.URL, searchURL url.URL, versionedLogoURL url.URL) *actionTemplateService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &actionTemplateService{
		categoriesURL:    categoriesURL,
		name:             serviceActionTemplateService,
		searchURL:        searchURL,
		sling:            sling,
		uriTemplate:      template,
		versionedLogoURL: versionedLogoURL,
	}
}

func (s actionTemplateService) getClient() *sling.Sling {
	return s.sling
}

func (s actionTemplateService) getName() string {
	return s.name
}

func (s actionTemplateService) getPagedResponse(path string) ([]model.ActionTemplate, error) {
	resources := []model.ActionTemplate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.ActionTemplates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.ActionTemplates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s actionTemplateService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new action template.
func (s actionTemplateService) Add(resource *model.ActionTemplate) (*model.ActionTemplate, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

// DeleteByID deletes the action template that matches the input ID.
func (s actionTemplateService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetAll returns all action templates. If none can be found or an error
// occurs, it returns an empty collection.
func (s actionTemplateService) GetAll() ([]model.ActionTemplate, error) {
	items := []model.ActionTemplate{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetCategories returns all action template categories.
func (s actionTemplateService) GetCategories() ([]model.ActionTemplateCategory, error) {
	err := validateInternalState(s)

	items := new([]model.ActionTemplateCategory)
	if err != nil {
		return *items, err
	}

	path := s.categoriesURL.String()

	_, err = apiGet(s.getClient(), items, path)

	return *items, err
}

// GetByID returns the action template that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s actionTemplateService) GetByID(id string) (*model.ActionTemplate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ActionTemplate), path)
	if err != nil {
		return nil, createResourceNotFoundError("action template", "ID", id)
	}

	return resp.(*model.ActionTemplate), nil
}

// Search lists all available action templates including built-in, custom, and community-contributed step templates.
func (s actionTemplateService) Search() ([]model.ActionTemplateSearch, error) {
	items := new([]model.ActionTemplateSearch)

	err := validateInternalState(s)
	if err != nil {
		return *items, err
	}

	path := s.searchURL.String()

	_, err = apiGet(s.getClient(), items, path)

	return *items, err
}

// GetByName returns the action templates with a matching partial name.
func (s actionTemplateService) GetByName(name string) ([]model.ActionTemplate, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []model.ActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an ActionTemplate based on the one provided as input.
func (s actionTemplateService) Update(resource model.ActionTemplate) (*model.ActionTemplate, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.ActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.ActionTemplate), nil
}

var _ ServiceInterface = &actionTemplateService{}
