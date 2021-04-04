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
	var spv SensitiveValue
	errUnmarshalSensitivePropertyValue := json.Unmarshal(data, &spv)

	if errUnmarshalSensitivePropertyValue != nil {
		var p string
		errUnmarshalString := json.Unmarshal(data, &p)

		if errUnmarshalString != nil {
			return errUnmarshalString
		}

		d.Value = p
		d.SensitiveValue = nil
		return nil
	}

	d.SensitiveValue = &spv
	return nil
}
