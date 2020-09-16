package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/go-playground/validator"
)

// APIKeyService handles communication with API key-related methods of the
// Octopus API.
type APIKeyService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewAPIKeyService returns an APIKeyService with a preconfigured client.
func NewAPIKeyService(sling *sling.Sling, uriTemplate string) *APIKeyService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &APIKeyService{
		sling: sling,
		path:  path,
	}
}

// Get lists all API keys for a user, returning the most recent results first.
func (s *APIKeyService) Get(userID string) (*[]model.APIKey, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(userID) {
		return nil, errors.New("APIKeyService: invalid parameter, userID")
	}

	var p []model.APIKey
	path := fmt.Sprintf("/%s/apikeys", userID)
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
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(userID) {
		return nil, errors.New("APIKeyService: invalid parameter, userID")
	}

	if isEmpty(apiKeyID) {
		return nil, errors.New("APIKeyService: invalid parameter, apiKeyID")
	}

	path := fmt.Sprintf("users/%s/apikeys/%s", userID, apiKeyID)
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

	path := fmt.Sprintf("users/%s/apikeys", *apiKey.UserID)
	resp, err := apiPost(s.sling, apiKey, new(model.APIKey), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}

func (s *APIKeyService) validateInternalState() error {
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

var _ ServiceInterface = &APIKeyService{}
