package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/stretchr/testify/require"
)

func createGlobalRootService(t *testing.T) *globalRootService {
	service := newGlobalRootService(nil, TestURIRoot)
	services.testNewService(t, service, TestURIRoot, ServiceRootService)
	return service
}

func BenchmarkRootServiceGet(b *testing.B) {
	newGlobalRootService(nil, TestURIRoot).Get()
}

func TestRootServiceGet(t *testing.T) {
	service := createGlobalRootService(t)
	require.NotNil(t, service)

	resource, err := service.Get()
	require.NoError(t, err)
	require.NotNil(t, resource)
}

func TestRootServiceNew(t *testing.T) {
	ServiceFunction := newGlobalRootService
	client := &Client{}
	uriTemplate := services.emptyString
	ServiceName := ServiceRootService

	testCases := []struct {
		name        string
		f           func(*client, string) *globalRootService
		client      *client
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
