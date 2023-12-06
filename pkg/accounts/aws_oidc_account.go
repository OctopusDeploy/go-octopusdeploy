package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AwsOIDCAccount represents an AWS OIDC account.
type AwsOIDCAccount struct {
	RoleArn                string   `json:"RoleArn"`
	SessionDuration        string   `json:"SessionDuration,omitempty"`
	Audience               string   `json:"Audience,omitempty"`
	DeploymentSubjectKeys  []string `json:"DeploymentSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type'"`
	HealthCheckSubjectKeys []string `json:"HealthCheckSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account target type'"`
	AccountTestSubjectKeys []string `json:"AccountTestSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account type'"`

	account
}

// NewAwsOIDCAccount creates and initializes an Aws OIDC account.
func NewAwsOIDCAccount(name string, roleArn string) (*AwsOIDCAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	account := AwsOIDCAccount{
		RoleArn: roleArn,
		account: *newAccount(name, AccountTypeAwsOIDC),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AwsOIDCAccount) Validate() error {
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
