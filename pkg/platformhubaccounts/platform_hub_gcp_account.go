package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PlatformHubGcpAccount represents a Platform Hub GCP account.
type PlatformHubGcpAccount struct {
	JsonKey *core.SensitiveValue `json:"JsonKey" validate:"required"`

	platformHubAccount
}

// NewPlatformHubGcpAccount initializes and returns a Platform Hub GCP account with a name and JSON key.
func NewPlatformHubGcpAccount(name string, jsonKey *core.SensitiveValue) (*PlatformHubGcpAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if jsonKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("jsonKey")
	}

	account := PlatformHubGcpAccount{
		JsonKey:            jsonKey,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubGcpAccount),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *PlatformHubGcpAccount) Validate() error {
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
