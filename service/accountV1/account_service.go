package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/service"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with accountV1-related methods of the
// Octopus API.
type accountService struct {
	service.CanGetByIDService[IAccount]
	service.CanAddService[IAccount]
	service.CanUpdateService[IAccount]
	service.CanDeleteService[IAccount]
	service.CanQueryService[IAccount, service.AccountsQuery]
	service.ISpaceScopedService
}

type IAccountService interface {
	service.GetsByIDer[IAccount]
	service.ResourceAdder[IAccount]
	service.ResourceUpdater[IAccount]
	service.DeleteByIDer[IAccount]
	service.ResourceQueryer[IAccount, service.AccountsQuery]
	service.ISpaceScopedService
}

// NewAccountService returns an accountV1 service with a preconfigured client.
func NewAccountService(client service.ISpaceScopedClient) IAccountService {
	baseService := service.NewSpaceScopedService(service.ServiceAccountService, accountsV1BasePath, client)
	accountService := &accountService{
		ISpaceScopedService: baseService,
		CanGetByIDService: service.CanGetByIDService[IAccount]{
			IService: baseService,
		},
		CanAddService: service.CanAddService[IAccount]{
			IService: baseService,
		},
		CanUpdateService: service.CanUpdateService[IAccount]{
			IService: baseService,
		},
		CanDeleteService: service.CanDeleteService[IAccount]{
			IService: baseService,
		},
		CanQueryService: service.CanQueryService[IAccount, service.AccountsQuery]{
			IService: baseService,
		},
	}
	return accountService
}

// GetUsages lists the projects and deployments which are using an accountV1.
// func (s *accountService) GetUsages(account IAccount) (*AccountUsage, error) {
// 	path := account.GetLinks()
// 	resp, err := s.client.apiGet(new(AccountUsage), path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp.(*AccountUsage), nil
// }

// Update modifies an accountV1 based on the one provided as input.
// func (s *accountService) Update(account IAccount) (IAccount, error) {
// 	if account == nil {
// 		return nil, internal.CreateInvalidParameterError(service.OperationUpdate, octopusdeploy.ParameterAccount)
// 	}

// 	accountResource, err := ToAccountResource(s.client, account)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := s.client.apiUpdate(accountResource, new(AccountResource))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ToAccount(resp.(*AccountResource))
// }
