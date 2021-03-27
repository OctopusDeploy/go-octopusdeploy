package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerStatusServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	serverStatus, err := client.ServerStatus.Get()
	require.NoError(t, err)
	require.NotNil(t, serverStatus)
}
