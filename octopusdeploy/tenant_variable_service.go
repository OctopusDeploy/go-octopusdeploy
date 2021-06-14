package octopusdeploy

import "github.com/dghubble/sling"

type tenantVariableService struct {
	service
}

func newTenantVariableService(sling *sling.Sling, uriTemplate string) *tenantVariableService {
	return &tenantVariableService{
		service: newService(ServiceTenantVariableService, sling, uriTemplate),
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
