package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type rootService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newRootService(sling *sling.Sling, uriTemplate string) *rootService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &rootService{
		name:        serviceRootService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s rootService) getClient() *sling.Sling {
	return s.sling
}

func (s rootService) getName() string {
	return s.name
}

func (s rootService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

func (s rootService) Get() (*model.RootResource, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.RootResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.RootResource), nil
}

var _ ServiceInterface = &rootService{}
