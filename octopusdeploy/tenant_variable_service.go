package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type tenantVariableService struct {
	services.service
}

func newTenantVariableService(sling *sling.Sling, uriTemplate string) *tenantVariableService {
	return &tenantVariableService{
		service: services.newService(ServiceTenantVariableService, sling, uriTemplate),
	}
}

func (s tenantVariableService) GetAll() ([]TenantVariables, error) {
	items := []TenantVariables{}
	path, err := getPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}
