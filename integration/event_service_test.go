package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventsService(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := service.EventsQuery{}

	events, err := client.Events.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	agents, err := client.Events.GetAgents()
	assert.NoError(t, err)
	assert.NotNil(t, agents)

	categories, err := client.Events.GetCategories(service.EventCategoriesQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	documentTypes, err := client.Events.GetDocumentTypes()
	assert.NoError(t, err)
	assert.NotNil(t, documentTypes)

	groups, err := client.Events.GetGroups(service.EventGroupsQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, groups)
}
