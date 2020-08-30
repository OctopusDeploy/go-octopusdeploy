package octopusdeploy

import "github.com/dghubble/sling"

// AccountService handles communication with Account-related methods of the
// Octopus API.
type AccountService struct {
	sling *sling.Sling
}

// NewAccountService returns an AccountService with a preconfigured client.
func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling,
	}
}
