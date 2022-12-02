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

func TestNewJiraServiceManagementExtensionSettings(t *testing.T) {
	var connectionID string
	var isChangeControlled bool
	var serviceDeskProjectName string

	jiraServiceManagementExtensionSettings := projects.NewJiraServiceManagementExtensionSettings(connectionID, isChangeControlled, serviceDeskProjectName)
	require.NotNil(t, jiraServiceManagementExtensionSettings)
	require.Equal(t, connectionID, jiraServiceManagementExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, jiraServiceManagementExtensionSettings.IsChangeControlled())
	require.Equal(t, serviceDeskProjectName, jiraServiceManagementExtensionSettings.ServiceDeskProjectName)

	connectionID = internal.GetRandomName()
	isChangeControlled = false
	serviceDeskProjectName = internal.GetRandomName()

	jiraServiceManagementExtensionSettings = projects.NewJiraServiceManagementExtensionSettings(connectionID, isChangeControlled, serviceDeskProjectName)
	require.NotNil(t, jiraServiceManagementExtensionSettings)
	require.Equal(t, connectionID, jiraServiceManagementExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, jiraServiceManagementExtensionSettings.IsChangeControlled())
	require.Equal(t, serviceDeskProjectName, jiraServiceManagementExtensionSettings.ServiceDeskProjectName)
}

func TestJiraServiceManagementExtensionSettingsMarshalJSON(t *testing.T) {
	connectionID := internal.GetRandomName()
	isChangeControlled := false
	serviceDeskProjectName := internal.GetRandomName()
	jiraServiceManagementExtensionSettings := projects.NewJiraServiceManagementExtensionSettings(connectionID, isChangeControlled, serviceDeskProjectName)

	jiraServiceManagementExtensionSettingsAsJSON, err := json.Marshal(jiraServiceManagementExtensionSettings)
	require.NoError(t, err)
	require.NotNil(t, jiraServiceManagementExtensionSettingsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"JsmChangeControlled": %v,
			"JsmConnectionId": "%s",
			"ServiceDeskProjectName": "%s"
		}
	}`, extensions.ExtensionIDJiraServiceManagement, isChangeControlled, connectionID, serviceDeskProjectName)

	jsonassert.New(t).Assertf(expectedJson, string(jiraServiceManagementExtensionSettingsAsJSON))
}

func TestJiraServiceManagementExtensionSettingsUnmarshalJSON(t *testing.T) {
	connectionID := internal.GetRandomName()
	isChangeControlled := false
	serviceDeskProjectName := internal.GetRandomName()

	inputJSON := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"JsmConnectionId": "%s",
			"JsmChangeControlled": %v,
			"ServiceDeskProjectName": "%s"
		}
	}`, extensions.ExtensionIDServiceNow, connectionID, isChangeControlled, serviceDeskProjectName)

	var j projects.JiraServiceManagementExtensionSettings
	err := json.Unmarshal([]byte(inputJSON), &j)
	require.NoError(t, err)
	require.NotNil(t, j)
	require.Equal(t, connectionID, j.ConnectionID())
	require.Equal(t, isChangeControlled, j.IsChangeControlled())
	require.Equal(t, serviceDeskProjectName, j.ServiceDeskProjectName)
}
