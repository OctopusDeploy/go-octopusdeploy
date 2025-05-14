package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
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

func (s *RootService) Get() (*RootResource, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return nil, err
	}

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

func GetSpaceRoot(client newclient.Client, spaceID *string) (*RootResource, error) {
	values := map[string]any{}
	if spaceID != nil {
		values["spaceId"] = *spaceID
	}
	path, err := client.URITemplateCache().Expand(template, values)

	if err != nil {
		return nil, err
	}
	return newclient.Get[RootResource](client.HttpSession(), path)
}

func GetServerRoot(client newclient.Client) (*RootResource, error) {
	return GetSpaceRoot(client, nil)
}
