package artifacts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func TestArtifactService(t *testing.T) {
	serviceName := constants.ServiceArtifactService
	uriTemplate := constants.TestURIArtifacts

	service := NewArtifactService(nil, uriTemplate)
	require.NotNil(t, service)
	require.NotNil(t, service.GetClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.GetURITemplate(), template)
	require.Equal(t, service.GetName(), serviceName)

	client := &sling.Sling{}

	testCases := []struct {
		name        string
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", nil, uriTemplate},
		{"EmptyURITemplate", client, ""},
		{"URITemplateWithWhitespace", client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewArtifactService(tc.client, tc.uriTemplate)
			require.NotNil(t, service)
			require.NotNil(t, service.GetClient())

			template, err := uritemplates.Parse(uriTemplate)
			require.NoError(t, err)
			require.Equal(t, service.GetURITemplate(), template)
			require.Equal(t, service.GetName(), serviceName)
		})
	}
}
