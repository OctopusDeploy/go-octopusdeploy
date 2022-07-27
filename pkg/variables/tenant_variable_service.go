package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type TenantVariableService struct {
	services.Service
}

func NewTenantVariableService(sling *sling.Sling, uriTemplate string) *TenantVariableService {
	return &TenantVariableService{
		Service: services.NewService(constants.ServiceTenantVariableService, sling, uriTemplate),
	}
}

func (s *TenantVariableService) GetAll() ([]TenantVariables, error) {
	items := []TenantVariables{}
	path, err := services.GetPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}
