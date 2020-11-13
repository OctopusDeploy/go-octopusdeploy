package octopusdeploy

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyMachine(t *testing.T) {
	machine := &machine{}
	require.NotNil(t, machine)
	require.Error(t, machine.Validate())
}

func TestKubernetesEndpoint(t *testing.T) {
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	machine := &machine{
		Endpoint: NewKubernetesEndpoint(url),
	}
	assert.NotNil(t, machine)
	assert.NoError(t, machine.Validate())
}
