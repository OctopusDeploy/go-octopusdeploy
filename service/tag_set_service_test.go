package service

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTagSetService(t *testing.T) {
	ServiceFunction := newTagSetService
	client := &sling.Sling{}
	uriTemplate := emptyString
	sortOrderPath := emptyString
	ServiceName := ServiceTagSetService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string) *tagSetService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, sortOrderPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, sortOrderPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, sortOrderPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath)
			testNewService(t, service, uriTemplate, ServiceName)
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
