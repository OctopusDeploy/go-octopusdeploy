package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type PlatformHubAzureServicePrincipalAccount struct {
	SubscriptionID              string               `json:"SubscriptionNumber" validate:"required"`
	TenantID                    string               `json:"TenantId" validate:"required"`
	ApplicationID               string               `json:"ClientId" validate:"required"`
	Password                    *core.SensitiveValue `json:"Password" validate:"required"`
	AzureEnvironment            string               `json:"AzureEnvironment,omitempty"`
	AuthenticationEndpoint      string               `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ResourceManagementEndpoint  string               `json:"ResourceManagementEndpointBaseUri,omitempty"`

	platformHubAccount
}

func NewPlatformHubAzureServicePrincipalAccount(name string, subscriptionID string, tenantID string, applicationID string, password *core.SensitiveValue) (*PlatformHubAzureServicePrincipalAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(subscriptionID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("subscriptionID")
	}

	if internal.IsEmpty(tenantID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("tenantID")
	}

	if internal.IsEmpty(applicationID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("applicationID")
	}

	if password == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterApplicationPassword)
	}

	account := PlatformHubAzureServicePrincipalAccount{
		SubscriptionID:     subscriptionID,
		TenantID:           tenantID,
		ApplicationID:      applicationID,
		Password:           password,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubAzureServicePrincipalAccount),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *PlatformHubAzureServicePrincipalAccount) Validate() error {
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
