package octopusdeploy

import (
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// actionTemplateService handles communication for any operations in the
// Octopus API that pertain to artifacts.
type artifactService struct {
	canDeleteService
}

// newArtifactService returns an artifactService with a preconfigured client.
func newArtifactService(sling *sling.Sling, uriTemplate string) *artifactService {
	artifactService := &artifactService{}
	artifactService.service = newService(ServiceArtifactService, sling, uriTemplate)

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

	resp, err := apiAdd(s.getClient(), resource, new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}

// Get returns a collection of artifacts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s artifactService) Get(artifactsQuery ArtifactsQuery) (*Artifacts, error) {
	v, _ := query.Values(artifactsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := apiGet(s.getClient(), new(Artifacts), path)
	if err != nil {
		return &Artifacts{}, err
	}

	return resp.(*Artifacts), nil
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

	resp, err := apiGet(s.getClient(), new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}

// Update modifies an Artifact based on the one provided as input.
func (s artifactService) Update(artifact Artifact) (*Artifact, error) {
	path, err := getUpdatePath(s, &artifact)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), artifact, new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}
