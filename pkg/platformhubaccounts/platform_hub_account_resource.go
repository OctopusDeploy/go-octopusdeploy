package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// PlatformHubAccountResource represents details for all Platform Hub accounts.
type PlatformHubAccountResource struct {
	AccountType PlatformHubAccountType `json:"AccountType" validate:"required,oneof=AmazonWebServicesAccount AmazonWebServicesOidcAccount AzureOidc AzureServicePrincipal GoogleCloudAccount GenericOidcAccount UsernamePassword"`
	Description string                 `json:"Description,omitempty"`
	Name        string                 `json:"Name" validate:"required,notall"`
	AccessKey                     string                 `json:"AccessKey,omitempty"`
	SecretKey                     *core.SensitiveValue   `json:"SecretKey,omitempty"`
	JsonKey                       *core.SensitiveValue   `json:"JsonKey,omitempty"`
	ExecutionSubjectKeys          []string               `json:"ExecutionSubjectKeys,omitempty"`
	Audience                      string                 `json:"Audience,omitempty"`
	Username                      string                 `json:"Username,omitempty"`
	Password                      *core.SensitiveValue   `json:"Password,omitempty"`
	SubscriptionNumber            string                 `json:"SubscriptionNumber,omitempty"`
	TenantID                      string                 `json:"TenantId,omitempty"`
	ApplicationID                 string                 `json:"ClientId,omitempty"`
	DeploymentSubjectKeys         []string               `json:"DeploymentSubjectKeys,omitempty"`
	HealthCheckSubjectKeys        []string               `json:"HealthCheckSubjectKeys,omitempty"`
	AccountTestSubjectKeys        []string               `json:"AccountTestSubjectKeys,omitempty"`
	AzureEnvironment              string                 `json:"AzureEnvironment,omitempty"`
	AuthenticationEndpoint        string                 `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ResourceManagementEndpoint    string                 `json:"ResourceManagementEndpointBaseUri,omitempty"`

	RoleArn                string `json:"RoleArn,omitempty"`
	SessionDuration        string `json:"SessionDuration,omitempty"`

	resources.Resource
}

// NewPlatformHubAccountResource creates and initializes a Platform Hub account resource.
func NewPlatformHubAccountResource(name string, accountType PlatformHubAccountType) *PlatformHubAccountResource {
	return &PlatformHubAccountResource{
		AccountType: accountType,
		Name:        name,
		Resource:    *resources.NewResource(),
	}
}

// GetAccountType returns the type of this account resource.
func (r *PlatformHubAccountResource) GetAccountType() PlatformHubAccountType {
	return r.AccountType
}

// GetDescription returns the description of this account resource.
func (r *PlatformHubAccountResource) GetDescription() string {
	return r.Description
}

// GetName returns the name of this account resource.
func (r *PlatformHubAccountResource) GetName() string {
	return r.Name
}

// SetDescription sets the description of the account resource.
func (r *PlatformHubAccountResource) SetDescription(description string) {
	r.Description = description
}

// SetName sets the name of this account resource.
func (r *PlatformHubAccountResource) SetName(name string) {
	r.Name = name
}

// ToPlatformHubAccount converts a PlatformHubAccountResource to the appropriate concrete account type.
func (r *PlatformHubAccountResource) ToPlatformHubAccount() (IPlatformHubAccount, error) {
	if r == nil {
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccount", "PlatformHubAccountResource")
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	var account IPlatformHubAccount

	switch r.AccountType {
	case AccountTypePlatformHubAwsAccount:
		awsAccount, err := NewPlatformHubAwsAccount(r.GetName(), r.AccessKey, r.SecretKey)
		if err != nil {
			return nil, err
		}
		account = awsAccount
	case AccountTypePlatformHubAwsOIDCAccount:
		oidcAccount, err := NewPlatformHubAwsOIDCAccount(r.GetName(), r.RoleArn)
		if err != nil {
			return nil, err
		}
		oidcAccount.SessionDuration = r.SessionDuration
		oidcAccount.DeploymentSubjectKeys = r.DeploymentSubjectKeys
		oidcAccount.HealthCheckSubjectKeys = r.HealthCheckSubjectKeys
		oidcAccount.AccountTestSubjectKeys = r.AccountTestSubjectKeys
		account = oidcAccount
	case AccountTypePlatformHubAzureServicePrincipalAccount:
		azureSpAccount, err := NewPlatformHubAzureServicePrincipalAccount(r.GetName(), r.SubscriptionNumber, r.TenantID, r.ApplicationID, r.Password)
		if err != nil {
			return nil, err
		}
		azureSpAccount.AzureEnvironment = r.AzureEnvironment
		azureSpAccount.AuthenticationEndpoint = r.AuthenticationEndpoint
		azureSpAccount.ResourceManagementEndpoint = r.ResourceManagementEndpoint
		account = azureSpAccount
	case AccountTypePlatformHubAzureOidcAccount:
		azureOidcAccount, err := NewPlatformHubAzureOidcAccount(r.GetName(), r.SubscriptionNumber, r.ApplicationID, r.TenantID)
		if err != nil {
			return nil, err
		}
		azureOidcAccount.ExecutionSubjectKeys = r.DeploymentSubjectKeys
		azureOidcAccount.HealthSubjectKeys = r.HealthCheckSubjectKeys
		azureOidcAccount.AccountTestSubjectKeys = r.AccountTestSubjectKeys
		azureOidcAccount.Audience = r.Audience
		azureOidcAccount.AzureEnvironment = r.AzureEnvironment
		azureOidcAccount.AuthenticationEndpoint = r.AuthenticationEndpoint
		azureOidcAccount.ResourceManagementEndpoint = r.ResourceManagementEndpoint
		account = azureOidcAccount
	case AccountTypePlatformHubGcpAccount:
		gcpAccount, err := NewPlatformHubGcpAccount(r.GetName(), r.JsonKey)
		if err != nil {
			return nil, err
		}
		account = gcpAccount
	case AccountTypePlatformHubGenericOidcAccount:
		oidcAccount, err := NewPlatformHubGenericOidcAccount(r.GetName())
		if err != nil {
			return nil, err
		}
		oidcAccount.ExecutionSubjectKeys = r.ExecutionSubjectKeys
		oidcAccount.Audience = r.Audience
		account = oidcAccount
	case AccountTypePlatformHubUsernamePasswordAccount:
		usernamePasswordAccount, err := NewPlatformHubUsernamePasswordAccount(r.GetName(), r.Username, r.Password)
		if err != nil {
			return nil, err
		}
		account = usernamePasswordAccount
	default:
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccount", "AccountType")
	}

	account.SetDescription(r.Description)
	account.SetID(r.GetID())
	account.SetLinks(r.GetLinks())
	account.SetModifiedBy(r.GetModifiedBy())
	account.SetModifiedOn(r.GetModifiedOn())

	return account, nil
}

// ToPlatformHubAccountResource converts a concrete account type to PlatformHubAccountResource.
func ToPlatformHubAccountResource(account IPlatformHubAccount) (*PlatformHubAccountResource, error) {
	if account == nil {
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccountResource", "PlatformHubAccount")
	}

	// conversion unnecessary if input account is *PlatformHubAccountResource
	if v, ok := account.(*PlatformHubAccountResource); ok {
		return v, nil
	}

	resource := NewPlatformHubAccountResource(account.GetName(), account.GetAccountType())

	switch resource.AccountType {
	case AccountTypePlatformHubAwsAccount:
		awsAccount := account.(*PlatformHubAwsAccount)
		resource.AccessKey = awsAccount.AccessKey
		resource.SecretKey = awsAccount.SecretKey
	case AccountTypePlatformHubAwsOIDCAccount:
		oidcAccount := account.(*PlatformHubAwsOIDCAccount)
		resource.RoleArn = oidcAccount.RoleArn
		resource.SessionDuration = oidcAccount.SessionDuration
		resource.DeploymentSubjectKeys = oidcAccount.DeploymentSubjectKeys
		resource.HealthCheckSubjectKeys = oidcAccount.HealthCheckSubjectKeys
		resource.AccountTestSubjectKeys = oidcAccount.AccountTestSubjectKeys
	case AccountTypePlatformHubAzureServicePrincipalAccount:
		azureSpAccount := account.(*PlatformHubAzureServicePrincipalAccount)
		resource.SubscriptionNumber = azureSpAccount.SubscriptionID
		resource.TenantID = azureSpAccount.TenantID
		resource.ApplicationID = azureSpAccount.ApplicationID
		resource.Password = azureSpAccount.Password
		resource.AzureEnvironment = azureSpAccount.AzureEnvironment
		resource.AuthenticationEndpoint = azureSpAccount.AuthenticationEndpoint
		resource.ResourceManagementEndpoint = azureSpAccount.ResourceManagementEndpoint
	case AccountTypePlatformHubAzureOidcAccount:
		azureOidcAccount := account.(*PlatformHubAzureOidcAccount)
		resource.SubscriptionNumber = azureOidcAccount.SubscriptionID
		resource.ApplicationID = azureOidcAccount.ApplicationID
		resource.TenantID = azureOidcAccount.TenantID
		resource.DeploymentSubjectKeys = azureOidcAccount.ExecutionSubjectKeys
		resource.HealthCheckSubjectKeys = azureOidcAccount.HealthSubjectKeys
		resource.AccountTestSubjectKeys = azureOidcAccount.AccountTestSubjectKeys
		resource.Audience = azureOidcAccount.Audience
		resource.AzureEnvironment = azureOidcAccount.AzureEnvironment
		resource.AuthenticationEndpoint = azureOidcAccount.AuthenticationEndpoint
		resource.ResourceManagementEndpoint = azureOidcAccount.ResourceManagementEndpoint
	case AccountTypePlatformHubGcpAccount:
		gcpAccount := account.(*PlatformHubGcpAccount)
		resource.JsonKey = gcpAccount.JsonKey
	case AccountTypePlatformHubGenericOidcAccount:
		oidcAccount := account.(*PlatformHubGenericOidcAccount)
		resource.ExecutionSubjectKeys = oidcAccount.ExecutionSubjectKeys
		resource.Audience = oidcAccount.Audience
	case AccountTypePlatformHubUsernamePasswordAccount:
		usernamePasswordAccount := account.(*PlatformHubUsernamePasswordAccount)
		resource.Username = usernamePasswordAccount.Username
		resource.Password = usernamePasswordAccount.Password
	default:
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccountResource", "AccountType")
	}

	resource.SetDescription(account.GetDescription())
	resource.SetID(account.GetID())
	resource.SetLinks(account.GetLinks())
	resource.SetModifiedBy(account.GetModifiedBy())
	resource.SetModifiedOn(account.GetModifiedOn())

	return resource, nil
}

var _ IPlatformHubAccount = &PlatformHubAccountResource{}
