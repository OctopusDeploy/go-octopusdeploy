package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// SSHKeyAccount represents a SSH key pair account.
type SSHKeyAccount struct {
	PrivateKeyFile       *SensitiveValue `validate:"required"`
	PrivateKeyPassphrase *SensitiveValue
	Username             string `validate:"required"`

	account
}

// NewSSHKeyAccount initializes and returns a SSH key pair account with a name,
// username, and private key file.
func NewSSHKeyAccount(name string, username string, privateKeyFile *SensitiveValue, options ...func(*SSHKeyAccount)) (*SSHKeyAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if isEmpty(username) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUsername)
	}

	if privateKeyFile == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterPrivateKeyFile)
	}

	account := SSHKeyAccount{
		account: *newAccount(name, AccountType("SshKeyPair")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.accountType = AccountType("SshKeyPair")
	account.ID = emptyString
	account.ModifiedBy = emptyString
	account.ModifiedOn = nil
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

// Validate checks the state of this account and returns an error if invalid.
func (s *SSHKeyAccount) Validate() error {
	v := validator.New()
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
