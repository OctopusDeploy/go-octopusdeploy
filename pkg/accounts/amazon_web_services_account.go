package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AmazonWebServicesAccount represents an Amazon Web Services (AWS) account.
type AmazonWebServicesAccount struct {
	AccessKey string               `json:"AccessKey" validate:"required"`
	SecretKey *core.SensitiveValue `json:"SecretKey" validate:"required"`

	account
}

// NewAmazonWebServicesAccount initializes and returns an AWS account with a name, access key, and secret key.
func NewAmazonWebServicesAccount(name string, accessKey string, secretKey *core.SensitiveValue) (*AmazonWebServicesAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(accessKey) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterAccessKey)
	}

	if secretKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterSecretKey)
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
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
