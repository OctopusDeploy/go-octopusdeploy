package model

import (
	"github.com/go-playground/validator/v10"
)

type ConfigurationSections struct {
	Items []ConfigurationSection `json:"Items"`
	PagedResults
}

type ConfigurationSection struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name,omitempty"`

	Resource
}

func (c *ConfigurationSection) GetID() string {
	return c.ID
}

func (c *ConfigurationSection) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &ConfigurationSection{}
