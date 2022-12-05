package environments

import (
	"encoding/json"

	ext "github.com/OctopusDeploy/go-octopusdeploy/v2/internal/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
)

type ServiceNowExtensionSettings struct {
	ext.ChangeControlExtensionSettings
}

// NewServiceNowExtensionSettings creates an instance of extension settings for ServiceNow.
func NewServiceNowExtensionSettings(isChangeControlled bool) *ServiceNowExtensionSettings {
	return &ServiceNowExtensionSettings{
		ChangeControlExtensionSettings: ext.NewChangeControlExtensionSettings(extensions.ServiceNowExtensionID, isChangeControlled),
	}
}

func (s *ServiceNowExtensionSettings) ExtensionID() extensions.ExtensionID {
	return s.ExtensionSettings.ExtensionID
}

func (s *ServiceNowExtensionSettings) IsChangeControlled() bool {
	return s.ChangeControlExtensionSettings.IsChangeControlled
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
			"ServiceNowChangeControlled": s.IsChangeControlled(),
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
		case "ServiceNowChangeControlled":
			s.SetIsChangeControlled(v.(bool))
		}
	}

	return nil
}

var _ extensions.ExtensionSettings = &ServiceNowExtensionSettings{}
var _ extensions.ChangeControlExtensionSettings = &ServiceNowExtensionSettings{}
