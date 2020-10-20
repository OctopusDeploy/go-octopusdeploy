package octopusdeploy

type Variables struct {
	ID          string      `json:"Id"`
	OwnerID     string      `json:"OwnerId"`
	Version     int         `json:"Version"`
	Variables   []Variable  `json:"Variables"`
	ScopeValues ScopeValues `json:"ScopeValues"`
	Links       map[string]string
}

type Variable struct {
	Description string                 `json:"Description"`
	ID          string                 `json:"Id"`
	IsEditable  bool                   `json:"IsEditable"`
	IsSensitive bool                   `json:"IsSensitive"`
	Name        string                 `json:"Name"`
	Prompt      *VariablePromptOptions `json:"Prompt"`
	Scope       *VariableScope         `json:"Scope,omitempty"`
	Type        string                 `json:"Type"`
	Value       string                 `json:"Value"`

	Resource
}

func NewVariable(name, valuetype, value, description string, scope *VariableScope, sensitive bool) *Variable {
	return &Variable{
		Name:        name,
		Value:       value,
		Description: description,
		Type:        valuetype,
		IsSensitive: sensitive,
		Scope:       scope,
		Resource:    *newResource(),
	}
}
