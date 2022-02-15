package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// AccountResource represents accountV1 details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs.
type AccountResource struct {
	AccessKey               string                           `json:"AccessKey,omitempty"`
	AccountType             AccountType                      `json:"AccountType"`
	ApplicationID           *uuid.UUID                       `json:"ClientId,omitempty"`
	ApplicationPassword     *resources.SensitiveValue        `json:"Password,omitempty"`
	AuthenticationEndpoint  string                           `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	AzureEnvironment        string                           `json:"AzureEnvironment,omitempty"`
	CertificateBytes        *resources.SensitiveValue        `json:"CertificateBytes,omitempty"`
	CertificateThumbprint   string                           `json:"CertificateThumbprint,omitempty"`
	Description             string                           `json:"Description,omitempty"`
	EnvironmentIDs          []string                         `json:"EnvironmentIds,omitempty"`
	JsonKey                 *resources.SensitiveValue        `json:"JsonKey,omitempty"`
	ManagementEndpoint      string                           `json:"ServiceManagementEndpointBaseUri,omitempty"`
	Name                    string                           `json:"Name" validate:"required,notall"`
	PrivateKeyFile          *resources.SensitiveValue        `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase    *resources.SensitiveValue        `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagerEndpoint string                           `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey               *resources.SensitiveValue        `json:"SecretKey,omitempty"`
	StorageEndpointSuffix   string                           `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                 string                           `json:"SpaceId"`
	SubscriptionID          *uuid.UUID                       `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentMode  resources.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantID                *uuid.UUID                       `json:"TenantId,omitempty"`
	TenantIDs               []string                         `json:"TenantIds,omitempty"`
	TenantTags              []string                         `json:"TenantTags,omitempty"`
	Token                   *resources.SensitiveValue        `json:"Token,omitempty"`
	Username                string                           `json:"Username,omitempty"`

	resources.Resource
}

type AccountsQuery struct {
	AccountType AccountType `url:"accountType,omitempty"`
	service.IdsQuery
	service.PartialNameQuery
}

// NewAccount creates and initializes an accountV1 resource with a name and type.
func NewAccountResource(spaceID string, name string, accountType AccountType) IAccount {
	return &AccountResource{
		SpaceID:                spaceID,
		AccountType:            accountType,
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
		Resource:               *resources.NewResource(),
	}
}

// GetAccountType returns the type of this accountV1 resource.
func (a *AccountResource) GetAccountType() AccountType {
	return a.AccountType
}

// GetDescription returns the description of this accountV1 resource.
func (a *AccountResource) GetDescription() string {
	return a.Description
}

func (a *AccountResource) GetEnvironmentIDs() []string {
	return a.EnvironmentIDs
}

// GetName returns the name of this accountV1 resource.
func (a *AccountResource) GetName() string {
	return a.Name
}

// GetSpaceID returns the space ID of this accountV1 resource.
func (a *AccountResource) GetSpaceID() string {
	return a.SpaceID
}

// GetTenantedDeploymentMode returns the tenanted deployment mode of this accountV1 resource.
func (a *AccountResource) GetTenantedDeploymentMode() resources.TenantedDeploymentMode {
	return a.TenantedDeploymentMode
}

// GetTenantIDs returns the tenant IDs associated with this accountV1 resource.
func (a *AccountResource) GetTenantIDs() []string {
	return a.TenantIDs
}

// GetTenantTags returns the tenant tags assigned to this accountV1 resource.
func (a *AccountResource) GetTenantTags() []string {
	return a.TenantTags
}

// SetDescription sets the description of the accountV1 resource.
func (a *AccountResource) SetDescription(description string) {
	a.Description = description
}

// SetEnvironmentIDs sets the associated environment IDs of the accountV1 resource.
func (a *AccountResource) SetEnvironmentIDs(environmentIds []string) {
	a.EnvironmentIDs = environmentIds
}

// SetName sets the name of this accountV1 resource.
func (a *AccountResource) SetName(name string) {
	a.Name = name
}

// SetSpaceID sets the space ID of this accountV1 resource.
func (a *AccountResource) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
}

// SetTenantedDeploymentMode sets the tenanted deployment mode of this accountV1 resource.
func (a *AccountResource) SetTenantedDeploymentMode(mode resources.TenantedDeploymentMode) {
	a.TenantedDeploymentMode = mode
}

// SetTenantIDs sets the tenant IDs associated with this accountV1 resource.
func (a *AccountResource) SetTenantIDs(tenantIds []string) {
	a.TenantIDs = tenantIds
}

// SetTenantTags sets the tenant tags associated with this accountV1 resource.
func (a *AccountResource) SetTenantTags(tenantTags []string) {
	a.TenantTags = tenantTags
}

// Validate checks the state of the accountV1 resource and returns an error if
// invalid.
func (a *AccountResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
