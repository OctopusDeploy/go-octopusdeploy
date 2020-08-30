package octopusdeploy

import "github.com/dghubble/sling"

type AccountService struct {
	sling *sling.Sling
}

func NewAccountService(sling *sling.Sling) *AccountService {
	return &AccountService{
		sling: sling,
	}
}
