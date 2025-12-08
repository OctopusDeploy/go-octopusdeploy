package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PlatformHubAwsAccount represents a Platform Hub AWS account.
type PlatformHubAwsAccount struct {
	AccessKey string               `json:"AccessKey" validate:"required"`
	SecretKey *core.SensitiveValue `json:"SecretKey" validate:"required"`

	platformHubAccount
}

// NewPlatformHubAwsAccount initializes and returns a Platform Hub AWS account with a name, access key, and secret key.
func NewPlatformHubAwsAccount(name string, accessKey string, secretKey *core.SensitiveValue) (*PlatformHubAwsAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(accessKey) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterAccessKey)
	}

	if secretKey == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterSecretKey)
	}

	account := PlatformHubAwsAccount{
		AccessKey:          accessKey,
		SecretKey:          secretKey,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubAwsAccount),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *PlatformHubAwsAccount) Validate() error {
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
