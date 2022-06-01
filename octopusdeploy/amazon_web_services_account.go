package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AmazonWebServicesAccount represents an Amazon Web Services (AWS) account.
type AmazonWebServicesAccount struct {
	AccessKey string          `json:"AccessKey" validate:"required"`
	SecretKey *SensitiveValue `json:"SecretKey" validate:"required"`

	account
}

// NewAmazonWebServicesAccount initializes and returns an AWS account with a name, access key, and secret key.
func NewAmazonWebServicesAccount(name string, accessKey string, secretKey *SensitiveValue) (*AmazonWebServicesAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if isEmpty(accessKey) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterAccessKey)
	}

	if secretKey == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterSecretKey)
	}

	account := AmazonWebServicesAccount{
		AccessKey: accessKey,
		SecretKey: secretKey,
		account:   *newAccount(name, AccountType("AmazonWebServicesAccount")),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AmazonWebServicesAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
