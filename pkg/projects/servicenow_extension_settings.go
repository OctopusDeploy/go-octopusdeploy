package projects

import (
	"encoding/json"

	ext "github.com/OctopusDeploy/go-octopusdeploy/v2/internal/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
)

type ServiceNowExtensionSettings struct {
	IsStateAutomaticallyTransitioned bool
	StandardChangeTemplateName       string

	ext.ConnectedChangeControlExtensionSettings
}

// NewServiceNowExtensionSettings creates an instance of extension settings for ServiceNow.
func NewServiceNowExtensionSettings(connectionID string, isChangeControlled bool, standardChangeTemplateName string, isStateAutomaticallyTransitioned bool) *ServiceNowExtensionSettings {
	return &ServiceNowExtensionSettings{
		IsStateAutomaticallyTransitioned:        isStateAutomaticallyTransitioned,
		StandardChangeTemplateName:              standardChangeTemplateName,
		ConnectedChangeControlExtensionSettings: ext.NewConnectedChangeControlExtensionSettings(extensions.ExtensionIDServiceNow, isChangeControlled, connectionID),
	}
}

func (s *ServiceNowExtensionSettings) ConnectionID() string {
	return s.ConnectedChangeControlExtensionSettings.ConnectionID
}

func (s *ServiceNowExtensionSettings) ExtensionID() extensions.ExtensionID {
	return s.ExtensionSettings.ExtensionID
}

func (s *ServiceNowExtensionSettings) IsChangeControlled() bool {
	return s.ChangeControlExtensionSettings.IsChangeControlled
}

func (s *ServiceNowExtensionSettings) SetConnectionID(connectionID string) {
	s.ConnectedChangeControlExtensionSettings.ConnectionID = connectionID
}

func (s *ServiceNowExtensionSettings) SetExtensionID(extensionID extensions.ExtensionID) {
	s.ExtensionSettings.ExtensionID = extensionID
}

func (s *ServiceNowExtensionSettings) SetIsChangeControlled(isChangeControlled bool) {
	s.ChangeControlExtensionSettings.IsChangeControlled = isChangeControlled
}

// MarshalJSON returns the ServiceNow extension settings as its JSON encoding.
func (s ServiceNowExtensionSettings) MarshalJSON() ([]byte, error) {
	extensionSettings := struct {
		ExtensionID extensions.ExtensionID `json:"ExtensionId"`
		Values      map[string]interface{} `json:"Values"`
	}{
		ExtensionID: s.ExtensionID(),
		Values: map[string]interface{}{
			"AutomaticStateTransition":   s.IsStateAutomaticallyTransitioned,
			"StandardChangeTemplateName": s.StandardChangeTemplateName,
			"ServiceNowChangeControlled": s.IsChangeControlled(),
			"ServiceNowConnectionId":     s.ConnectionID(),
		},
	}

	return json.Marshal(extensionSettings)
}

// UnmarshalJSON sets the ServiceNow extension settings to its representation in JSON.
func (s *ServiceNowExtensionSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		ExtensionID extensions.ExtensionID `json:"ExtensionId"`
		Values      map[string]interface{} `json:"Values"`
	}

	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	s.SetExtensionID(fields.ExtensionID)

	for k, v := range fields.Values {
		switch k {
		case "AutomaticStateTransition":
			s.IsStateAutomaticallyTransitioned = v.(bool)
		case "StandardChangeTemplateName":
			s.StandardChangeTemplateName = v.(string)
		case "ServiceNowChangeControlled":
			s.SetIsChangeControlled(v.(bool))
		case "ServiceNowConnectionId":
			s.SetConnectionID(v.(string))
		}
	}

	return nil
}

var _ extensions.ExtensionSettings = &ServiceNowExtensionSettings{}
var _ extensions.ChangeControlExtensionSettings = &ServiceNowExtensionSettings{}
var _ extensions.ConnectedChangeControlExtensionSettings = &ServiceNowExtensionSettings{}
