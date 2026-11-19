package variables

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type Variable struct {
	Description string                 `json:"Description"`
	IsEditable  bool                   `json:"IsEditable"`
	IsSensitive bool                   `json:"IsSensitive"`
	Name        string                 `json:"Name"`
	Prompt      *VariablePromptOptions `json:"Prompt,omitempty"`
	Scope       VariableScope          `json:"Scope"`
	Type        string                 `json:"Type"`
	// Value is never populated for a sensitive variable: the server returns null and it reads
	// back as an empty string. Preserving an existing secret across an update therefore depends
	// on sending the variable's ID, not on the value. Writing a sensitive variable with an empty
	// Value and no ID replaces the stored secret with an empty string.
	Value   string `json:"Value"`
	SpaceID string `json:"SpaceId,omitempty"`

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
