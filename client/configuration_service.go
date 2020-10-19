package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type configurationService struct {
	versionControlClearCachePath string

	service
}

func newConfigurationService(sling *sling.Sling, uriTemplate string, versionControlClearCachePath string) *configurationService {
	return &configurationService{
		versionControlClearCachePath: versionControlClearCachePath,
		service:                      newService(serviceConfigurationService, sling, uriTemplate, nil),
	}
}

// GetByID returns a ConfigurationSection that matches the input ID. If one cannot be found, it returns nil and an error.
func (s configurationService) GetByID(id string) (*model.ConfigurationSection, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.ConfigurationSection), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.ConfigurationSection), nil
}

func (s configurationService) getPagedResponse(path string) ([]*model.ConfigurationSection, error) {
	resources := []*model.ConfigurationSection{}
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
