package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// TokenAccount represents a token account.
type TokenAccount struct {
	Token *core.SensitiveValue `json:"Token,omitempty" validate:"required"`

	account
}

// NewTokenAccount creates and initializes a token account with a name and
// token.
func NewTokenAccount(name string, token *core.SensitiveValue) (*TokenAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	if token == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("token")
	}

	account := TokenAccount{
		Token:   token,
		account: *newAccount(name, AccountType("Token")),
	}

	// validate to ensure that all expectations are met
	if err := account.Validate(); err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (t *TokenAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(t)
}
