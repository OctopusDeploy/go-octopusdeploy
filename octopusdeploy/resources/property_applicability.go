package resources

import "github.com/go-playground/validator/v10"

type PropertyApplicability struct {
	DependsOnPropertyName  string      `json:"DependsOnPropertyName,omitempty"`
	DependsOnPropertyValue interface{} `json:"DependsOnPropertyValue,omitempty"`
	Mode                   string      `json:"Mode,omitempty" validate:"required,oneof=ApplicableIfHasAnyValue ApplicableIfHasNoValue ApplicableIfNotSpecificValue ApplicableIfSpecificValue"`
}

func NewPropertyApplicability() *PropertyApplicability {
	return &PropertyApplicability{}
}

// Validate checks the state of the property applicability and returns an error
// if invalid.
func (p PropertyApplicability) Validate() error {
	return validator.New().Struct(p)
}
