package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTagSetService(t *testing.T) {
	serviceFunction := newTagSetService
	client := &sling.Sling{}
	uriTemplate := emptyString
	sortOrderPath := emptyString
	serviceName := serviceTagSetService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string) *tagSetService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, sortOrderPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, sortOrderPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, sortOrderPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestTagSetServiceGetWithEmptyID(t *testing.T) {
	service := newTagSetService(nil, emptyString, TestURITagSetSortOrder)

	resource, err := service.GetByID(emptyString)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Error(t, err)
	assert.Nil(t, resource)
}
