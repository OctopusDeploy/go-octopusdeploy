package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

func NewUsernamePasswordAccount(name string) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewUsernamePasswordAccount", "name")
	}

	account := &Account{
		Name:        name,
		AccountType: enum.UsernamePassword,
	}

	return account, nil
}

func validateUsernamePasswordAccount(account *Account) error {
	if account == nil {
		return createInvalidParameterError("validateUsernamePasswordAccount", "account")
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
