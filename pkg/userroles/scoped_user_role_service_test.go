package userroles

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createScopedUserRoleService(t *testing.T) *ScopedUserRoleService {
	service := NewScopedUserRoleService(nil, constants.TestURIScopedUserRoles)
	services.NewServiceTests(t, service, constants.TestURIScopedUserRoles, constants.ServiceScopedUserRoleService)
	return service
}

func TestScopedUserRoleServiceAddGetDelete(t *testing.T) {
	scopedUserRoleService := createScopedUserRoleService(t)
	require.NotNil(t, scopedUserRoleService)

	resource, err := scopedUserRoleService.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterScopedUserRole))
	assert.Nil(t, resource)

	invalidResource := &ScopedUserRole{}
	resource, err = scopedUserRoleService.Add(invalidResource)
	assert.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestScopedUserRoleServiceNew(t *testing.T) {
	ServiceFunction := NewScopedUserRoleService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceScopedUserRoleService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *ScopedUserRoleService
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
