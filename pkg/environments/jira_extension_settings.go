package environments

import (
	"encoding/json"

	ext "github.com/OctopusDeploy/go-octopusdeploy/v2/internal/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
)

type JiraExtensionSettings struct {
	JiraEnvironmentType string
	ext.ExtensionSettings
}

// NewJiraExtensionSettings creates an instance of extension settings for Jira.
func NewJiraExtensionSettings(environmentType string) *JiraExtensionSettings {
	return &JiraExtensionSettings{
		JiraEnvironmentType: environmentType,
		ExtensionSettings:   ext.NewExtensionSettings(extensions.ExtensionIDJira),
	}
}

func (j *JiraExtensionSettings) ExtensionID() extensions.ExtensionID {
	return j.ExtensionSettings.ExtensionID
}

func (j *JiraExtensionSettings) SetExtensionID(extensionID extensions.ExtensionID) {
	j.ExtensionSettings.ExtensionID = extensionID
}

// MarshalJSON returns the Jira extension settings as its JSON encoding.
func (j JiraExtensionSettings) MarshalJSON() ([]byte, error) {
	extensionSettings := struct {
		ExtensionID extensions.ExtensionID `json:"ExtensionId"`
		Values      map[string]interface{} `json:"Values"`
	}{
		ExtensionID: j.ExtensionID(),
		Values: map[string]interface{}{
			"JiraEnvironmentType": j.JiraEnvironmentType,
		},
	}

	return json.Marshal(extensionSettings)
}

// UnmarshalJSON sets the Jira extension settings to its representation in JSON.
func (j *JiraExtensionSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		ExtensionID extensions.ExtensionID `json:"ExtensionId"`
		Values      map[string]interface{} `json:"Values"`
	}

	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	j.SetExtensionID(fields.ExtensionID)

	for k, v := range fields.Values {
		switch k {
		case "JiraEnvironmentType":
			j.JiraEnvironmentType = v.(string)
		}
	}

	return nil
}

var _ extensions.ExtensionSettings = &JiraExtensionSettings{}
