package subscriptions_test

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/subscriptions"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestNewSubscriptionSlackDefaults(t *testing.T) {
	s := subscriptions.NewSubscription("my-subscription")

	require.NotNil(t, s.EventNotificationSubscription.SlackChannelIds)
	require.Empty(t, s.EventNotificationSubscription.SlackChannelIds)
	require.NotNil(t, s.EventNotificationSubscription.SlackChannelNames)
	require.Empty(t, s.EventNotificationSubscription.SlackChannelNames)
	require.Equal(t, "01:00:00", s.EventNotificationSubscription.SlackFrequencyPeriod)
	require.Equal(t, "Summary", s.EventNotificationSubscription.SlackDigestFormat)
}

func TestEventNotificationSubscriptionSlackFieldsJSON(t *testing.T) {
	inputJSON := `{
		"EmailFrequencyPeriod": "01:00:00",
		"EmailPriority": "Normal",
		"EmailShowDatesInTimeZoneId": "UTC",
		"EmailTeams": [],
		"Filter": null,
		"SlackChannelIds": ["C0123", "C0456"],
		"SlackChannelNames": ["general", "releases"],
		"SlackFrequencyPeriod": "01:00:00",
		"SlackDigestFormat": "Detailed",
		"TeamsFrequencyPeriod": "",
		"TeamsChannels": null,
		"WebhookHeaderKey": "",
		"WebhookHeaderValue": "",
		"WebhookTeams": [],
		"WebhookTimeout": "00:00:10",
		"WebhookURI": ""
	}`

	var sub subscriptions.EventNotificationSubscription
	err := json.Unmarshal([]byte(inputJSON), &sub)
	require.NoError(t, err)
	require.Equal(t, []string{"C0123", "C0456"}, sub.SlackChannelIds)
	require.Equal(t, []string{"general", "releases"}, sub.SlackChannelNames)
	require.Equal(t, "01:00:00", sub.SlackFrequencyPeriod)
	require.Equal(t, "Detailed", sub.SlackDigestFormat)

	outputJSON, err := json.Marshal(sub)
	require.NoError(t, err)

	jsonassert.New(t).Assertf(inputJSON, "%s", string(outputJSON))
}

func TestNewSubscriptionTeamsDefaults(t *testing.T) {
	s := subscriptions.NewSubscription("my-subscription")

	require.NotNil(t, s.EventNotificationSubscription.TeamsChannels)
	require.Empty(t, s.EventNotificationSubscription.TeamsChannels)
	require.Equal(t, "01:00:00", s.EventNotificationSubscription.TeamsFrequencyPeriod)
}

func TestEventNotificationSubscriptionTeamsFieldsJSON(t *testing.T) {
	// The server never returns the URL value, WebhookUrl is HasValue: true with no NewValue.
	inputJSON := `{
		"TeamsChannels": [
			{"Id": "id-1", "Name": "general", "WebhookUrl": {"HasValue": true}},
			{"Id": "id-2", "Name": "releases", "WebhookUrl": {"HasValue": true}}
		],
		"TeamsFrequencyPeriod": "01:00:00"
	}`

	var sub subscriptions.EventNotificationSubscription
	err := json.Unmarshal([]byte(inputJSON), &sub)
	require.NoError(t, err)

	require.Len(t, sub.TeamsChannels, 2)
	require.Equal(t, "id-1", sub.TeamsChannels[0].Id)
	require.Equal(t, "general", sub.TeamsChannels[0].Name)
	require.NotNil(t, sub.TeamsChannels[0].WebhookUrl)
	require.True(t, sub.TeamsChannels[0].WebhookUrl.HasValue)
	require.Nil(t, sub.TeamsChannels[0].WebhookUrl.NewValue)
	require.Equal(t, "01:00:00", sub.TeamsFrequencyPeriod)
}

func TestEventNotificationSubscriptionTeamsAppChannelJSON(t *testing.T) {
	channelId := "19:abc123@thread.tacv2"
	teamId := "team-guid-1234"
	teamName := "Octopus Deploy"

	inputJSON := `{
		"TeamsChannels": [
			{"Id": "id-1", "Type": "AppChannel", "Name": "general", "ChannelId": "19:abc123@thread.tacv2", "TeamId": "team-guid-1234", "TeamName": "Octopus Deploy"}
		],
		"TeamsFrequencyPeriod": "01:00:00"
	}`

	var sub subscriptions.EventNotificationSubscription
	err := json.Unmarshal([]byte(inputJSON), &sub)
	require.NoError(t, err)

	require.Len(t, sub.TeamsChannels, 1)
	ch := sub.TeamsChannels[0]
	require.Equal(t, "id-1", ch.Id)
	require.Equal(t, "AppChannel", ch.Type)
	require.Equal(t, "general", ch.Name)
	require.Nil(t, ch.WebhookUrl)
	require.NotNil(t, ch.ChannelId)
	require.Equal(t, channelId, *ch.ChannelId)
	require.NotNil(t, ch.TeamId)
	require.Equal(t, teamId, *ch.TeamId)
	require.NotNil(t, ch.TeamName)
	require.Equal(t, teamName, *ch.TeamName)
}
