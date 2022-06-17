package octopusdeploy

import (
	"encoding/json"
)

type PropertyValue struct {
	IsSensitive    bool            `json:"IsSensitive,omitempty"`
	SensitiveValue *SensitiveValue `json:"SensitiveValue,omitempty"`
	Value          string          `json:"Value,omitempty"`
}

func NewPropertyValue(value string, isSensitive bool) PropertyValue {
	propertyValue := PropertyValue{
		IsSensitive: isSensitive,
	}

	if isSensitive {
		propertyValue.SensitiveValue = NewSensitiveValue(value)
	} else {
		propertyValue.Value = value
	}

	return propertyValue
}

func (p PropertyValue) MarshalJSON() ([]byte, error) {
	if p.IsSensitive {
		return json.Marshal(p.SensitiveValue)
	}

	return json.Marshal(p.Value)
}

// UnmarshalJSON sets this property value to its representation in JSON.
func (d *PropertyValue) UnmarshalJSON(data []byte) error {
	// try unmarshal into a sensitive property, if that fails, it's just a normal property
	var sensitiveValue SensitiveValue
	if err := json.Unmarshal(data, &sensitiveValue); err == nil {
		d.IsSensitive = true
		d.SensitiveValue = &sensitiveValue
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	d.Value = value
	d.SensitiveValue = nil
	return nil
}
