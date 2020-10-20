package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createRunbookProcessService(t *testing.T) *runbookProcessService {
	service := newRunbookProcessService(nil, TestURIRunbookProcesses)
	testNewService(t, service, TestURIRunbookProcesses, serviceRunbookProcessService)
	return service
}

func TestRunbookProcessServiceGetAll(t *testing.T) {
	service := createRunbookProcessService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.ID)
		require.NoError(t, err)
		require.NotNil(t, resourceToCompare)
	}
}

func TestRunbookProcessServiceNew(t *testing.T) {
	serviceFunction := newRunbookProcessService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceRunbookProcessService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *runbookProcessService
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
