package model

import (
	"github.com/go-playground/validator/v10"
)

type Variables struct {
	ID          string      `json:"Id"`
	OwnerID     string      `json:"OwnerId"`
	Version     int         `json:"Version"`
	Variables   []Variable  `json:"Variables"`
	ScopeValues ScopeValues `json:"ScopeValues"`
	Links       map[string]string
}

type Variable struct {
	ID          string                 `json:"Id"`
	Name        string                 `json:"Name"`
	Value       string                 `json:"Value"`
	Description string                 `json:"Description"`
	Scope       *VariableScope         `json:"Scope,omitempty"`
	IsEditable  bool                   `json:"IsEditable"`
	Prompt      *VariablePromptOptions `json:"Prompt"`
	Type        string                 `json:"Type"`
	IsSensitive bool                   `json:"IsSensitive"`
}

func (t *Variable) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewVariable(name, valuetype, value, description string, scope *VariableScope, sensitive bool) *Variable {
	return &Variable{
		Name:        name,
		Value:       value,
		Description: description,
		Type:        valuetype,
		IsSensitive: sensitive,
		Scope:       scope,
	}
}
