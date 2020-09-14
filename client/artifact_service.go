package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/go-playground/validator"
)

// ArtifactService handles communication with Account-related methods of the
// Octopus API.
type ArtifactService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewArtifactService returns an ArtifactService with a preconfigured client.
func NewArtifactService(sling *sling.Sling) *ArtifactService {
	if sling == nil {
		return nil
	}

	return &ArtifactService{
		sling: sling,
		path:  "artifacts",
	}
}

// Get returns an Artifact that matches the input ID.
func (s *ArtifactService) Get(id string) (*model.Artifact, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("ArtifactService: invalid parameter, id")
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
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if artifact == nil {
		return nil, errors.New("ArtifactService: invalid parameter, artifact")
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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("ArtifactService: invalid parameter, id")
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
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ServiceInterface = &ArtifactService{}
