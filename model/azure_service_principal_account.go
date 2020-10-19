package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureServicePrincipalAccount represents an Azure service principal account.
type AzureServicePrincipalAccount struct {
	AccountType             string          `json:"AccountType" validate:"required,eq=AzureServicePrincipal"`
	ApplicationID           *uuid.UUID      `json:"ClientId" validate:"required"`
	ApplicationPassword     *SensitiveValue `json:"Password" validate:"required"`
	AuthenticationEndpoint  string          `json:"ActiveDirectoryEndpointBaseUri,omitempty" validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string          `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string          `json:"ResourceManagementEndpointBaseUri,omitempty" validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID      `json:"SubscriptionNumber" validate:"required"`
	TenantID                *uuid.UUID      `json:"TenantId" validate:"required"`

	AccountResource
}

// NewAzureServicePrincipalAccount creates and initializes an Azure service
// principal account.
func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, applicationPassword SensitiveValue) *AzureServicePrincipalAccount {
	return &AzureServicePrincipalAccount{
		AccountType:         "AzureServicePrincipal",
		ApplicationID:       &applicationID,
		ApplicationPassword: &applicationPassword,
		SubscriptionID:      &subscriptionID,
		TenantID:            &tenantID,
		AccountResource:     *newAccountResource(name),
	}
}

// GetAccountType returns the account type of this Azure service principal
// account.
func (a *AzureServicePrincipalAccount) GetAccountType() string {
	return a.AccountType
}

// Validate checks the state of this Azure service principal account and
// returns an error if invalid.
func (a *AzureServicePrincipalAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &AzureServicePrincipalAccount{}
