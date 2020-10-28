package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookService(t *testing.T) *runbookService {
	service := newRunbookService(nil, TestURIRunbooks)
	testNewService(t, service, TestURIRunbooks, ServiceRunbookService)
	return service
}

func TestRunbookServiceAddGetDelete(t *testing.T) {
	runbookService := createRunbookService(t)
	require.NotNil(t, runbookService)

	resource, err := runbookService.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterRunbook))
	assert.Nil(t, resource)

	invalidResource := &Runbook{}
	resource, err = runbookService.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestRunbookServiceNew(t *testing.T) {
	ServiceFunction := newRunbookService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceRunbookService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *runbookService
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
