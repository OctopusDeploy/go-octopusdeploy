package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// GenericOIDCAccount represents a Generic OIDC account.
type GenericOIDCAccount struct {
	Audience              string            `json:"Audience,omitempty"`
	DeploymentSubjectKeys []string          `json:"DeploymentSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type'"`
	CustomClaims          map[string]string `json:"CustomClaims,omitempty"`

	account
}

// NewGenericOIDCAccount creates and initializes a Generic OIDC account.
func NewGenericOIDCAccount(name string) (*GenericOIDCAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	account := GenericOIDCAccount{
		account: *newAccount(name, AccountTypeGenericOIDCAccount),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *GenericOIDCAccount) Validate() error {
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
