package artifacts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// ArtifactService handles communication for any operations in the
// Octopus API that pertain to artifacts.
type ArtifactService struct {
	services.CanDeleteService
}

// NewArtifactService returns an artifactService with a preconfigured client.
func NewArtifactService(sling *sling.Sling, uriTemplate string) *ArtifactService {
	return &ArtifactService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceArtifactService, sling, uriTemplate),
		},
	}
}

// Add creates a new artifact.
func (s *ArtifactService) Add(artifact *Artifact) (*Artifact, error) {
	if IsNil(artifact) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterArtifact)
	}

	path, err := services.GetAddPath(s, artifact)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), artifact, new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}

// Get returns a collection of artifacts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *ArtifactService) Get(artifactsQuery Query) (*resources.Resources[*Artifact], error) {
	v, _ := query.Values(artifactsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Artifact]), path)
	if err != nil {
		return &resources.Resources[*Artifact]{}, err
	}

	return resp.(*resources.Resources[*Artifact]), nil
}

// GetAll returns all artifacts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *ArtifactService) GetAll() ([]*Artifact, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return []*Artifact{}, err
	}

	return services.GetPagedResponse[Artifact](s, path)
}

// GetByID returns the artifact that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *ArtifactService) GetByID(id string) (*Artifact, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}

// Update modifies an Artifact based on the one provided as input.
func (s *ArtifactService) Update(artifact Artifact) (*Artifact, error) {
	path, err := services.GetUpdatePath(s, &artifact)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), artifact, new(Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Artifact), nil
}
