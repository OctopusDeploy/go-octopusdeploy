package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// IUsernamePasswordAccount defines the interface for username-password accounts.
type IUsernamePasswordAccount interface {
	GetUsername() string
	SetPassword(*core.SensitiveValue)
	SetUsername(string)

	IAccount
}

// UsernamePasswordAccount represents a username/password account.
type UsernamePasswordAccount struct {
	Username string
	Password *core.SensitiveValue

	account
}

// GetUsername returns the username of this username/password account.
func (u *UsernamePasswordAccount) GetUsername() string {
	return u.Username
}

// SetPassword sets the password of this username/password account.
func (u *UsernamePasswordAccount) SetPassword(password *core.SensitiveValue) {
	u.Password = password
}

// SetUsername sets the username of this username/password account.
func (u *UsernamePasswordAccount) SetUsername(username string) {
	u.Username = username
}

// NewUsernamePasswordAccount creates and initializes a username/password account with a name.
func NewUsernamePasswordAccount(name string, options ...func(*UsernamePasswordAccount)) (*UsernamePasswordAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	account := UsernamePasswordAccount{
		account: *newAccount(name, AccountType("UsernamePassword")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("UsernamePassword")
	account.ID = ""
	account.ModifiedBy = ""
	account.ModifiedOn = nil
	account.Name = name

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (u *UsernamePasswordAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(u)
}

var _ IUsernamePasswordAccount = &UsernamePasswordAccount{}
