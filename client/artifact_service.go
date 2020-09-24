package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// artifactService handles communication with Account-related methods of the Octopus API.
type artifactService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newArtifactService returns an artifactService with a preconfigured client.
func newArtifactService(sling *sling.Sling, uriTemplate string) *artifactService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &artifactService{
		name:        serviceArtifactService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s artifactService) getClient() *sling.Sling {
	return s.sling
}

func (s artifactService) getName() string {
	return s.name
}

func (s artifactService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns an Artifact that matches the input ID. If one cannot be found, it returns nil and an error.
func (s artifactService) GetByID(id string) (*model.Artifact, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// GetByPartialName performs a lookup and returns all instances of Artifact with a matching partial name.
func (s artifactService) GetByPartialName(name string) ([]model.Artifact, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return nil, err
	}

	return s.getPagedResponse(path)
}

// Add creates a new Artifact.
func (s artifactService) Add(artifact *model.Artifact) (*model.Artifact, error) {
	path, err := getAddPath(s, artifact)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), artifact, new(model.Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

// DeleteByID deletes the Artifact that matches the input ID.
func (s artifactService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// Update modifies an Artifact based on the one provided as input.
func (s artifactService) Update(resource model.Artifact) (*model.Artifact, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Artifact), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Artifact), nil
}

func (s artifactService) getPagedResponse(path string) ([]model.Artifact, error) {
	var resources []model.Artifact
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Artifacts), path)
		if err != nil {
			return nil, err
		}

		responseList := resp.(*model.Artifacts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

var _ ServiceInterface = &artifactService{}
