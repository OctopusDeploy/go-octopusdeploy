package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureServicePrincipalAccount represents an Azure service principal account.
type AzureServicePrincipalAccount struct {
	ApplicationID           *uuid.UUID                `validate:"required"`
	ApplicationPassword     *resources.SensitiveValue `validate:"required"`
	AuthenticationEndpoint  string                    `validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string                    `validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string                    `validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID                `validate:"required"`
	TenantID                *uuid.UUID                `validate:"required"`

	Account
}

// NewAzureServicePrincipalAccount creates and initializes an Azure service
// principal account.
func NewAzureServicePrincipalAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID, applicationPassword *resources.SensitiveValue, options ...func(*AzureServicePrincipalAccount)) (*AzureServicePrincipalAccount, error) {
	if octopusdeploy.IsEmpty(name) {
		return nil, octopusdeploy.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
	}

	if applicationPassword == nil {
		return nil, octopusdeploy.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterApplicationPassword)
	}

	account := AzureServicePrincipalAccount{
		Account: *NewAccount(name, AccountType("AzureServicePrincipal")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("AzureServicePrincipal")
	account.ApplicationID = &applicationID
	account.ApplicationPassword = applicationPassword
	account.Name = name
	account.SubscriptionID = &subscriptionID
	account.TenantID = &tenantID

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
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
