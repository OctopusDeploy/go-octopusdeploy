package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// SSHKeyAccount represents a SSH key pair account.
type SSHKeyAccount struct {
	AccountType          string          `json:"AccountType" validate:"required,eq=SshKeyPair"`
	PrivateKeyFile       *SensitiveValue `json:"PrivateKeyFile" validate:"required"`
	PrivateKeyPassphrase *SensitiveValue `json:"PrivateKeyPassphrase,omitempty"`
	Username             string          `json:"Username" validate:"required"`

	AccountResource
}

// NewSSHKeyAccount initializes and returns a SSH key pair account with a name,
// username, and private key file.
func NewSSHKeyAccount(name string, username string, privateKeyFile SensitiveValue) *SSHKeyAccount {
	return &SSHKeyAccount{
		AccountType:     "SshKeyPair",
		Username:        username,
		PrivateKeyFile:  &privateKeyFile,
		AccountResource: *newAccountResource(name),
	}
}

// GetAccountType returns the account type for this account.
func (s *SSHKeyAccount) GetAccountType() string {
	return s.AccountType
}

// Validate checks the state of this account and returns an error if invalid.
func (s *SSHKeyAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(s)
}

var _ IAccount = &SSHKeyAccount{}
