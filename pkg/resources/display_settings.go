package resources

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ControlType string

const (
	ControlTypeSingleLineText     = ControlType("SingleLineText")
	ControlTypeMultiLineText      = ControlType("MultiLineText")
	ControlTypeSelect             = ControlType("Select")
	ControlTypeCheckbox           = ControlType("Checkbox")
	ControlTypeSensitive          = ControlType("Sensitive")
	ControlTypeStepName           = ControlType("StepName")
	ControlTypeCertificate        = ControlType("Certificate")
	ControlTypeWorkerPool         = ControlType("WorkerPool")
	ControlTypeAzureAccount       = ControlType("AzureAccount")
	ControlTypeGoogleCloudAccount = ControlType("GoogleCloudAccount")
	ControlTypeAwsAccount         = ControlType("AmazonWebServicesAccount")
)

type SelectOption struct {
	Value       string
	DisplayName string
}

type DisplaySettings struct {
	ControlType   ControlType     `json:"Octopus.ControlType"`
	SelectOptions []*SelectOption `json:"Octopus.SelectOptions,omitempty"`
}

func NewDisplaySettings(controlType ControlType, selectOptions []*SelectOption) *DisplaySettings {
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
		ControlType: string(d.ControlType),
	}

	for _, opt := range d.SelectOptions {
		displaySettings.SelectOptions += fmt.Sprintf("%s|%s\n", opt.Value, opt.DisplayName)
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

	d.ControlType = ControlType(fields.ControlType)

	var displaySettings map[string]*json.RawMessage
	if err := json.Unmarshal(b, &displaySettings); err != nil {
		return err
	}

	if displaySettings["Octopus.SelectOptions"] != nil {
		d.SelectOptions = make([]*SelectOption, 0)

		var selectOptionsDelimitedString *string
		if err := json.Unmarshal(*displaySettings["Octopus.SelectOptions"], &selectOptionsDelimitedString); err != nil {
			return err
		}

		for _, kv := range strings.Split(*selectOptionsDelimitedString, "\n") {
			pairs := strings.SplitN(kv, "|", 2)
			if len(pairs) == 2 { // ignore malformed options; server shouldn't send them anyway
				d.SelectOptions = append(d.SelectOptions, &SelectOption{Value: pairs[0], DisplayName: pairs[1]})
			}
		}
	}

	return nil
}
