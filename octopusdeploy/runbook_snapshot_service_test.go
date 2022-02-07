package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookSnapshotService(t *testing.T) *runbookSnapshotService {
	service := newRunbookSnapshotService(nil, TestURIRunbookSnapshots)
	services.testNewService(t, service, TestURIRunbookSnapshots, ServiceRunbookSnapshotService)
	return service
}

func TestRunbookSnapshotServiceAdd(t *testing.T) {
	service := createRunbookSnapshotService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterRunbookSnapshot))
	assert.Nil(t, resource)

	invalidResource := &RunbookSnapshot{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestRunbookSnapshotServiceNew(t *testing.T) {
	ServiceFunction := newRunbookSnapshotService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceRunbookSnapshotService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *runbookSnapshotService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
