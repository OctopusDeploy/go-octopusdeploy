package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createRunbookProcessService(t *testing.T) *runbookProcessService {
	service := newRunbookProcessService(nil, TestURIRunbookProcesses)
	services.testNewService(t, service, TestURIRunbookProcesses, ServiceRunbookProcessService)
	return service
}

func TestRunbookProcessServiceGetAll(t *testing.T) {
	service := createRunbookProcessService(t)
	require.NotNil(t, service)

	runbookProcesses, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, runbookProcesses)

	for _, runbookProcess := range runbookProcesses {
		runbookProcessToCompare, err := service.GetByID(runbookProcess.GetID())
		require.NoError(t, err)
		require.NotNil(t, runbookProcessToCompare)
	}
}

func TestRunbookProcessServiceNew(t *testing.T) {
	ServiceFunction := newRunbookProcessService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceRunbookProcessService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *runbookProcessService
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
