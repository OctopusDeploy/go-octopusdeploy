package channels

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createChannelService(t *testing.T) *ChannelService {
	service := NewChannelService(nil, constants.TestURIChannels, constants.TestURIVersionRuleTest)
	services.NewServiceTests(t, service, constants.TestURIChannels, constants.ServiceChannelService)
	return service
}

func TestChannelServiceAdd(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterChannel))
	assert.Nil(t, resource)

	invalidResource := &Channel{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)

	invalidResource = NewChannel("test-channel", "Projects-1")
	invalidResource.Type = "invalid"
	resource, err = service.Add(invalidResource)
	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestChannelServiceGetByID(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	assert.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	assert.Nil(t, resource)
}

func TestChannelServiceNew(t *testing.T) {
	ServiceFunction := NewChannelService
	client := &sling.Sling{}
	uriTemplate := ""
	versionRuleTestPath := ""
	ServiceName := constants.ServiceChannelService

	testCases := []struct {
		name                string
		f                   func(*sling.Sling, string, string) *ChannelService
		client              *sling.Sling
		uriTemplate         string
		versionRuleTestPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, versionRuleTestPath},
		{"EmptyURITemplate", ServiceFunction, client, "", versionRuleTestPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", versionRuleTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.versionRuleTestPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
