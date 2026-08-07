package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/serverstatus"
	"github.com/stretchr/testify/require"
)

func TestServerStatusServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	serverStatus, err := client.ServerStatus.Get()
	require.NoError(t, err)
	require.NotNil(t, serverStatus)
}

func TestServerStatusGetServerStatus(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	serverStatus, err := serverstatus.GetServerStatus(client)
	require.NoError(t, err)
	require.NotNil(t, serverStatus)
	require.NotEmpty(t, serverStatus.Links)
}

func TestServerStatusGetHealthStatus(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	healthStatus, err := serverstatus.GetHealthStatus(client)
	require.NoError(t, err)
	require.NotNil(t, healthStatus)
	require.NotEmpty(t, healthStatus.Description)
}

func TestServerStatusGetTimezones(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	timezones, err := serverstatus.GetTimezones(client)
	require.NoError(t, err)
	require.NotEmpty(t, timezones)

	for _, timezone := range timezones {
		require.NotEmpty(t, timezone.GetID())
		require.NotEmpty(t, timezone.Name)
	}
}

func TestServerStatusGetDocumentCounts(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	documentCounts, err := serverstatus.GetDocumentCounts(client)
	require.NoError(t, err)
	require.NotNil(t, documentCounts)
	require.GreaterOrEqual(t, documentCounts.Global.Spaces, 1)
}
