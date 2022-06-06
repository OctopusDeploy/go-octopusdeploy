package octopusdeploy

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/stretchr/testify/require"
)

func testNewService(t *testing.T, service IService, uriTemplate string, ServiceName string) {
	require.NotNil(t, service)
	require.NotNil(t, service.getClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.getURITemplate(), template)
	require.Equal(t, service.getName(), ServiceName)
}
