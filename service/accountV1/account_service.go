package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/service"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with accountV1-related methods of the
// Octopus API.
type accountService struct {
	client *service.SpaceScopedClient
	service.SpaceScopedService
	service.GetsByIDer[IAccount]
	service.ResourceQueryer[IAccount]
	service.CanAddService[IAccount]
	service.CanUpdateService[IAccount]
	service.CanDeleteService[IAccount]
}

// NewAccountService returns an accountV1 service with a preconfigured client.
func NewAccountService(client *service.SpaceScopedClient) *accountService {
	accountService := &accountService{
		SpaceScopedService: service.NewSpaceScopedService(service.ServiceAccountService, accountsV1BasePath, client),
	}

	return accountService
}

// GetUsages lists the projects and deployments which are using an accountV1.
func (s *accountService) GetUsages(account IAccount) (*AccountUsage, error) {
	path := account.GetLinks()
	resp, err := s.client.apiGet(new(AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*AccountUsage), nil
}

// Update modifies an accountV1 based on the one provided as input.
func (s *accountService) Update(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, internal.CreateInvalidParameterError(service.OperationUpdate, octopusdeploy.ParameterAccount)
	}

	accountResource, err := ToAccountResource(s.client, account)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(accountResource, new(AccountResource))
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}
