package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type configurationService struct {
	versionControlClearCachePath string

	services.service
}

func newConfigurationService(sling *sling.Sling, uriTemplate string, versionControlClearCachePath string) *configurationService {
	return &configurationService{
		versionControlClearCachePath: versionControlClearCachePath,
		service:                      services.newService(ServiceConfigurationService, sling, uriTemplate),
	}
}

// GetByID returns a ConfigurationSection that matches the input ID. If one cannot be found, it returns nil and an error.
func (s configurationService) GetByID(id string) (*ConfigurationSection, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(ConfigurationSection), path)
	if err != nil {
		return nil, err
	}

	return resp.(*ConfigurationSection), nil
}

func (s configurationService) getPagedResponse(path string) ([]*ConfigurationSection, error) {
	resources := []*ConfigurationSection{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(ConfigurationSections), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*ConfigurationSections)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}
