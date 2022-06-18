package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// ISSHKeyAccount defines the interface for SSH key accounts.
type ISSHKeyAccount interface {
	SetPrivateKeyPassphrase(*core.SensitiveValue)

	IAccount
}

// SSHKeyAccount represents a SSH key pair account.
type SSHKeyAccount struct {
	PrivateKeyFile       *core.SensitiveValue `validate:"required"`
	PrivateKeyPassphrase *core.SensitiveValue
	Username             string `validate:"required"`

	account
}

// NewSSHKeyAccount initializes and returns a SSH key pair account with a name,
// username, and private key file.
func NewSSHKeyAccount(name string, username string, privateKeyFile *core.SensitiveValue) (*SSHKeyAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(username) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUsername)
	}

	if privateKeyFile == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterPrivateKeyFile)
	}

	account := SSHKeyAccount{
		PrivateKeyFile: privateKeyFile,
		Username:       username,
		account:        *newAccount(name, AccountType("SshKeyPair")),
	}

	// validate to ensure that all expectations are met
	if err := account.Validate(); err != nil {
		return nil, err
	}

	return &account, nil
}

// SetPrivateKeyPassphrase sets the private key [assphrase of this SSH key pair account.
func (s *SSHKeyAccount) SetPrivateKeyPassphrase(privateKeyPassphrase *core.SensitiveValue) {
	s.PrivateKeyPassphrase = privateKeyPassphrase
}

// Validate checks the state of this account and returns an error if invalid.
func (s *SSHKeyAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(s)
}

var _ ISSHKeyAccount = &SSHKeyAccount{}
