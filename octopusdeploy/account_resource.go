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
	AccountType             AccountType            `json:"AccountType"`
	ApplicationID           *uuid.UUID             `json:"ClientId,omitempty"`
	ApplicationPassword     *SensitiveValue        `json:"Password,omitempty"`
	AuthenticationEndpoint  string                 `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	AzureEnvironment        string                 `json:"AzureEnvironment,omitempty"`
	CertificateBytes        *SensitiveValue        `json:"CertificateBytes,omitempty"`
	CertificateThumbprint   string                 `json:"CertificateThumbprint,omitempty"`
	Description             string                 `json:"Description,omitempty"`
	EnvironmentIDs          []string               `json:"EnvironmentIds,omitempty"`
	ManagementEndpoint      string                 `json:"ServiceManagementEndpointBaseUri,omitempty"`
	Name                    string                 `json:"Name" validate:"required,notall"`
	PrivateKeyFile          *SensitiveValue        `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase    *SensitiveValue        `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagerEndpoint string                 `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey               *SensitiveValue        `json:"SecretKey,omitempty"`
	StorageEndpointSuffix   string                 `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                 string                 `json:"SpaceId,omitempty"`
	SubscriptionID          *uuid.UUID             `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentMode  TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
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
		TenantedDeploymentMode: "Untenanted",
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

func (a *AccountResource) GetTenantedDeploymentMode() TenantedDeploymentMode {
	return a.TenantedDeploymentMode
}

func (a *AccountResource) GetTenantIDs() []string {
	return a.TenantIDs
}

func (a *AccountResource) GetTenantTags() []string {
	return a.TenantTags
}

// SetDescription sets the description of the account resource.
func (a *AccountResource) SetDescription(description string) {
	a.Description = description
}

// SetName sets the name of this account resource.
func (a *AccountResource) SetName(name string) {
	a.Name = name
}

// SetSpaceID sets the space ID of this account resource.
func (a *AccountResource) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
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
