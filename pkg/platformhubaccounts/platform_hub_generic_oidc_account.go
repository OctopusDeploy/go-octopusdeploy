package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PlatformHubGenericOidcAccount represents a Platform Hub Generic OIDC account.
type PlatformHubGenericOidcAccount struct {
	ExecutionSubjectKeys []string `json:"ExecutionSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type"`
	Audience             string   `json:"Audience,omitempty"`

	platformHubAccount
}

// NewPlatformHubGenericOidcAccount initializes and returns a Platform Hub Generic OIDC account with a name.
func NewPlatformHubGenericOidcAccount(name string) (*PlatformHubGenericOidcAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	account := PlatformHubGenericOidcAccount{
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubGenericOidcAccount),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *PlatformHubGenericOidcAccount) Validate() error {
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
