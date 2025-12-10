package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type PlatformHubUsernamePasswordAccount struct {
	Username string               `json:"Username" validate:"required"`
	Password *core.SensitiveValue `json:"Password" validate:"required"`

	platformHubAccount
}

func NewPlatformHubUsernamePasswordAccount(name string, username string, password *core.SensitiveValue) (*PlatformHubUsernamePasswordAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(username) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("username")
	}

	if password == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("password")
	}

	account := PlatformHubUsernamePasswordAccount{
		Username:           username,
		Password:           password,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubUsernamePasswordAccount),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *PlatformHubUsernamePasswordAccount) Validate() error {
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
