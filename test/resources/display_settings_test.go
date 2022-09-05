package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestDisplaySettings(t *testing.T) {
	controlType := variables.ControlTypeSelect
	option1 := internal.GetRandomName()
	value1 := internal.GetRandomName()
	option2 := internal.GetRandomName()
	value2 := internal.GetRandomName()

	selectOptions := []*variables.SelectOption{
		{Key: value1, Value: option1},
		{Key: value2, Value: option2},
	}

	displaySettings := variables.NewDisplaySettings(controlType, selectOptions)
	require.NotNil(t, displaySettings)
	require.Equal(t, controlType, displaySettings.ControlType)
	require.Equal(t, selectOptions, displaySettings.SelectOptions)
}

func TestDisplaySettingsAsJson(t *testing.T) {
	controlType := variables.ControlTypeSelect

	displaySettings := variables.NewDisplaySettings(controlType, []*variables.SelectOption{
		{Key: "Value-1", Value: "Option-1"},
		{Key: "Value-2", Value: "Option-2"},
		{Key: "Value-3", Value: "Option-3"},
	})
	require.NotNil(t, displaySettings)

	expectedJson := `{
		"Octopus.ControlType": "Select",
		"Octopus.SelectOptions":"Value-1|Option-1\nValue-2|Option-2\nValue-3|Option-3"
	}`

	displaySettingsAsJson, err := json.Marshal(displaySettings)
	require.NoError(t, err)
	require.NotNil(t, displaySettingsAsJson)

	jsonassert.New(t).Assertf(expectedJson, string(displaySettingsAsJson))
}

func TestSelectOptions(t *testing.T) {
	controlType := variables.ControlTypeSelect
	option1 := internal.GetRandomName()
	value1 := internal.GetRandomName()
	option2 := internal.GetRandomName()
	value2 := internal.GetRandomName()
	option3 := internal.GetRandomName()
	value3 := internal.GetRandomName()

	selectOptions := []*variables.SelectOption{
		{Key: value1, Value: option1},
		{Key: value2, Value: option2},
		{Key: value3, Value: option3},
	}
	displaySettings := variables.NewDisplaySettings(controlType, selectOptions)

	displaySettingsAsJson, err := json.Marshal(displaySettings)
	require.NoError(t, err)
	require.NotNil(t, displaySettingsAsJson)

	// TODO: loop through each select option; verify count and option/value pairs
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
	require.Equal(t, variables.ControlTypeSelect, displaySettings.ControlType)
	require.NotNil(t, displaySettings.SelectOptions)
	require.Len(t, displaySettings.SelectOptions, 3)
	require.Equal(t, []*variables.SelectOption{
		{Key: "Value-1", Value: "Option-1"},
		{Key: "Value-2", Value: "Option-2"},
		{Key: "Value-3", Value: "Option-3"},
	}, displaySettings.SelectOptions)
}
