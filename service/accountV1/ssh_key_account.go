package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// SSHKeyAccount represents a SSH key pair accountV1.
type SSHKeyAccount struct {
	PrivateKeyFile       *resources.SensitiveValue `validate:"required"`
	PrivateKeyPassphrase *resources.SensitiveValue
	Username             string `validate:"required"`

	Account
}

// ISSHKeyAccount defines the interface for SSH key accounts.
type ISSHKeyAccount interface {
	SetPrivateKeyPassphrase(*resources.SensitiveValue)

	IAccount
}

// NewSSHKeyAccount initializes and returns a SSH key pair accountV1 with a name,
// username, and private key file.
func NewSSHKeyAccount(name string, username string, privateKeyFile *resources.SensitiveValue, options ...func(*SSHKeyAccount)) (*SSHKeyAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
	}

	if internal.IsEmpty(username) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterUsername)
	}

	if privateKeyFile == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterPrivateKeyFile)
	}

	account := SSHKeyAccount{
		Account: *NewAccount(name, AccountType("SshKeyPair")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("SshKeyPair")
	account.PrivateKeyFile = privateKeyFile
	account.Name = name
	account.Username = username

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this accountV1 and returns an error if invalid.
func (s *SSHKeyAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(s)
}
