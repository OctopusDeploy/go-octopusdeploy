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

func TestNewServiceNowExtensionSettings(t *testing.T) {
	var connectionID string
	var isChangeControlled bool
	var standardChangeTemplateName string
	var isStateAutomaticallyTransitioned bool

	serviceNowExtensionSettings := projects.NewServiceNowExtensionSettings(connectionID, isChangeControlled, standardChangeTemplateName, isStateAutomaticallyTransitioned)
	require.NotNil(t, serviceNowExtensionSettings)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, serviceNowExtensionSettings.IsChangeControlled())
	require.Equal(t, standardChangeTemplateName, serviceNowExtensionSettings.StandardChangeTemplateName)
	require.Equal(t, isStateAutomaticallyTransitioned, serviceNowExtensionSettings.IsStateAutomaticallyTransitioned)

	connectionID = internal.GetRandomName()
	isChangeControlled = false
	standardChangeTemplateName = internal.GetRandomName()
	isStateAutomaticallyTransitioned = false

	serviceNowExtensionSettings = projects.NewServiceNowExtensionSettings(connectionID, isChangeControlled, standardChangeTemplateName, isStateAutomaticallyTransitioned)
	require.NotNil(t, serviceNowExtensionSettings)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, serviceNowExtensionSettings.IsChangeControlled())
	require.Equal(t, standardChangeTemplateName, serviceNowExtensionSettings.StandardChangeTemplateName)
	require.Equal(t, isStateAutomaticallyTransitioned, serviceNowExtensionSettings.IsStateAutomaticallyTransitioned)
}

func TestServiceNowExtensionSettingsMarshalJSON(t *testing.T) {
	connectionID := internal.GetRandomName()
	isChangeControlled := false
	standardChangeTemplateName := internal.GetRandomName()
	isStateAutomaticallyTransitioned := false
	serviceNowExtensionSettings := projects.NewServiceNowExtensionSettings(connectionID, isChangeControlled, standardChangeTemplateName, isStateAutomaticallyTransitioned)

	serviceNowExtensionSettingsAsJSON, err := json.Marshal(serviceNowExtensionSettings)
	require.NoError(t, err)
	require.NotNil(t, serviceNowExtensionSettingsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"AutomaticStateTransition": %v,
			"StandardChangeTemplateName": "%s",
			"ServiceNowChangeControlled": %v,
			"ServiceNowConnectionId": "%s"
		}
	}`, extensions.ServiceNowExtensionID, isStateAutomaticallyTransitioned, standardChangeTemplateName, isChangeControlled, connectionID)

	jsonassert.New(t).Assertf(expectedJson, "%s", string(serviceNowExtensionSettingsAsJSON))
}

func TestServiceNowExtensionSettingsMarshalJSONWithoutStandardChangeTemplateName(t *testing.T) {
	connectionID := internal.GetRandomName()
	serviceNowExtensionSettings := projects.NewServiceNowExtensionSettings(connectionID, true, "", false)

	serviceNowExtensionSettingsAsJSON, err := json.Marshal(serviceNowExtensionSettings)
	require.NoError(t, err)

	expectedJson := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"AutomaticStateTransition": false,
			"StandardChangeTemplateName": "",
			"ServiceNowChangeControlled": true,
			"ServiceNowConnectionId": "%s"
		}
	}`, extensions.ServiceNowExtensionID, connectionID)

	jsonassert.New(t).Assertf(expectedJson, "%s", string(serviceNowExtensionSettingsAsJSON))
}

func TestServiceNowExtensionSettingsUnmarshalJSON(t *testing.T) {
	connectionID := internal.GetRandomName()
	isChangeControlled := false
	standardChangeTemplateName := internal.GetRandomName()
	isStateAutomaticallyTransitioned := false

	inputJSON := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"AutomaticStateTransition": %v,
			"StandardChangeTemplateName": "%s",
			"ServiceNowChangeControlled": %v,
			"ServiceNowConnectionId": "%s"
		}
	}`, extensions.ServiceNowExtensionID, isStateAutomaticallyTransitioned, standardChangeTemplateName, isChangeControlled, connectionID)

	var serviceNowExtensionSettings projects.ServiceNowExtensionSettings
	err := json.Unmarshal([]byte(inputJSON), &serviceNowExtensionSettings)
	require.NoError(t, err)
	require.NotNil(t, serviceNowExtensionSettings)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, serviceNowExtensionSettings.IsChangeControlled())
	require.Equal(t, standardChangeTemplateName, serviceNowExtensionSettings.StandardChangeTemplateName)
	require.Equal(t, isStateAutomaticallyTransitioned, serviceNowExtensionSettings.IsStateAutomaticallyTransitioned)
}

func TestServiceNowExtensionSettingsUnmarshalJSONWithNullStandardChangeTemplateName(t *testing.T) {
	connectionID := internal.GetRandomName()

	inputJSON := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"AutomaticStateTransition": false,
			"StandardChangeTemplateName": null,
			"ServiceNowChangeControlled": true,
			"ServiceNowConnectionId": "%s"
		}
	}`, extensions.ServiceNowExtensionID, connectionID)

	var serviceNowExtensionSettings projects.ServiceNowExtensionSettings
	err := json.Unmarshal([]byte(inputJSON), &serviceNowExtensionSettings)
	require.NoError(t, err)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.True(t, serviceNowExtensionSettings.IsChangeControlled())
	require.Empty(t, serviceNowExtensionSettings.StandardChangeTemplateName)
}

func TestServiceNowExtensionSettingsUnmarshalJSONWithMissingStandardChangeTemplateName(t *testing.T) {
	connectionID := internal.GetRandomName()

	inputJSON := fmt.Sprintf(`{
		"ExtensionId": "%s",
		"Values": {
			"AutomaticStateTransition": false,
			"ServiceNowChangeControlled": true,
			"ServiceNowConnectionId": "%s"
		}
	}`, extensions.ServiceNowExtensionID, connectionID)

	var serviceNowExtensionSettings projects.ServiceNowExtensionSettings
	err := json.Unmarshal([]byte(inputJSON), &serviceNowExtensionSettings)
	require.NoError(t, err)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.True(t, serviceNowExtensionSettings.IsChangeControlled())
	require.Empty(t, serviceNowExtensionSettings.StandardChangeTemplateName)
}
