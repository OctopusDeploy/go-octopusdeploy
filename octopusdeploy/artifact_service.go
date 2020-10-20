package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// actionTemplateService handles communication for any operations in the
// Octopus API that pertain to artifacts.
type artifactService struct {
	canDeleteService
}

// newArtifactService returns an artifactService with a preconfigured client.
func newArtifactService(sling *sling.Sling, uriTemplate string) *artifactService {
	artifactService := &artifactService{}
	artifactService.service = newService(serviceArtifactService, sling, uriTemplate, new(Artifact))

	return artifactService
}

func (s artifactService) getPagedResponse(path string) ([]*Artifact, error) {
	resources := []*Artifact{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Artifacts), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Artifacts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new artifact.
func (s artifactService) Add(resource *Artifact) (*Artifact, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}

// GetAll returns all artifacts. If none can be found or an error occurs, it
// returns an empty collection.
func (s artifactService) GetAll() ([]*Artifact, error) {
	path, err := getPath(s)
	if err != nil {
		return []*Artifact{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the artifact that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s artifactService) GetByID(id string) (*Artifact, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Artifact), nil
}

// GetByPartialName performs a lookup and returns all instances of Artifact with a matching partial name.
func (s artifactService) GetByPartialName(name string) ([]*Artifact, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Artifact{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an Artifact based on the one provided as input.
func (s artifactService) Update(resource Artifact) (*Artifact, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}
