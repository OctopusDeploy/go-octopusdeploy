package proxies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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

// Deprecated: use proxies.GetByID
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

// Deprecated: use proxies.GetAll
func (p *ProxyService) GetAll() ([]*Proxy, error) {
	items := []*Proxy{}
	path, err := services.GetAllPath(p)
	if err != nil {
		return nil, err
	}

	_, err = api.ApiGet(p.GetClient(), &items, path)
	return items, err
}

// --- new ---
const template = "/api/{spaceId}/proxies{/id}{?skip,take,ids,partialName}"

// Get return the machine proxies that match the provided proxies query.
func Get(client newclient.Client, spaceID string, proxyQuery ProxiesQuery) (*resources.Resources[*Proxy], error) {
	return newclient.GetByQuery[Proxy](client, template, spaceID, proxyQuery)
}

// Update modifies a machine proxy based on the one provided as input.
func Update(client newclient.Client, resource *Proxy) (*Proxy, error) {
	return newclient.Update[Proxy](client, template, resource.SpaceID, resource.ID, resource)
}

// Add creates a new machine proxy.
func Add(client newclient.Client, proxy *Proxy) (*Proxy, error) {
	return newclient.Add[Proxy](client, template, proxy.SpaceID, proxy)
}

// GetByID returns the machine proxy that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*Proxy, error) {
	return newclient.GetByID[Proxy](client, template, spaceID, ID)
}

// DeleteByID deletes the machine proxy that matches the input ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetAll returns all machine proxies. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*Proxy, error) {
	return newclient.GetAll[Proxy](client, template, spaceID)
}
