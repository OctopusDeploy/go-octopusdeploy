package variables

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type Variable struct {
	Description string                 `json:"Description"`
	IsEditable  bool                   `json:"IsEditable"`
	IsSensitive bool                   `json:"IsSensitive"`
	Name        string                 `json:"Name"`
	Prompt      *VariablePromptOptions `json:"Prompt,omitempty"`
	Scope       VariableScope          `json:"Scope"`
	Type        string                 `json:"Type"`
	Value       string                 `json:"Value"`

	resources.Resource
}

func NewVariable(name string) *Variable {
	return &Variable{
		IsEditable:  true,
		IsSensitive: false,
		Name:        name,
		Type:        "String",

		Resource: *resources.NewResource(),
	}
}
