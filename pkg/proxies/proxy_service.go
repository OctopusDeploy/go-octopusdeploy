package proxies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type ProxyService struct {
	services.CanDeleteService
}

func NewProxyService(sling *sling.Sling, uriTemplate string) *ProxyService {
	return &ProxyService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceProxyService, sling, uriTemplate),
		},
	}
}

func (p *ProxyService) GetById(id string) (*Proxy, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, "id")
	}

	path := p.BasePath + "/" + id
	resp, err := api.ApiGet(p.GetClient(), new(Proxy), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Proxy), nil
}

func (p *ProxyService) GetAll() ([]*Proxy, error) {
	items := []*Proxy{}
	path, err := services.GetAllPath(p)
	if err != nil {
		return nil, err
	}

	_, err = api.ApiGet(p.GetClient(), &items, path)
	return items, err
}
