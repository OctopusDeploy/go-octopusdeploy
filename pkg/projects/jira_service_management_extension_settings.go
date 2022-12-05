package projects

import (
	"encoding/json"

	ext "github.com/OctopusDeploy/go-octopusdeploy/v2/internal/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
)

type JiraServiceManagementExtensionSettings struct {
	ServiceDeskProjectName string

	ext.ConnectedChangeControlExtensionSettings
}

// NewJiraServiceManagementExtensionSettings creates an instance of extension settings for Jira Service Management (JSM).
func NewJiraServiceManagementExtensionSettings(connectionID string, isChangeControlled bool, serviceDeskProjectName string) *JiraServiceManagementExtensionSettings {
	return &JiraServiceManagementExtensionSettings{
		ServiceDeskProjectName:                  serviceDeskProjectName,
		ConnectedChangeControlExtensionSettings: ext.NewConnectedChangeControlExtensionSettings(extensions.JiraServiceManagementExtensionID, isChangeControlled, connectionID),
	}
}

func (j *JiraServiceManagementExtensionSettings) ConnectionID() string {
	return j.ConnectedChangeControlExtensionSettings.ConnectionID
}

func (j *JiraServiceManagementExtensionSettings) ExtensionID() extensions.ExtensionID {
	return j.ExtensionSettings.ExtensionID
}

func (j *JiraServiceManagementExtensionSettings) IsChangeControlled() bool {
	return j.ChangeControlExtensionSettings.IsChangeControlled
}

func (j *JiraServiceManagementExtensionSettings) SetConnectionID(connectionID string) {
	j.ConnectedChangeControlExtensionSettings.ConnectionID = connectionID
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
			"JsmChangeControlled":    j.IsChangeControlled(),
			"JsmConnectionId":        j.ConnectionID(),
			"ServiceDeskProjectName": j.ServiceDeskProjectName,
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
		case "JsmConnectionId":
			j.SetConnectionID(v.(string))
		case "ServiceDeskProjectName":
			j.ServiceDeskProjectName = v.(string)
		}
	}

	return nil
}

var _ extensions.ExtensionSettings = &JiraServiceManagementExtensionSettings{}
var _ extensions.ChangeControlExtensionSettings = &JiraServiceManagementExtensionSettings{}
var _ extensions.ConnectedChangeControlExtensionSettings = &JiraServiceManagementExtensionSettings{}
