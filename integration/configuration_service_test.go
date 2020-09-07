package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestGetConfiguration(t *testing.T) {
	configurationSection, err := octopusClient.Configuration.Get("authentication")

	assert.NoError(t, err)
	assert.NotNil(t, configurationSection)

	if err != nil || configurationSection == nil {
		return
	}

	assert.NotEmpty(t, configurationSection.Links)
	assert.NotEmpty(t, configurationSection.Links["Values"])
	assert.NotEmpty(t, configurationSection.Links["Metadata"])
}

func TestGetAllConfiguration(t *testing.T) {
	configurationSections, err := octopusClient.Configuration.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, configurationSections)

	if err != nil || configurationSections == nil {
		return
	}

	assert.NotEmpty(t, configurationSections.Items)

	for _, configurationSection := range configurationSections.Items {
		assert.NotEmpty(t, configurationSection.Links["Metadata"])
		assert.NotEmpty(t, configurationSection.Links["Values"])
	}
}
