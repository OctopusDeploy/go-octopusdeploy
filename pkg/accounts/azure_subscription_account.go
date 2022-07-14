package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureSubscriptionAccount represents an Azure subscription account.
type AzureSubscriptionAccount struct {
	AzureEnvironment      string               `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	CertificateBytes      *core.SensitiveValue `json:"CertificateBytes,omitempty"`
	CertificateThumbprint string               `json:"CertificateThumbprint,omitempty"`
	ManagementEndpoint    string               `json:"ServiceManagementEndpointBaseUri,omitempty" validate:"omitempty,uri"`
	StorageEndpointSuffix string               `json:"ServiceManagementEndpointSuffix,omitempty" validate:"omitempty,hostname"`
	SubscriptionID        *uuid.UUID           `json:"SubscriptionNumber" validate:"required"`

	account
}

// NewAzureSubscriptionAccount creates and initializes an Azure subscription account with a name.
func NewAzureSubscriptionAccount(name string, subscriptionID uuid.UUID) (*AzureSubscriptionAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	account := AzureSubscriptionAccount{
		SubscriptionID: &subscriptionID,
		account:        *newAccount(name, AccountType("AzureSubscription")),
	}

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
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
