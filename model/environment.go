package model

import (
	"github.com/go-playground/validator/v10"
)

type Environments struct {
	Items []Environment `json:"Items"`
	PagedResults
}

type Environment struct {
	ID                         string `json:"Id"`
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	SortOrder                  int    `json:"SortOrder"`
	UseGuidedFailure           bool   `json:"UseGuidedFailure"`
	AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`
}

func (e *Environment) GetID() string {
	return e.ID
}

func (e *Environment) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)

	if err != nil {
		return err
	}

	return nil
}

func NewEnvironment(name, description string, useguidedfailure bool) *Environment {
	return &Environment{
		Name:             name,
		Description:      description,
		UseGuidedFailure: useguidedfailure,
	}
}

var _ ResourceInterface = &Environment{}
