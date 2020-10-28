package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureSubscriptionAccount represents an Azure subscription account.
type AzureSubscriptionAccount struct {
	AzureEnvironment      string `validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	CertificateBytes      *SensitiveValue
	CertificateThumbprint string
	ManagementEndpoint    string     `validate:"omitempty,uri"`
	StorageEndpointSuffix string     `validate:"omitempty,hostname"`
	SubscriptionID        *uuid.UUID `validate:"required"`

	account
}

// NewAzureSubscriptionAccount creates and initializes an Azure subscription
// account with a name.
func NewAzureSubscriptionAccount(name string, subscriptionID uuid.UUID, options ...func(*AzureSubscriptionAccount)) (*AzureSubscriptionAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	account := AzureSubscriptionAccount{
		account: *newAccount(name, AccountType("AzureSubscription")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("AzureSubscription")
	account.ID = emptyString
	account.ModifiedBy = emptyString
	account.ModifiedOn = nil
	account.Name = name
	account.SubscriptionID = &subscriptionID

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AzureSubscriptionAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
