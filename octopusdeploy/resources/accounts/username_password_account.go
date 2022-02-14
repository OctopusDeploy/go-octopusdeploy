package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// UsernamePasswordAccount represents a username/password account.
type UsernamePasswordAccount struct {
	Username string
	Password *resources.SensitiveValue

	account
}

// IUsernamePasswordAccount defines the interface for username-password accounts.
type IUsernamePasswordAccount interface {
	GetUsername() string
	SetPassword(*resources.SensitiveValue)
	SetUsername(string)

	IAccount
}

// GetUsername returns the username of this username/password account.
func (u *UsernamePasswordAccount) GetUsername() string {
	return u.Username
}

// SetPassword sets the password of this username/password account.
func (u *UsernamePasswordAccount) SetPassword(password *resources.SensitiveValue) {
	u.Password = password
}

// SetUsername sets the username of this username/password account.
func (u *UsernamePasswordAccount) SetUsername(username string) {
	u.Username = username
}

// NewUsernamePasswordAccount creates and initializes a username/password account with a name.
func NewUsernamePasswordAccount(name string, options ...func(*UsernamePasswordAccount)) (*UsernamePasswordAccount, error) {
	if octopusdeploy.IsEmpty(name) {
		return nil, octopusdeploy.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
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
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(u)
}

var _ IUsernamePasswordAccount = &UsernamePasswordAccount{}
