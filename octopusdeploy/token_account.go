package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// TokenAccount represents a token account.
type TokenAccount struct {
	Token *SensitiveValue `json:"Token,omitempty" validate:"required"`

	account
}

// NewTokenAccount creates and initializes a token account with a name and
// token.
func NewTokenAccount(name string, token *SensitiveValue, options ...func(*TokenAccount)) (*TokenAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if token == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterToken)
	}

	account := TokenAccount{
		Token:   token,
		account: *newAccount(name, AccountType("Token")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("Token")
	account.ID = emptyString
	account.ModifiedBy = emptyString
	account.ModifiedOn = nil
	account.Name = name
	account.Token = token

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
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
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(t)
}
