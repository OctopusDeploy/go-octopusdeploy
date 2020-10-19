package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// TokenAccount represents a token account.
type TokenAccount struct {
	AccountType string          `json:"AccountType" validate:"required,eq=Token"`
	Token       *SensitiveValue `json:"Token,omitempty" validate:"required"`

	AccountResource
}

// NewTokenAccount creates and initializes a token account with a name and
// token.
func NewTokenAccount(name string, token SensitiveValue) *TokenAccount {
	return &TokenAccount{
		AccountType:     "Token",
		Token:           &token,
		AccountResource: *newAccountResource(name),
	}
}

// GetAccountType returns the account type for this account.
func (t *TokenAccount) GetAccountType() string {
	return t.AccountType
}

// Validate checks the state of this account and returns an error if invalid.
func (t *TokenAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(t)
}

var _ IAccount = &TokenAccount{}
