package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AmazonWebServicesAccount represents an Amazon Web Services (AWS) account.
type AmazonWebServicesAccount struct {
	AccountType string          `json:"AccountType" validate:"required,eq=AmazonWebServicesAccount"`
	AccessKey   string          `json:"AccessKey,omitempty" validate:"required"`
	SecretKey   *SensitiveValue `json:"SecretKey,omitempty" validate:"required"`

	AccountResource
}

// NewAmazonWebServicesAccount initializes and returns an AWS account with a
// name, access key, and secret key.
func NewAmazonWebServicesAccount(name string, accessKey string, secretKey SensitiveValue) *AmazonWebServicesAccount {
	return &AmazonWebServicesAccount{
		AccountType:     "AmazonWebServicesAccount",
		AccessKey:       accessKey,
		SecretKey:       &secretKey,
		AccountResource: *newAccountResource(name),
	}
}

// GetAccountType returns the account type for this account.
func (a *AmazonWebServicesAccount) GetAccountType() string {
	return a.AccountType
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AmazonWebServicesAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &AmazonWebServicesAccount{}
