package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type buildInformationService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newBuildInformationService(sling *sling.Sling, uriTemplate string) *buildInformationService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &buildInformationService{
		name:        serviceBuildInformationService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s buildInformationService) getClient() *sling.Sling {
	return s.sling
}

func (s buildInformationService) getName() string {
	return s.name
}

func (s buildInformationService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

var _ ServiceInterface = &buildInformationService{}
