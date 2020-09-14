package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

func NewAwsServicePrincipalAccount(name string, accessKey string, secretKey SensitiveValue) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAwsServicePrincipalAccount", "name")
	}

	if isEmpty(accessKey) {
		return nil, createInvalidParameterError("NewAwsServicePrincipalAccount", "accessKey")
	}

	account := &Account{
		Name:        name,
		AccountType: enum.AmazonWebServicesAccount,
		AccessKey:   accessKey,
		SecretKey:   &secretKey,
	}

	return account, nil
}

func validateAwsServicePrincipalAccount(account *Account) error {
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
