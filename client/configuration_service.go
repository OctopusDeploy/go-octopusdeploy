package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type configurationService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newConfigurationService(sling *sling.Sling, uriTemplate string) *configurationService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &configurationService{
		name:        serviceConfigurationService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s configurationService) getClient() *sling.Sling {
	return s.sling
}

func (s configurationService) getName() string {
	return s.name
}

func (s configurationService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns a ConfigurationSection that matches the input ID. If one cannot be found, it returns nil and an error.
func (s configurationService) GetByID(id string) (*model.ConfigurationSection, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ConfigurationSection), path)
	if err != nil {
		return nil, createResourceNotFoundError("configuration", "ID", id)
	}

	return resp.(*model.ConfigurationSection), nil
}

func (s configurationService) getPagedResponse(path string) ([]model.ConfigurationSection, error) {
	resources := []model.ConfigurationSection{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.ConfigurationSections), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.ConfigurationSections)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

var _ ServiceInterface = &configurationService{}
