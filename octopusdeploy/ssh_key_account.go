package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// SSHKeyAccount represents a SSH key pair account.
type SSHKeyAccount struct {
	PrivateKeyFile       *SensitiveValue `json:"PrivateKeyFile" validate:"required"`
	PrivateKeyPassphrase *SensitiveValue `json:"PrivateKeyPassphrase,omitempty"`
	Username             string          `json:"Username" validate:"required"`

	AccountResource
}

// NewSSHKeyAccount initializes and returns a SSH key pair account with a name,
// username, and private key file.
func NewSSHKeyAccount(name string, username string, privateKeyFile SensitiveValue) *SSHKeyAccount {
	return &SSHKeyAccount{
		Username:        username,
		PrivateKeyFile:  &privateKeyFile,
		AccountResource: *newAccountResource(name, accountTypeSshKeyPair),
	}
}

// Validate checks the state of this account and returns an error if invalid.
func (s *SSHKeyAccount) Validate() error {
	v := validator.New()
	v.RegisterStructValidation(validateSSHKeyAccount, SSHKeyAccount{})
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(s)
}

func validateSSHKeyAccount(sl validator.StructLevel) {
	account := sl.Current().Interface().(SSHKeyAccount)
	if account.AccountType != accountTypeSshKeyPair {
		sl.ReportError(account.AccountType, "AccountType", "AccountType", "accounttype", accountTypeSshKeyPair)
	}
}
