package octopusdeploy

type Variable struct {
	Description string                 `json:"Description"`
	IsEditable  bool                   `json:"IsEditable"`
	IsSensitive bool                   `json:"IsSensitive"`
	Name        string                 `json:"Name"`
	Prompt      *VariablePromptOptions `json:"Prompt,omitempty"`
	Scope       VariableScope          `json:"Scope"`
	Type        string                 `json:"Type"`
	Value       string                 `json:"Value"`

	resource
}

func NewVariable(name string) *Variable {
	return &Variable{
		IsEditable:  true,
		IsSensitive: false,
		Name:        name,
		Type:        "String",

		resource: *newResource(),
	}
}
