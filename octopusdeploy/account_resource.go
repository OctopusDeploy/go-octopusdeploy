package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// AccountResources defines a collection of account resources with built-in
// support for paged results.
type AccountResources struct {
	Items []*AccountResource `json:"Items"`
	PagedResults
}

// AccountResource represents account details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs.
type AccountResource struct {
	AccessKey               string                 `json:"AccessKey,omitempty"`
	AccountType             AccountType            `json:"AccountType" validate:"required,oneof=None UsernamePassword SshKeyPair AzureSubscription AzureServicePrincipal AmazonWebServicesAccount AmazonWebServicesRoleAccount GoogleCloudAccount Token"`
	ApplicationID           *uuid.UUID             `json:"ClientId,omitempty"`
	ApplicationPassword     *SensitiveValue        `json:"Password,omitempty"`
	AuthenticationEndpoint  string                 `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	AzureEnvironment        string                 `json:"AzureEnvironment,omitempty"`
	CertificateBytes        *SensitiveValue        `json:"CertificateBytes,omitempty"`
	CertificateThumbprint   string                 `json:"CertificateThumbprint,omitempty"`
	Description             string                 `json:"Description,omitempty"`
	EnvironmentIDs          []string               `json:"EnvironmentIds,omitempty"`
	JsonKey                 *SensitiveValue        `json:"JsonKey,omitempty"`
	ManagementEndpoint      string                 `json:"ServiceManagementEndpointBaseUri,omitempty"`
	Name                    string                 `json:"Name" validate:"required,notall"`
	PrivateKeyFile          *SensitiveValue        `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase    *SensitiveValue        `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagerEndpoint string                 `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey               *SensitiveValue        `json:"SecretKey,omitempty"`
	StorageEndpointSuffix   string                 `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                 string                 `json:"SpaceId,omitempty"`
	SubscriptionID          *uuid.UUID             `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentMode  TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantID                *uuid.UUID             `json:"TenantId,omitempty"`
	TenantIDs               []string               `json:"TenantIds,omitempty"`
	TenantTags              []string               `json:"TenantTags,omitempty"`
	Token                   *SensitiveValue        `json:"Token,omitempty"`
	Username                string                 `json:"Username,omitempty"`

	resource
}

// NewAccount creates and initializes an account resource with a name and type.
func NewAccountResource(name string, accountType AccountType) *AccountResource {
	return &AccountResource{
		AccountType:            accountType,
		Name:                   name,
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
		resource:               *newResource(),
	}
}

// GetAccountType returns the type of this account resource.
func (a *AccountResource) GetAccountType() AccountType {
	return a.AccountType
}

// GetDescription returns the description of this account resource.
func (a *AccountResource) GetDescription() string {
	return a.Description
}

func (a *AccountResource) GetEnvironmentIDs() []string {
	return a.EnvironmentIDs
}

// GetName returns the name of this account resource.
func (a *AccountResource) GetName() string {
	return a.Name
}

// GetSpaceID returns the space ID of this account resource.
func (a *AccountResource) GetSpaceID() string {
	return a.SpaceID
}

// GetTenantedDeploymentMode returns the tenanted deployment mode of this account resource.
func (a *AccountResource) GetTenantedDeploymentMode() TenantedDeploymentMode {
	return a.TenantedDeploymentMode
}

// GetTenantIDs returns the tenant IDs associated with this account resource.
func (a *AccountResource) GetTenantIDs() []string {
	return a.TenantIDs
}

// GetTenantTags returns the tenant tags assigned to this account resource.
func (a *AccountResource) GetTenantTags() []string {
	return a.TenantTags
}

// SetDescription sets the description of the account resource.
func (a *AccountResource) SetDescription(description string) {
	a.Description = description
}

// SetEnvironmentIDs sets the associated environment IDs of the account resource.
func (a *AccountResource) SetEnvironmentIDs(environmentIds []string) {
	a.EnvironmentIDs = environmentIds
}

// SetName sets the name of this account resource.
func (a *AccountResource) SetName(name string) {
	a.Name = name
}

// SetSpaceID sets the space ID of this account resource.
func (a *AccountResource) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
}

// SetTenantedDeploymentMode sets the tenanted deployment mode of this account resource.
func (a *AccountResource) SetTenantedDeploymentMode(mode TenantedDeploymentMode) {
	a.TenantedDeploymentMode = mode
}

// SetTenantIDs sets the tenant IDs associated with this account resource.
func (a *AccountResource) SetTenantIDs(tenantIds []string) {
	a.TenantIDs = tenantIds
}

// SetTenantTags sets the tenant tags associated with this account resource.
func (a *AccountResource) SetTenantTags(tenantTags []string) {
	a.TenantTags = tenantTags
}

// Validate checks the state of the account resource and returns an error if
// invalid.
func (a *AccountResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &AccountResource{}
