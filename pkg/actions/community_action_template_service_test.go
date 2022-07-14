package actions

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createCommunityActionTemplateService(t *testing.T) *CommunityActionTemplateService {
	service := NewCommunityActionTemplateService(nil, constants.TestURICommunityActionTemplates)
	services.NewServiceTests(t, service, constants.TestURICommunityActionTemplates, constants.ServiceCommunityActionTemplateService)
	return service
}

func TestCommunityActionTemplateServiceNew(t *testing.T) {
	ServiceFunction := NewCommunityActionTemplateService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceCommunityActionTemplateService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *CommunityActionTemplateService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestCommunityActionTemplateServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createCommunityActionTemplateService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
			assert.Nil(t, resource)
		})
	}
}

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Install(CommunityActionTemplate{})
	require.Error(t, err)
	require.Nil(t, resource)
}
