package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type RootService struct {
	services.Service
}

func NewRootService(sling *sling.Sling, uriTemplate string) *RootService {
	return &RootService{
		Service: services.NewService(constants.ServiceRootService, sling, uriTemplate),
	}
}

func (s *RootService) GetPath() string {
	return "/api"
}

func (s *RootService) Get() (*RootResource, error) {
	path := s.GetPath()
	resp, err := api.ApiGet(s.GetClient(), new(RootResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RootResource), nil
}

var _ services.IService = &RootService{}

const (
	template = "/api/{spaceId}"
)

func GetSpaceRoot(client newclient.Client, spaceID string) (*resources.Resource, error) {
	values := map[string]any{
		"spaceId": spaceID,
	}
	path, err := client.URITemplateCache().Expand(template, values)

	if err != nil {
		return nil, err
	}

	return newclient.Get[resources.Resource](client.HttpSession(), path)
}
