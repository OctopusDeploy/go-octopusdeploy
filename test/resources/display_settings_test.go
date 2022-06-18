package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/variables"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestDisplaySettings(t *testing.T) {
	controlType := "Select"
	option1 := internal.GetRandomName()
	value1 := internal.GetRandomName()
	option2 := internal.GetRandomName()
	value2 := internal.GetRandomName()

	selectOptions := map[string]string{}
	selectOptions[value1] = option1
	selectOptions[value2] = option2

	displaySettings := variables.NewDisplaySettings(controlType, selectOptions)
	require.NotNil(t, displaySettings)
	require.Equal(t, controlType, displaySettings.ControlType)
	require.Equal(t, selectOptions, displaySettings.SelectOptions)
}

func TestDisplaySettingsAsJson(t *testing.T) {
	controlType := "Select"

	displaySettings := variables.NewDisplaySettings(controlType, nil)
	require.NotNil(t, displaySettings)

	expectedJson := `{
		"Octopus.ControlType": "Select"
	}`

	displaySettingsAsJson, err := json.Marshal(displaySettings)
	require.NoError(t, err)
	require.NotNil(t, displaySettingsAsJson)

	jsonassert.New(t).Assertf(expectedJson, string(displaySettingsAsJson))

	option1 := "Option-1"
	value1 := "Value-1"
	option2 := "Option-2"
	value2 := "Value-2"
	option3 := "Option-3"
	value3 := "Value-3"

	selectOptions := map[string]string{}
	selectOptions[value1] = option1
	selectOptions[value2] = option2
	selectOptions[value3] = option3

	displaySettings = variables.NewDisplaySettings(controlType, selectOptions)

	expectedJson = fmt.Sprintf(`{
		"Octopus.ControlType": "%s",
		"Octopus.SelectOptions": "Value-1|Option-1\nValue-2|Option-2\nValue-3|Option-3"
	}`, controlType)

	displaySettingsAsJson, err = json.Marshal(displaySettings)
	require.NoError(t, err)
	require.NotNil(t, displaySettingsAsJson)

	jsonassert.New(t).Assertf(string(displaySettingsAsJson), expectedJson)
}

func TestDisplaySettingsFromJson(t *testing.T) {
	controlType := "Select"

	displaySettingsAsJson := fmt.Sprintf(`{
		"Octopus.ControlType": "%s",
		"Octopus.SelectOptions": "Value-1|Option-1\nValue-2|Option-2\nValue-3|Option-3"
	}`, controlType)

	var displaySettings variables.DisplaySettings
	err := json.Unmarshal([]byte(displaySettingsAsJson), &displaySettings)
	require.NoError(t, err)
	require.NotNil(t, displaySettings)
	require.NotNil(t, displaySettings.ControlType)
	require.Equal(t, controlType, displaySettings.ControlType)
	require.NotNil(t, displaySettings.SelectOptions)
	require.Len(t, displaySettings.SelectOptions, 3)
	require.Equal(t, "Option-1", displaySettings.SelectOptions["Value-1"])
	require.Equal(t, "Option-2", displaySettings.SelectOptions["Value-2"])
	require.Equal(t, "Option-3", displaySettings.SelectOptions["Value-3"])
}
