package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// APIKeyService handles communication with API key-related methods of the
// Octopus API.
type APIKeyService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewAPIKeyService returns an APIKeyService with a preconfigured client.
func NewAPIKeyService(sling *sling.Sling, uriTemplate string) *APIKeyService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &APIKeyService{
		name:  "APIKeyService",
		path:  path,
		sling: sling,
	}
}

// Get lists all API keys for a user, returning the most recent results first.
func (s *APIKeyService) Get(userID string) (*[]model.APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError("Get", "userID")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	var p []model.APIKey
	path := fmt.Sprintf(s.path+"/%s/apikeys", userID)
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.APIKeys), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.APIKeys)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByID the API key that belongs to the user by its ID.
func (s *APIKeyService) GetByID(userID string, apiKeyID string) (*model.APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError("GetByID", "userID")
	}

	if isEmpty(apiKeyID) {
		return nil, createInvalidParameterError("GetByID", "apiKeyID")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s/apikeys/%s", userID, apiKeyID)
	resp, err := apiGet(s.sling, new(model.APIKey), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}

// Create generates a new API key for the specified user ID. The API key
// returned in the result must be saved by the caller, as it cannot be
// retrieved subsequently from the Octopus server.
func (s *APIKeyService) Create(apiKey *model.APIKey) (*model.APIKey, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = apiKey.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s/apikeys", *apiKey.UserID)
	resp, err := apiPost(s.sling, apiKey, new(model.APIKey), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}

func (s *APIKeyService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &APIKeyService{}
