package runbooks

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRunbookService(t *testing.T) *RunbookService {
	service := NewRunbookService(nil, constants.TestURIRunbooks)
	services.NewServiceTests(t, service, constants.TestURIRunbooks, constants.ServiceRunbookService)
	return service
}

func TestRunbookServiceAddGetDelete(t *testing.T) {
	runbookService := createRunbookService(t)
	require.NotNil(t, runbookService)

	resource, err := runbookService.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbook))
	assert.Nil(t, resource)

	invalidResource := &Runbook{}
	resource, err = runbookService.Add(invalidResource)
	assert.Equal(t, internal.CreateValidationFailureError(constants.OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestRunbookServiceNew(t *testing.T) {
	ServiceFunction := NewRunbookService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceRunbookService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *RunbookService
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

func TestRunbookSnapshotVariablesValidation(t *testing.T) {
	_, err := SnapshotVariables(nil, "Spaces-1", "RunbookSnapshots-1")
	assert.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("client"), err)

	client := newclient.NewClient(&newclient.HttpSession{})

	_, err = SnapshotVariables(client, "", "RunbookSnapshots-1")
	assert.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID"), err)

	_, err = SnapshotVariables(client, "Spaces-1", "")
	assert.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("snapshotID"), err)
}
