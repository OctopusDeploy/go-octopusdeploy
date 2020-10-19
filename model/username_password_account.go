package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// UsernamePasswordAccount represents a username/password account.
type UsernamePasswordAccount struct {
	AccountType string          `json:"AccountType" validate:"required,eq=UsernamePassword"`
	Username    string          `json:"Username,omitempty"`
	Password    *SensitiveValue `json:"Password,omitempty"`

	AccountResource
}

// NewUsernamePasswordAccount creates and initializes a username/password
// account with a name.
func NewUsernamePasswordAccount(name string) *UsernamePasswordAccount {
	return &UsernamePasswordAccount{
		AccountType:     "UsernamePassword",
		AccountResource: *newAccountResource(name),
	}
}

// GetAccountType returns the account type for this account.
func (u *UsernamePasswordAccount) GetAccountType() string {
	return u.AccountType
}

// Validate checks the state of this account and returns an error if invalid.
func (u *UsernamePasswordAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(u)
}

var _ IAccount = &UsernamePasswordAccount{}
