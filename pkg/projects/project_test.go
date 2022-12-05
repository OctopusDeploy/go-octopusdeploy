package projects_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestProjectExtensionSettingsAsJSON(t *testing.T) {
	jsmConnectionID := internal.GetRandomName()
	lifecycleID := internal.GetRandomName()
	name := internal.GetRandomName()
	projectGroupID := internal.GetRandomName()
	serviceDeskProjectName := internal.GetRandomName()
	serviceNowConnectionID := internal.GetRandomName()
	standardChangeTemplateName := internal.GetRandomName()

	expectedJSON := fmt.Sprintf(`{
		"ExtensionSettings": [
			{
				"ExtensionId": "%s",
				"Values": {
					"JsmChangeControlled": true,
					"JsmConnectionId": "%s",
					"ServiceDeskProjectName": "%s"
				}
			},
			{
				"ExtensionId": "servicenow-integration",
				"Values": {
					"AutomaticStateTransition": true,
					"ServiceNowChangeControlled": true,
					"ServiceNowConnectionId": "%s",
					"StandardChangeTemplateName": "%s"
				}
			}
		],
		"LifecycleId": "%s",
		"Name": "%s",
		"ProjectGroupId": "%s"
	}`,
		extensions.JiraServiceManagementExtensionID,
		jsmConnectionID,
		serviceDeskProjectName,
		serviceNowConnectionID,
		standardChangeTemplateName,
		lifecycleID,
		name,
		projectGroupID,
	)

	var project projects.Project
	err := json.Unmarshal([]byte(expectedJSON), &project)
	require.NoError(t, err)
	require.NotNil(t, project)
	require.Len(t, project.ExtensionSettings, 2)
	require.Equal(t, lifecycleID, project.LifecycleID)
	require.Equal(t, name, project.Name)
	require.Equal(t, projectGroupID, project.ProjectGroupID)

	actualJSON, err := json.Marshal(project)
	require.NoError(t, err)
	require.NotNil(t, actualJSON)

	jsonassert.New(t).Assertf(expectedJSON, string(actualJSON))
}
