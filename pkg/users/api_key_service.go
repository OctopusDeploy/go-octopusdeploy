package users

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

// ApiKeyService handles communication with API key-related methods of the Octopus API.
type ApiKeyService struct {
	services.Service
}

// NewAPIKeyService returns an ApiKeyService with a preconfigured client.
func NewAPIKeyService(sling *sling.Sling, uriTemplate string) *ApiKeyService {
	return &ApiKeyService{
		Service: services.NewService("APIKeyService", sling, uriTemplate),
	}
}

// GetByUserID lists all API keys for a user, returning the most recent results first.
func (s *ApiKeyService) GetByUserID(userID string) ([]*APIKey, error) {
	if internal.IsEmpty(userID) {
		return nil, internal.CreateInvalidParameterError("GetByUserID", "userID")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	var p []*APIKey

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/%s/apikeys", userID)

	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(resources.Resources[APIKey]), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*resources.Resources[APIKey])
		p = append(p, r.Items...)
		path, loadNextPage = services.LoadNextPage(r.PagedResults)
	}

	return p, nil
}

// GetByID the API key that belongs to the user by its ID.
func (s *ApiKeyService) GetByID(userID string, apiKeyID string) (*APIKey, error) {
	if internal.IsEmpty(userID) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, "userID")
	}

	if internal.IsEmpty(apiKeyID) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, "apiKeyID")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"%s/apikeys/%s", userID, apiKeyID)

	resp, err := services.ApiGet(s.GetClient(), new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*APIKey), nil
}

// Create generates a new API key for the specified user ID. The API key
// returned in the result must be saved by the caller, as it cannot be
// retrieved subsequently from the Octopus server.
func (s *ApiKeyService) Create(apiKey *APIKey) (*APIKey, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	if err := apiKey.Validate(); err != nil {
		return nil, err
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/%s/apikeys", apiKey.UserID)

	resp, err := services.ApiPost(s.GetClient(), apiKey, new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return resp.(*APIKey), nil
}
