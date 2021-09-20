package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventsService(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := octopusdeploy.EventsQuery{}

	events, err := client.Events.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	agents, err := client.Events.GetAgents()
	assert.NoError(t, err)
	assert.NotNil(t, agents)

	categories, err := client.Events.GetCategories(octopusdeploy.EventCategoriesQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	documentTypes, err := client.Events.GetDocumentTypes()
	assert.NoError(t, err)
	assert.NotNil(t, documentTypes)

	groups, err := client.Events.GetGroups(octopusdeploy.EventGroupsQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, groups)
}
