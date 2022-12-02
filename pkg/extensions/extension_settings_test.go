package extensions_test

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

func TestExtensionSettingsMarshalJSON(t *testing.T) {
	extensionSettings := []extensions.ExtensionSettings{}

	jsmConnectionID := internal.GetRandomName()
	jsmChangeControlled := false
	serviceDeskProjectName := internal.GetRandomName()
	jiraServiceManagementExtensionSettings := projects.NewJiraServiceManagementExtensionSettings(jsmConnectionID, jsmChangeControlled, serviceDeskProjectName)

	extensionSettings = append(extensionSettings, jiraServiceManagementExtensionSettings)

	serviceNowConnectionID := internal.GetRandomName()
	serviceNowChangeControlled := false
	standardChangeTemplateName := internal.GetRandomName()
	isStateAutomaticallyTransitioned := false
	serviceNowExtensionSettings := projects.NewServiceNowExtensionSettings(serviceNowConnectionID, serviceNowChangeControlled, standardChangeTemplateName, isStateAutomaticallyTransitioned)

	extensionSettings = append(extensionSettings, serviceNowExtensionSettings)

	serviceNowExtensionSettingsAsJSON, err := json.Marshal(extensionSettings)
	require.NoError(t, err)
	require.NotNil(t, serviceNowExtensionSettingsAsJSON)

	expectedJson := fmt.Sprintf(`[
		{
			"ExtensionId": "%s",
			"Values": {
				"JsmChangeControlled": %v,
				"JsmConnectionId": "%s",
				"ServiceDeskProjectName": "%s"
			}
		},
		{
			"ExtensionId": "%s",
			"Values": {
				"AutomaticStateTransition": %v,
				"ServiceNowChangeControlled": %v,
				"ServiceNowConnectionId": "%s",
				"StandardChangeTemplateName": "%s"
			}
		}
	]`,
		extensions.ExtensionIDJiraServiceManagement,
		jsmChangeControlled,
		jsmConnectionID,
		serviceDeskProjectName,
		extensions.ExtensionIDServiceNow,
		isStateAutomaticallyTransitioned,
		serviceNowChangeControlled,
		serviceNowConnectionID,
		standardChangeTemplateName,
	)

	jsonassert.New(t).Assertf(expectedJson, string(serviceNowExtensionSettingsAsJSON))
}
