package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GoogleCloudPlatformAccount represents a Google cloud accountV1.
type GoogleCloudPlatformAccount struct {
	JsonKey *resources.SensitiveValue `validate:"required"`

	Account
}

// NewGoogleCloudPlatformAccount initializes and returns a Google cloud accountV1.
func NewGoogleCloudPlatformAccount(name string, jsonKey *resources.SensitiveValue, options ...func(*GoogleCloudPlatformAccount)) (*GoogleCloudPlatformAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
	}

	if jsonKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("jsonKey")
	}

	account := GoogleCloudPlatformAccount{
		Account: *NewAccount(name, AccountType("GoogleCloudAccount")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.JsonKey = jsonKey
	account.AccountType = AccountType("GoogleCloudAccount")
	account.Name = name

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this accountV1 and returns an error if invalid.
func (a *GoogleCloudPlatformAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
