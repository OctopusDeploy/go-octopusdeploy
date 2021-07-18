package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

// apiKeyService handles communication with API key-related methods of the Octopus API.
type apiKeyService struct {
	service
}

// newAPIKeyService returns an apiKeyService with a preconfigured client.
func newAPIKeyService(sling *sling.Sling, uriTemplate string) *apiKeyService {
	return &apiKeyService{
		service: newService(ServiceAPIKeyService, sling, uriTemplate),
	}
}

// GetByUserID lists all API keys for a user, returning the most recent results first.
func (s apiKeyService) GetByUserID(userID string) ([]*APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError(OperationGetByUserID, ParameterUserID)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	var p []*APIKey

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/apikeys", userID)

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(APIKeys), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*APIKeys)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// GetByID the API key that belongs to the user by its ID.
func (s apiKeyService) GetByID(userID string, apiKeyID string) (*APIKey, error) {
	if isEmpty(userID) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterUserID)
	}

	if isEmpty(apiKeyID) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterAPIKeyID)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"%s/apikeys/%s", userID, apiKeyID)

	resp, err := apiGet(s.getClient(), new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*APIKey), nil
}

// Create generates a new API key for the specified user ID. The API key
// returned in the result must be saved by the caller, as it cannot be
// retrieved subsequently from the Octopus server.
func (s apiKeyService) Create(apiKey *APIKey) (*APIKey, error) {
	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	if err := apiKey.Validate(); err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/apikeys", apiKey.UserID)

	resp, err := apiPost(s.getClient(), apiKey, new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*APIKey), nil
}
