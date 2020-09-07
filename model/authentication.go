package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Authentication struct {
	AnyAuthenticationProvidersSupportPasswordManagement bool                             `json:"AnyAuthenticationProvidersSupportPasswordManagement"`
	AuthenticationProviders                             []*AuthenticationProviderElement `json:"AuthenticationProviders"`
	AutoLoginEnabled                                    bool                             `json:"AutoLoginEnabled"`
	Resource
}

func (a *Authentication) GetID() string {
	return a.ID
}

func (a *Authentication) Validate() error {
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

var _ ResourceInterface = &Authentication{}
