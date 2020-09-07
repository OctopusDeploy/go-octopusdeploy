package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ActionTemplateParameter struct {
	DefaultValue    *PropertyValueResource `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string      `json:"DisplaySettings,omitempty"`
	HelpText        string                 `json:"HelpText,omitempty"`
	Label           string                 `json:"Label,omitempty"`
	Name            string                 `json:"Name,omitempty"`
	Resource
}

func (a *ActionTemplateParameter) GetID() string {
	return a.ID
}

func (a *ActionTemplateParameter) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &ActionTemplateParameter{}
