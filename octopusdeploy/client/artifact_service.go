package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

// ArtifactService handles communication with Account-related methods of the
// Octopus API.
type ArtifactService struct {
	sling *sling.Sling
	path  string
}

// NewArtifactService returns an ArtifactService with a preconfigured client.
func NewArtifactService(sling *sling.Sling) *ArtifactService {
	return &ArtifactService{
		sling: sling,
		path:  "artifacts",
	}
}

// Get returns an Artifact that matches the input ID.
func (s *ArtifactService) Get(id string) (*model.Artifact, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Artifact), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// GetAll returns all of the Accounts for a Space.
func (s *ArtifactService) GetAll() (*[]model.Artifact, error) {
	var p []model.Artifact
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Artifacts), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Artifacts)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// Add creates a new Artifact.
func (s *ArtifactService) Add(resource *model.Artifact) (*model.Artifact, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Artifact), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// Delete removes the Artifact that matches the input ID.
func (s *ArtifactService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Artifact based on the one provided as input.
func (s *ArtifactService) Update(resource *model.Artifact) (*model.Artifact, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Artifact), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}
