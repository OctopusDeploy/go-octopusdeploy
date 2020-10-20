package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AmazonWebServicesAccount represents an Amazon Web Services (AWS) account.
type AmazonWebServicesAccount struct {
	AccessKey string          `json:"AccessKey,omitempty" validate:"required"`
	SecretKey *SensitiveValue `json:"SecretKey,omitempty" validate:"required"`

	AccountResource
}

// NewAmazonWebServicesAccount initializes and returns an AWS account with a
// name, access key, and secret key.
func NewAmazonWebServicesAccount(name string, accessKey string, secretKey SensitiveValue) *AmazonWebServicesAccount {
	return &AmazonWebServicesAccount{
		AccessKey:       accessKey,
		SecretKey:       &secretKey,
		AccountResource: *newAccountResource(name, accountTypeAmazonWebServicesAccount),
	}
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AmazonWebServicesAccount) Validate() error {
	v := validator.New()
	v.RegisterStructValidation(validateAmazonWebServicesAccount, AmazonWebServicesAccount{})
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

func validateAmazonWebServicesAccount(sl validator.StructLevel) {
	account := sl.Current().Interface().(AmazonWebServicesAccount)
	if account.AccountType != accountTypeAmazonWebServicesAccount {
		sl.ReportError(account.AccountType, "AccountType", "AccountType", "accounttype", accountTypeSshKeyPair)
	}
}
