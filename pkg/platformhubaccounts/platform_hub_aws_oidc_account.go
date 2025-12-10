package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PlatformHubAwsOIDCAccount represents a Platform Hub AWS OIDC account.
type PlatformHubAwsOIDCAccount struct {
	RoleArn                string   `json:"RoleArn" validate:"required"`
	SessionDuration        string   `json:"SessionDuration,omitempty"`
	DeploymentSubjectKeys  []string `json:"DeploymentSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type'"`
	HealthCheckSubjectKeys []string `json:"HealthCheckSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account target type'"`
	AccountTestSubjectKeys []string `json:"AccountTestSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account type'"`

	platformHubAccount
}

// NewPlatformHubAwsOIDCAccount initializes and returns a Platform Hub AWS OIDC account with a name and role ARN.
func NewPlatformHubAwsOIDCAccount(name string, roleArn string) (*PlatformHubAwsOIDCAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(roleArn) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("roleArn")
	}

	account := PlatformHubAwsOIDCAccount{
		RoleArn:            roleArn,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubAwsOIDCAccount),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *PlatformHubAwsOIDCAccount) Validate() error {
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
