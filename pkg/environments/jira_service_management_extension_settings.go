package environments

import (
	"encoding/json"

	ext "github.com/OctopusDeploy/go-octopusdeploy/v2/internal/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
)

type JiraServiceManagementExtensionSettings struct {
	ext.ChangeControlExtensionSettings
}

// NewJiraServiceManagementExtensionSettings creates an instance of extension settings for Jira Service Management (JSM).
func NewJiraServiceManagementExtensionSettings(isChangeControlled bool) *JiraServiceManagementExtensionSettings {
	return &JiraServiceManagementExtensionSettings{
		ChangeControlExtensionSettings: ext.NewChangeControlExtensionSettings(extensions.JiraServiceManagementExtensionID, isChangeControlled),
	}
}

func (j *JiraServiceManagementExtensionSettings) ExtensionID() extensions.ExtensionID {
	return j.ExtensionSettings.ExtensionID
}

func (j *JiraServiceManagementExtensionSettings) IsChangeControlled() bool {
	return j.ChangeControlExtensionSettings.IsChangeControlled
}

func (j *JiraServiceManagementExtensionSettings) SetExtensionID(extensionID extensions.ExtensionID) {
	j.ExtensionSettings.ExtensionID = extensionID
}

func (j *JiraServiceManagementExtensionSettings) SetIsChangeControlled(isChangeControlled bool) {
	j.ChangeControlExtensionSettings.IsChangeControlled = isChangeControlled
}

// MarshalJSON returns the Jira Service Management (JSM) extension settings as its JSON encoding.
func (j JiraServiceManagementExtensionSettings) MarshalJSON() ([]byte, error) {
	extensionSettings := struct {
		ExtensionID extensions.ExtensionID `json:"ExtensionId"`
		Values      map[string]interface{} `json:"Values"`
	}{
		ExtensionID: j.ExtensionID(),
		Values: map[string]interface{}{
			"JsmChangeControlled": j.IsChangeControlled(),
		},
	}

	return json.Marshal(extensionSettings)
}

// UnmarshalJSON sets the Jira Service Management (JSM) extension settings to its representation in JSON.
func (j *JiraServiceManagementExtensionSettings) UnmarshalJSON(b []byte) error {
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
		case "JsmChangeControlled":
			j.SetIsChangeControlled(v.(bool))
		}
	}

	return nil
}

var _ extensions.ExtensionSettings = &JiraServiceManagementExtensionSettings{}
var _ extensions.ChangeControlExtensionSettings = &JiraServiceManagementExtensionSettings{}
