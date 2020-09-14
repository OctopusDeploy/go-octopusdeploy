package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

func NewSSHKeyAccount(name string, username string, privateKey SensitiveValue) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewSSHKeyAccount", "name")
	}

	if isEmpty(username) {
		return nil, createInvalidParameterError("NewSSHKeyAccount", "username")
	}

	return &Account{
		Name:           name,
		Username:       username,
		AccountType:    enum.SshKeyPair,
		PrivateKeyFile: &privateKey,
	}, nil
}

func validateSSHKeyAccount(account *Account) error {
	if account == nil {
		return createInvalidParameterError("validateSSHKeyAccount", "account")
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
