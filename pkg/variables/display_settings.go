package variables

import (
	"encoding/json"
	"strings"
)

type DisplaySettings struct {
	ControlType   string            `json:"Octopus.ControlType"`
	SelectOptions map[string]string `json:"Octopus.SelectOptions,omitempty"`
}

func NewDisplaySettings(controlType string, selectOptions map[string]string) *DisplaySettings {
	return &DisplaySettings{
		ControlType:   controlType,
		SelectOptions: selectOptions,
	}
}

// MarshalJSON returns display settings as its JSON encoding.
func (d *DisplaySettings) MarshalJSON() ([]byte, error) {
	displaySettings := struct {
		ControlType   string `json:"Octopus.ControlType"`
		SelectOptions string `json:"Octopus.SelectOptions,omitempty"`
	}{
		ControlType: d.ControlType,
	}

	for k, v := range d.SelectOptions {
		displaySettings.SelectOptions += k + "|" + v + "\n"
	}

	displaySettings.SelectOptions = strings.TrimSuffix(displaySettings.SelectOptions, "\n")

	return json.Marshal(displaySettings)
}

// UnmarshalJSON sets display settings from its representation in JSON.
func (d *DisplaySettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		ControlType string `json:"Octopus.ControlType"`
	}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	d.ControlType = fields.ControlType

	var displaySettings map[string]*json.RawMessage
	if err := json.Unmarshal(b, &displaySettings); err != nil {
		return err
	}

	if displaySettings["Octopus.SelectOptions"] != nil {
		d.SelectOptions = map[string]string{}

		var selectOptionsDelimitedString *string
		if err := json.Unmarshal(*displaySettings["Octopus.SelectOptions"], &selectOptionsDelimitedString); err != nil {
			return err
		}

		for _, kv := range strings.Split(*selectOptionsDelimitedString, "\n") {
			pairs := strings.Split(kv, "|")
			d.SelectOptions[pairs[0]] = pairs[1]
		}
	}

	return nil
}
