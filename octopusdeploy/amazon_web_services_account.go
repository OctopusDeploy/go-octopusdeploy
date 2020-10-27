package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AmazonWebServicesAccount represents an Amazon Web Services (AWS) account.
type AmazonWebServicesAccount struct {
	AccessKey string          `validate:"required"`
	SecretKey *SensitiveValue `validate:"required"`

	account
}

// NewAmazonWebServicesAccount initializes and returns an AWS account with a name, access key, and secret key.
func NewAmazonWebServicesAccount(name string, accessKey string, secretKey *SensitiveValue, options ...func(*AmazonWebServicesAccount)) (*AmazonWebServicesAccount, error) {
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
		account: *newAccount(name, AccountType("AmazonWebServicesAccount")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccessKey = accessKey
	account.accountType = AccountType("AmazonWebServicesAccount")
	account.ID = emptyString
	account.ModifiedBy = emptyString
	account.ModifiedOn = nil
	account.Name = name
	account.SecretKey = secretKey

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
