package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// APIKeyService handles communication with API key-related methods of the
// Octopus API.
type APIKeyService struct {
	sling *sling.Sling
}

// NewAPIKeyService returns an APIKeyService with a preconfigured client.
func NewAPIKeyService(sling *sling.Sling) *APIKeyService {
	if sling == nil {
		fmt.Println(fmt.Errorf("APIKeyService: input parameter (sling) is nil"))
		return nil
	}

	return &APIKeyService{
		sling: sling,
	}
}

// Get lists all API keys for a user, returning the most recent results first.
func (s *APIKeyService) Get(userID string) (*model.APIKey, error) {
	path := fmt.Sprintf("users/%s/apikeys", userID)
	resp, err := apiGet(s.sling, new(model.APIKey), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.APIKey), nil
}
