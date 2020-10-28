package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createRootService(t *testing.T) *rootService {
	service := newRootService(nil, TestURIRoot)
	testNewService(t, service, TestURIRoot, ServiceRootService)
	return service
}

func BenchmarkRootServiceGet(b *testing.B) {
	newRootService(nil, TestURIRoot).Get()
}

func TestRootServiceGet(t *testing.T) {
	service := createRootService(t)
	require.NotNil(t, service)

	resource, err := service.Get()
	require.NoError(t, err)
	require.NotNil(t, resource)
}

func TestRootServiceNew(t *testing.T) {
	ServiceFunction := newRootService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceRootService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *rootService
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
