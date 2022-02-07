package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GoogleCloudPlatformAccount represents a Google cloud account.
type GoogleCloudPlatformAccount struct {
	JsonKey *SensitiveValue `validate:"required"`

	accounts.account
}

// NewGoogleCloudPlatformAccount initializes and returns a Google cloud account.
func NewGoogleCloudPlatformAccount(name string, jsonKey *SensitiveValue, options ...func(*GoogleCloudPlatformAccount)) (*GoogleCloudPlatformAccount, error) {
	if IsEmpty(name) {
		return nil, CreateRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if jsonKey == nil {
		return nil, CreateRequiredParameterIsEmptyOrNilError("jsonKey")
	}

	account := GoogleCloudPlatformAccount{
		account: *accounts.newAccount(name, accounts.AccountType("GoogleCloudAccount")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.JsonKey = jsonKey
	account.AccountType = accounts.AccountType("GoogleCloudAccount")
	account.ID = services.emptyString
	account.ModifiedBy = services.emptyString
	account.ModifiedOn = nil
	account.Name = name

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *GoogleCloudPlatformAccount) Validate() error {
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
