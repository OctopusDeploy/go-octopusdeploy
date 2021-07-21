package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GoogleCloudAccount represents a Google cloud account.
type GoogleCloudAccount struct {
	JsonKey *SensitiveValue `validate:"required"`

	account
}

// NewGoogleCloudAccount initializes and returns a Google cloud account.
func NewGoogleCloudAccount(name string, jsonKey *SensitiveValue, options ...func(*GoogleCloudAccount)) (*GoogleCloudAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if jsonKey == nil {
		return nil, createRequiredParameterIsEmptyOrNilError("jsonKey")
	}

	account := GoogleCloudAccount{
		account: *newAccount(name, AccountType("GoogleCloudAccount")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.JsonKey = jsonKey
	account.AccountType = AccountType("GoogleCloudAccount")
	account.ID = emptyString
	account.ModifiedBy = emptyString
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
func (a *GoogleCloudAccount) Validate() error {
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
