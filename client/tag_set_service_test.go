package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTagSetService(t *testing.T) {
	serviceFunction := newTagSetService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceTagSetService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *tagSetService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestTagSetServiceGetWithEmptyID(t *testing.T) {
	service := newTagSetService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Error(t, err)
	assert.Nil(t, resource)
}
