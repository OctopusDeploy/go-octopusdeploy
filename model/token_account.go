package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// TokenAccount represents a token account.
type TokenAccount struct {
	Token *SensitiveValue `json:"Token,omitempty" validate:"required"`

	AccountResource
}

// NewTokenAccount creates and initializes a token account with a name and
// token.
func NewTokenAccount(name string, token SensitiveValue) *TokenAccount {
	return &TokenAccount{
		Token:           &token,
		AccountResource: *newAccountResource(name, accountTypeToken),
	}
}

// Validate checks the state of this account and returns an error if invalid.
func (t *TokenAccount) Validate() error {
	v := validator.New()
	v.RegisterStructValidation(validateTokenAccount, TokenAccount{})
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(t)
}

func validateTokenAccount(sl validator.StructLevel) {
	account := sl.Current().Interface().(TokenAccount)
	if account.AccountType != accountTypeToken {
		sl.ReportError(account.AccountType, "AccountType", "AccountType", "accounttype", accountTypeSshKeyPair)
	}
}
