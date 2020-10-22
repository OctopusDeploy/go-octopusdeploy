package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// Account represents account details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs.
type Account struct {
	AccessKey               string          `json:"AccessKey,omitempty"`
	AccountType             string          `json:"AccountType" validate:"required,oneof=None UsernamePassword SshKeyPair AzureSubscription AzureServicePrincipal AmazonWebServicesAccount AmazonWebServicesRoleAccount Token"`
	ApplicationID           *uuid.UUID      `json:"ClientId,omitempty"`
	ApplicationPassword     *SensitiveValue `json:"Password,omitempty"`
	AuthenticationEndpoint  string          `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	AzureEnvironment        string          `json:"AzureEnvironment,omitempty"`
	CertificateBytes        *SensitiveValue `json:"CertificateBytes,omitempty"`
	CertificateThumbprint   string          `json:"CertificateThumbprint,omitempty"`
	Description             string          `json:"Description,omitempty"`
	EnvironmentIDs          []string        `json:"EnvironmentIds,omitempty"`
	Name                    string          `json:"Name" validate:"required,notall"`
	PrivateKeyFile          *SensitiveValue `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase    *SensitiveValue `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagerEndpoint string          `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey               *SensitiveValue `json:"SecretKey,omitempty"`
	ManagementEndpoint      string          `json:"ServiceManagementEndpointBaseUri,omitempty"`
	StorageEndpointSuffix   string          `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                 string          `json:"SpaceId,omitempty"`
	SubscriptionID          *uuid.UUID      `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentMode  string          `json:"TenantedDeploymentParticipation" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantID                *uuid.UUID      `json:"TenantId,omitempty"`
	TenantIDs               []string        `json:"TenantIds,omitempty"`
	TenantTags              []string        `json:"TenantTags,omitempty"`
	Token                   *SensitiveValue `json:"Token,omitempty"`
	Username                string          `json:"Username,omitempty"`

	resource
}

// NewAccount creates and initializes an account with a name and type.
func NewAccount(name string, accountType string) *Account {
	return &Account{
		AccountType:            accountType,
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
		resource:               *newResource(),
	}
}

// GetAccountType returns the type of this account.
func (a *Account) GetAccountType() string {
	return a.AccountType
}

// GetDescription returns the description of this account.
func (a *Account) GetDescription() string {
	return a.Description
}

// GetName returns the name of this account.
func (a *Account) GetName() string {
	return a.Name
}

// SetDescription sets the description of the account.
func (a *Account) SetDescription(description string) {
	a.Description = description
}

// SetName sets the name of this account.
func (a *Account) SetName(name string) {
	a.Name = name
}

// Validate checks the state of the account and returns an error if invalid.
func (a *Account) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &Account{}
