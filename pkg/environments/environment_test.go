package environments_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentExtensionSettingsAsJSON(t *testing.T) {
	jiraEnvironmentType := internal.GetRandomName()
	name := internal.GetRandomName()
	slug := internal.GetRandomName()
	spaceID := internal.GetRandomName()

	expectedJSON := fmt.Sprintf(`{
		"AllowDynamicInfrastructure": true,
		"ExtensionSettings": [
			{
				"ExtensionId": "%s",
				"Values": {
					"JiraEnvironmentType": "%s"
				}
			},
			{
				"ExtensionId": "%s",
				"Values": {
					"JsmChangeControlled": true
				}
			},
			{
				"ExtensionId": "%s",
				"Values": {
					"ServiceNowChangeControlled": true
				}
			}
		],
		"Name": "%s",
		"Slug": "%s",
		"SortOrder": 0,
		"SpaceId": "%s",
		"UseGuidedFailure": true
	}`,
		extensions.JiraExtensionID,
		jiraEnvironmentType,
		extensions.JiraServiceManagementExtensionID,
		extensions.ServiceNowExtensionID,
		name,
		slug,
		spaceID,
	)

	var environment environments.Environment
	err := json.Unmarshal([]byte(expectedJSON), &environment)
	require.NoError(t, err)
	require.NotNil(t, environment)
	require.Len(t, environment.ExtensionSettings, 3)

	actualJSON, err := json.Marshal(environment)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))
}
