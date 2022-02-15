package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

const TokenAccountType string = "Token"

// TokenAccount represents a token accountV1.
type TokenAccount struct {
	Token *resources.SensitiveValue `json:"Token,omitempty" validate:"required"`

	Account
}

// NewTokenAccount creates and initializes a token accountV1 with a name and
// token.
func NewTokenAccount(name string, token *resources.SensitiveValue, options ...func(*TokenAccount)) (*TokenAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
	}

	if token == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterToken)
	}

	account := TokenAccount{
		Token:   token,
		Account: *NewAccount(name, AccountType(TokenAccountType)),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType(TokenAccountType)
	account.Name = name
	account.Token = token

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this accountV1 and returns an error if invalid.
func (t *TokenAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(t)
}
