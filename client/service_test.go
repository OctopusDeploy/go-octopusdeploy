package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestGetOperations(t *testing.T) {
	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *accountService
		uriTemplate string
	}{
		{"Accounts", newAccountService, TestURIAccounts},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(nil, tc.uriTemplate)

			assert.NotNil(t, service)
			if service == nil {
				return
			}
		})
	}
}

func testNewService(t *testing.T, service ServiceInterface, uriTemplate string, serviceName string) {
	assert := assert.New(t)

	assert.NotNil(service)
	assert.NotNil(service.getClient())

	template, err := uritemplates.Parse(uriTemplate)
	assert.NoError(err)

	assert.Equal(service.getURITemplate(), template)
	assert.Equal(service.getName(), serviceName)
}
