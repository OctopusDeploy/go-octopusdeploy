package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// UsernamePasswordAccount represents a username/password account.
type UsernamePasswordAccount struct {
	Username string          `json:"Username,omitempty"`
	Password *SensitiveValue `json:"Password,omitempty"`

	AccountResource
}

// NewUsernamePasswordAccount creates and initializes a username/password
// account with a name.
func NewUsernamePasswordAccount(name string) *UsernamePasswordAccount {
	return &UsernamePasswordAccount{
		AccountResource: *newAccountResource(name, accountTypeUsernamePassword),
	}
}

// Validate checks the state of this account and returns an error if invalid.
func (u *UsernamePasswordAccount) Validate() error {
	v := validator.New()
	v.RegisterStructValidation(validateUsernamePasswordAccount, UsernamePasswordAccount{})
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(u)
}

func validateUsernamePasswordAccount(sl validator.StructLevel) {
	account := sl.Current().Interface().(UsernamePasswordAccount)
	if account.AccountType != accountTypeUsernamePassword {
		sl.ReportError(account.AccountType, "AccountType", "AccountType", "accounttype", accountTypeSshKeyPair)
	}
}
