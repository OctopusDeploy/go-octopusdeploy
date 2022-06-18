package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createRootService(t *testing.T) *RootService {
	service := NewRootService(nil, constants.TestURIRoot)
	services.NewServiceTests(t, service, constants.TestURIRoot, constants.ServiceRootService)
	return service
}

func BenchmarkRootServiceGet(b *testing.B) {
	NewRootService(nil, constants.TestURIRoot).Get()
}

func TestRootServiceGet(t *testing.T) {
	service := createRootService(t)
	require.NotNil(t, service)

	resource, err := service.Get()
	require.NoError(t, err)
	require.NotNil(t, resource)
}

func TestRootServiceNew(t *testing.T) {
	ServiceFunction := NewRootService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceRootService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *RootService
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
