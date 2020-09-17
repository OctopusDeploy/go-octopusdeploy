package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// ArtifactService handles communication with Account-related methods of the
// Octopus API.
type ArtifactService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewArtifactService returns an ArtifactService with a preconfigured client.
func NewArtifactService(sling *sling.Sling, uriTemplate string) *ArtifactService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &ArtifactService{
		name:  "ArtifactService",
		path:  path,
		sling: sling,
	}
}

// Get returns an Artifact that matches the input ID.
func (s *ArtifactService) Get(id string) (*model.Artifact, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Artifact), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// GetAll returns all instances of an Artifact.
func (s *ArtifactService) GetAll() (*[]model.Artifact, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

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
func (s *ArtifactService) Add(artifact *model.Artifact) (*model.Artifact, error) {
	if artifact == nil {
		return nil, createInvalidParameterError("Add", "artifact")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = artifact.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, artifact, new(model.Artifact), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// Delete removes the Artifact that matches the input ID.
func (s *ArtifactService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()
	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Artifact based on the one provided as input.
func (s *ArtifactService) Update(artifact model.Artifact) (*model.Artifact, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = artifact.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", artifact.ID)
	resp, err := apiUpdate(s.sling, artifact, new(model.Artifact), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

func (s *ArtifactService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &ArtifactService{}
