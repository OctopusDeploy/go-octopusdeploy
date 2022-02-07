package octopusdeploy

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"

type rootService struct {
	services.service
}

func newRootService(client AdminClient) *rootService {
	return &rootService{
		service: services.newService(ServiceRootService, sling, uriTemplate),
	}
}

func (s rootService) Get() (*RootResource, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(RootResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RootResource), nil
}

var _ IService = &rootService{}
