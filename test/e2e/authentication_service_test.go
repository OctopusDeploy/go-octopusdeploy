package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthenticationServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources, err := authentication.Get(client)
	require.NoError(t, err)
	require.NotNil(t, resources)
	require.NotNil(t, resources.AuthenticationProviders)

	for _, provider := range resources.AuthenticationProviders {
		assert.NoError(t, err)
		assert.NotEmpty(t, provider)
	}
}
