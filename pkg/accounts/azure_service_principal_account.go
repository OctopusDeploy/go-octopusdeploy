package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureServicePrincipalAccount represents an Azure service principal account.
type AzureServicePrincipalAccount struct {
	ApplicationID           *uuid.UUID           `json:"ClientId" validate:"required"`
	ApplicationPassword     *core.SensitiveValue `json:"Password" validate:"required"`
	AuthenticationEndpoint  string               `json:"ActiveDirectoryEndpointBaseUri,omitempty" validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string               `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string               `json:"ResourceManagementEndpointBaseUri" validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID           `json:"SubscriptionNumber" validate:"required"`
	TenantID                *uuid.UUID           `json:"TenantId" validate:"required"`

	account
}

// NewAzureServicePrincipalAccount creates and initializes an Azure service principal account.
func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, applicationPassword *core.SensitiveValue) (*AzureServicePrincipalAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	if applicationPassword == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterApplicationPassword)
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
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
