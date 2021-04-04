package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createScopedUserRoleService(t *testing.T) *scopedUserRoleService {
	service := newScopedUserRoleService(nil, TestURIScopedUserRoles)
	testNewService(t, service, TestURIScopedUserRoles, ServiceScopedUserRoleService)
	return service
}

func TestScopedUserRoleServiceAddGetDelete(t *testing.T) {
	scopedUserRoleService := createScopedUserRoleService(t)
	require.NotNil(t, scopedUserRoleService)

	resource, err := scopedUserRoleService.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterScopedUserRole))
	assert.Nil(t, resource)

	invalidResource := &ScopedUserRole{}
	resource, err = scopedUserRoleService.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestScopedUserRoleServiceNew(t *testing.T) {
	ServiceFunction := newScopedUserRoleService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceScopedUserRoleService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *scopedUserRoleService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
