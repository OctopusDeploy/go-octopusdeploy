package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	resources "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// AccountResource represents account details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs.
type AccountResource struct {
	AccessKey               string                      `json:"AccessKey,omitempty"`
	AccountType             AccountType                 `json:"AccountType" validate:"required,oneof=None UsernamePassword SshKeyPair AzureSubscription AzureServicePrincipal AzureOidc AmazonWebServicesAccount AmazonWebServicesRoleAccount AmazonWebServicesOidcAccount GoogleCloudAccount GenericOidcAccount Token"`
	ApplicationID           *uuid.UUID                  `json:"ClientId,omitempty"`
	ApplicationPassword     *core.SensitiveValue        `json:"Password,omitempty"`
	AuthenticationEndpoint  string                      `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	AzureEnvironment        string                      `json:"AzureEnvironment,omitempty"`
	CertificateBytes        *core.SensitiveValue        `json:"CertificateBytes,omitempty"`
	CertificateThumbprint   string                      `json:"CertificateThumbprint,omitempty"`
	Description             string                      `json:"Description,omitempty"`
	EnvironmentIDs          []string                    `json:"EnvironmentIds,omitempty"`
	JsonKey                 *core.SensitiveValue        `json:"JsonKey,omitempty"`
	ManagementEndpoint      string                      `json:"ServiceManagementEndpointBaseUri,omitempty"`
	Name                    string                      `json:"Name" validate:"required,notall"`
	PrivateKeyFile          *core.SensitiveValue        `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase    *core.SensitiveValue        `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagerEndpoint string                      `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey               *core.SensitiveValue        `json:"SecretKey,omitempty"`
	Slug                    string                      `json:"Slug,omitempty"`
	StorageEndpointSuffix   string                      `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                 string                      `json:"SpaceId,omitempty"`
	SubscriptionID          *uuid.UUID                  `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentMode  core.TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantID                *uuid.UUID                  `json:"TenantId,omitempty"`
	TenantIDs               []string                    `json:"TenantIds,omitempty"`
	TenantTags              []string                    `json:"TenantTags,omitempty"`
	Token                   *core.SensitiveValue        `json:"Token,omitempty"`
	Username                string                      `json:"Username,omitempty"`
	Audience                string                      `json:"Audience,omitempty"`
	DeploymentSubjectKeys   []string                    `json:"DeploymentSubjectKeys,omitempty"`
	HealthCheckSubjectKeys  []string                    `json:"HealthCheckSubjectKeys,omitempty"`
	AccountTestSubjectKeys  []string                    `json:"AccountTestSubjectKeys,omitempty"`
	RoleArn                 string                      `json:"RoleArn,omitempty"`
	SessionDuration         string                      `json:"SessionDuration,omitempty"`
	CustomClaims            map[string]string           `json:"CustomClaims,omitempty"`

	resources.Resource
}

// NewAccount creates and initializes an account resource with a name and type.
func NewAccountResource(name string, accountType AccountType) *AccountResource {
	return &AccountResource{
		AccountType:            accountType,
		Name:                   name,
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
		Resource:               *resources.NewResource(),
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
func (a *AccountResource) GetTenantedDeploymentMode() core.TenantedDeploymentMode {
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

// GetSlug returns the slug to this account.
func (a *AccountResource) GetSlug() string {
	return a.Slug
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
func (a *AccountResource) SetTenantedDeploymentMode(mode core.TenantedDeploymentMode) {
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

// SetSlug sets the slug of this account.
func (a *AccountResource) SetSlug(slug string) {
	a.Slug = slug
}

// Validate checks the state of the account resource and returns an error if
// invalid.
func (a *AccountResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &AccountResource{}
