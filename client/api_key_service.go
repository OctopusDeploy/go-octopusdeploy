package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// apiKeyService handles communication with API key-related methods of the Octopus API.
type apiKeyService struct {
	service
}

// newAPIKeyService returns an apiKeyService with a preconfigured client.
func newAPIKeyService(sling *sling.Sling, uriTemplate string) *apiKeyService {
	return &apiKeyService{
		service: newService(serviceAPIKeyService, sling, uriTemplate, new(model.APIKey)),
	}
}

// GetByUserID lists all API keys for a user, returning the most recent results first.
func (s apiKeyService) GetByUserID(userID string) ([]*model.APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError(operationGetByUserID, parameterUserID)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	var p []*model.APIKey

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/apikeys", userID)

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.APIKeys), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*model.APIKeys)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// GetByID the API key that belongs to the user by its ID.
func (s apiKeyService) GetByID(userID string, apiKeyID string) (*model.APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError(operationGetByID, parameterUserID)
	}

	if isEmpty(apiKeyID) {
		return nil, createInvalidParameterError(operationGetByID, parameterAPIKeyID)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"%s/apikeys/%s", userID, apiKeyID)

	resp, err := apiGet(s.getClient(), new(model.APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}

// Create generates a new API key for the specified user ID. The API key
// returned in the result must be saved by the caller, as it cannot be
// retrieved subsequently from the Octopus server.
func (s apiKeyService) Create(apiKey *model.APIKey) (*model.APIKey, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	err = apiKey.Validate()

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/apikeys", apiKey.UserID)

	resp, err := apiPost(s.getClient(), apiKey, new(model.APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}
