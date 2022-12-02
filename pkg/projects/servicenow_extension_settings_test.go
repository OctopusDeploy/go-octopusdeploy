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
	}`, extensions.ExtensionIDServiceNow, isStateAutomaticallyTransitioned, standardChangeTemplateName, isChangeControlled, connectionID)

	jsonassert.New(t).Assertf(expectedJson, string(serviceNowExtensionSettingsAsJSON))
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
	}`, extensions.ExtensionIDServiceNow, isStateAutomaticallyTransitioned, standardChangeTemplateName, isChangeControlled, connectionID)

	var serviceNowExtensionSettings projects.ServiceNowExtensionSettings
	err := json.Unmarshal([]byte(inputJSON), &serviceNowExtensionSettings)
	require.NoError(t, err)
	require.NotNil(t, serviceNowExtensionSettings)
	require.Equal(t, connectionID, serviceNowExtensionSettings.ConnectionID())
	require.Equal(t, isChangeControlled, serviceNowExtensionSettings.IsChangeControlled())
	require.Equal(t, standardChangeTemplateName, serviceNowExtensionSettings.StandardChangeTemplateName)
	require.Equal(t, isStateAutomaticallyTransitioned, serviceNowExtensionSettings.IsStateAutomaticallyTransitioned)
}
