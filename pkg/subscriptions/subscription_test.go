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

	jsonassert.New(t).Assertf(inputJSON, string(outputJSON))
}
