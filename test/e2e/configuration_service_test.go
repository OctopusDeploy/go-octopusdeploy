package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfiguration(t *testing.T) {
	octopusClient := getOctopusClient()

	configurationSection, err := octopusClient.Configuration.GetByID("authentication")

	assert.NoError(t, err)
	assert.NotNil(t, configurationSection)

	if err != nil || configurationSection == nil {
		return
	}

	assert.NotEmpty(t, configurationSection.Links)
	assert.NotEmpty(t, configurationSection.Links["Values"])
	assert.NotEmpty(t, configurationSection.Links["Metadata"])
}
