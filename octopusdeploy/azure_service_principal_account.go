package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureServicePrincipalAccount represents an Azure service principal account.
type AzureServicePrincipalAccount struct {
	ApplicationID           *uuid.UUID      `json:"ClientId" validate:"required"`
	ApplicationPassword     *SensitiveValue `json:"Password" validate:"required"`
	AuthenticationEndpoint  string          `json:"ActiveDirectoryEndpointBaseUri,omitempty" validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string          `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string          `json:"ResourceManagementEndpointBaseUri" validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID      `json:"SubscriptionNumber" validate:"required"`
	TenantID                *uuid.UUID      `json:"TenantId" validate:"required"`

	account
}

// NewAzureServicePrincipalAccount creates and initializes an Azure service principal account.
func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, applicationPassword *SensitiveValue) (*AzureServicePrincipalAccount, error) {
	if isEmpty(name) {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterName)
	}

	if applicationPassword == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterApplicationPassword)
	}

	account := AzureServicePrincipalAccount{
		ApplicationID:       &applicationID,
		ApplicationPassword: applicationPassword,
		SubscriptionID:      &subscriptionID,
		TenantID:            &tenantID,
		account:             *newAccount(name, AccountType("AzureServicePrincipal")),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this Azure service principal account and
// returns an error if invalid.
func (a *AzureServicePrincipalAccount) Validate() error {
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
