package tenants

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantService(t *testing.T) {
	ServiceFunction := NewTenantService
	client := &sling.Sling{}
	uriTemplate := ""
	missingVariablesPath := ""
	statusPath := ""
	tagTestPath := ""
	ServiceName := constants.ServiceTenantService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *TenantService
		client               *sling.Sling
		uriTemplate          string
		missingVariablesPath string
		statusPath           string
		tagTestPath          string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, missingVariablesPath, statusPath, tagTestPath},
		{"EmptyURITemplate", ServiceFunction, client, "", missingVariablesPath, statusPath, tagTestPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", missingVariablesPath, statusPath, tagTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.missingVariablesPath, tc.statusPath, tc.tagTestPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := NewTenantService(nil, "", constants.TestURITenantsMissingVariables, constants.TestURITenantsStatus, constants.TestURITenantTagTest)

	resource, err := service.GetByID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)
}
