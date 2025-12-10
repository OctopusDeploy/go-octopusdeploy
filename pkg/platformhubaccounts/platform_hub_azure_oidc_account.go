package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type PlatformHubAzureOidcAccount struct {
	SubscriptionID             string   `json:"SubscriptionNumber" validate:"required"`
	ApplicationID              string   `json:"ClientId" validate:"required"`
	TenantID                   string   `json:"TenantId" validate:"required"`
	ExecutionSubjectKeys       []string `json:"DeploymentSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space environment project tenant runbook account type"`
	HealthSubjectKeys          []string `json:"HealthCheckSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account target type"`
	AccountTestSubjectKeys     []string `json:"AccountTestSubjectKeys,omitempty" validate:"omitempty,dive,oneof=space account type"`
	Audience                   string   `json:"Audience,omitempty"`
	AzureEnvironment           string   `json:"AzureEnvironment,omitempty"`
	AuthenticationEndpoint     string   `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ResourceManagementEndpoint string   `json:"ResourceManagementEndpointBaseUri,omitempty"`

	platformHubAccount
}

func NewPlatformHubAzureOidcAccount(name string, subscriptionID string, applicationID string, tenantID string) (*PlatformHubAzureOidcAccount, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterName)
	}

	if internal.IsEmpty(subscriptionID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("subscriptionID")
	}

	if internal.IsEmpty(applicationID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("applicationID")
	}

	if internal.IsEmpty(tenantID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("tenantID")
	}

	account := PlatformHubAzureOidcAccount{
		SubscriptionID:     subscriptionID,
		ApplicationID:      applicationID,
		TenantID:           tenantID,
		platformHubAccount: *newPlatformHubAccount(name, AccountTypePlatformHubAzureOidcAccount),
	}

	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *PlatformHubAzureOidcAccount) Validate() error {
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
