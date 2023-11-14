package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureOIDCAccount represents an Azure OIDC account.
type AzureOIDCAccount struct {
	ApplicationID           *uuid.UUID `json:"ClientId" validate:"required"`
	AuthenticationEndpoint  string     `json:"ActiveDirectoryEndpointBaseUri,omitempty" validate:"required_with=AzureEnvironment,omitempty,uri"`
	AzureEnvironment        string     `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	ResourceManagerEndpoint string     `json:"ResourceManagementEndpointBaseUri" validate:"required_with=AzureEnvironment,omitempty,uri"`
	SubscriptionID          *uuid.UUID `json:"SubscriptionNumber" validate:"required"`
	TenantID                *uuid.UUID `json:"TenantId" validate:"required"`
	Audience                string     `json:"Audience,omitempty"`
	DeploymentSubjectKeys   []string   `json:"DeploymentSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type'"`
	HealthCheckSubjectKeys  []string   `json:"HealthCheckSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account target type'"`
	AccountTestSubjectKeys  []string   `json:"AccountTestSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account type'"`

	account
}

// NewAzureOIDCAccount creates and initializes an Azure OIDC account.
func NewAzureOIDCAccount(name string, subscriptionID uuid.UUID, tenantID uuid.UUID, applicationID uuid.UUID) (*AzureOIDCAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}

	account := AzureOIDCAccount{
		ApplicationID:  &applicationID,
		SubscriptionID: &subscriptionID,
		TenantID:       &tenantID,
		account:        *newAccount(name, AccountType("AzureOIDC")),
	}

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AzureOIDCAccount) Validate() error {
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
