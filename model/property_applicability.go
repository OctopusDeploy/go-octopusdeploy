package model

type PropertyApplicability struct {
	DependsOnPropertyName  string      `json:"DependsOnPropertyName,omitempty"`
	DependsOnPropertyValue interface{} `json:"DependsOnPropertyValue,omitempty"`
	Mode                   string      `json:"Mode,omitempty" validate:"required,oneof=ApplicableIfHasAnyValue ApplicableIfHasNoValue ApplicableIfNotSpecificValue ApplicableIfSpecificValue"`
}
