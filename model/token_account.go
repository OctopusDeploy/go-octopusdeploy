package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

func NewTokenAccount(name string, token SensitiveValue) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewTokenAccount", "name")
	}

	return &Account{
		Name:        name,
		Token:       &token,
		AccountType: enum.Token,
	}, nil
}

func validateTokenAccount(account *Account) error {
	if account == nil {
		return createInvalidParameterError("validateTokenAccount", "account")
	}

	validate := validator.New()
	err := validate.Struct(account)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}
