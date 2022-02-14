package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventsService(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := services.EventsQuery{}

	events, err := client.Events.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	agents, err := client.Events.GetAgents()
	assert.NoError(t, err)
	assert.NotNil(t, agents)

	categories, err := client.Events.GetCategories(services.EventCategoriesQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	documentTypes, err := client.Events.GetDocumentTypes()
	assert.NoError(t, err)
	assert.NotNil(t, documentTypes)

	groups, err := client.Events.GetGroups(services.EventGroupsQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, groups)
}
