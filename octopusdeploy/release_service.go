package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type releaseService struct {
	canDeleteService
}

func newReleaseService(sling *sling.Sling, uriTemplate string) *releaseService {
	releaseService := &releaseService{}
	releaseService.service = newService(serviceReleaseService, sling, uriTemplate, new(Release))

	return releaseService
}

func (s releaseService) getPagedResponse(path string) ([]*Release, error) {
	resources := []*Release{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Releases), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Releases)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetByID returns the release that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s releaseService) GetByID(id string) (*Release, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Release), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Release), nil
}
