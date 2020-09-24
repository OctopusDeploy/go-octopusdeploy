package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type AzureDevOpsService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func NewAzureDevOpsService(sling *sling.Sling, uriTemplate string) *AzureDevOpsService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &AzureDevOpsService{
		name:        serviceAzureDevOpsService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s AzureDevOpsService) getClient() *sling.Sling {
	return s.sling
}

func (s AzureDevOpsService) getName() string {
	return s.name
}

func (s AzureDevOpsService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

var _ ServiceInterface = &AzureDevOpsService{}
