package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTenantVariablesServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	allTenantVariables, err := client.TenantVariables.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTenantVariables)
	require.True(t, len(allTenantVariables) > 0)

	for k, v := range allTenantVariables {
		require.NotNil(t, k)
		require.NotNil(t, v)
	}
}
