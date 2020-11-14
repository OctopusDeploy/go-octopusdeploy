package octopusdeploy

type Variable struct {
	Description string                 `json:"Description"`
	IsEditable  bool                   `json:"IsEditable"`
	IsSensitive bool                   `json:"IsSensitive"`
	Name        string                 `json:"Name"`
	Prompt      *VariablePromptOptions `json:"Prompt"`
	Scope       *VariableScope         `json:"Scope,omitempty"`
	Type        string                 `json:"Type"`
	Value       string                 `json:"Value"`

	resource
}

func NewVariable(name string, valueType string, value string, description string, scope *VariableScope, isSensitive bool) *Variable {
	return &Variable{
		Description: description,
		IsSensitive: isSensitive,
		Name:        name,
		Scope:       scope,
		Type:        valueType,
		Value:       value,

		resource: *newResource(),
	}
}
