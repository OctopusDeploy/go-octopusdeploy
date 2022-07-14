package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventsService(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := events.EventsQuery{}

	eventsFromQuery, err := client.Events.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, eventsFromQuery)

	agents, err := client.Events.GetAgents()
	assert.NoError(t, err)
	assert.NotNil(t, agents)

	categories, err := client.Events.GetCategories(events.EventCategoriesQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, categories)

	documentTypes, err := client.Events.GetDocumentTypes()
	assert.NoError(t, err)
	assert.NotNil(t, documentTypes)

	groups, err := client.Events.GetGroups(events.EventGroupsQuery{})
	assert.NoError(t, err)
	assert.NotNil(t, groups)
}
