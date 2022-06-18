package runbooks

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookSnapshotService(t *testing.T) *RunbookSnapshotService {
	service := NewRunbookSnapshotService(nil, constants.TestURIRunbookSnapshots)
	services.NewServiceTests(t, service, constants.TestURIRunbookSnapshots, constants.ServiceRunbookSnapshotService)
	return service
}

func TestRunbookSnapshotServiceAdd(t *testing.T) {
	service := createRunbookSnapshotService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbookSnapshot))
	assert.Nil(t, resource)

	invalidResource := &RunbookSnapshot{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestRunbookSnapshotServiceNew(t *testing.T) {
	ServiceFunction := NewRunbookSnapshotService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceRunbookSnapshotService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *RunbookSnapshotService
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
