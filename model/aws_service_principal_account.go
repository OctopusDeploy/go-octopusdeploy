package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

// NewAwsServicePrincipalAccount initializes and returns an AWS service
// principal account with a name, access key, and secret key. If any of the
// input parameters are invalid, it will return nil and an error.
func NewAwsServicePrincipalAccount(name string, accessKey string, secretKey SensitiveValue) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAwsServicePrincipalAccount", "name")
	}

	if isEmpty(accessKey) {
		return nil, createInvalidParameterError("NewAwsServicePrincipalAccount", "accessKey")
	}

	account, err := NewAccount(name, enum.AmazonWebServicesAccount)
	if err != nil {
		return nil, err
	}

	account.AccessKey = accessKey
	account.SecretKey = &secretKey

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
