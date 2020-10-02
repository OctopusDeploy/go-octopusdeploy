package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// Account represents account details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs.
type Account struct {
	AccessKey                        string                      `json:"AccessKey,omitempty"`
	AccountType                      enum.AccountType            `json:"AccountType" validate:"required"`
	ActiveDirectoryEndpointBase      string                      `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ApplicationID                    *uuid.UUID                  `json:"ClientId,omitempty"`
	AzureEnvironment                 string                      `json:"AzureEnvironment,omitempty"`
	CertificateBytes                 *SensitiveValue             `json:"CertificateBytes,omitempty"`
	CertificateThumbprint            string                      `json:"CertificateThumbprint,omitempty"`
	Description                      string                      `json:"Description,omitempty"`
	EnvironmentIDs                   []string                    `json:"EnvironmentIds,omitempty"`
	Name                             string                      `json:"Name" validate:"required"`
	Password                         *SensitiveValue             `json:"Password,omitempty"`
	PrivateKeyFile                   *SensitiveValue             `json:"PrivateKeyFile,omitempty"`
	PrivateKeyPassphrase             *SensitiveValue             `json:"PrivateKeyPassphrase,omitempty"`
	ResourceManagementEndpointBase   string                      `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SecretKey                        *SensitiveValue             `json:"SecretKey,omitempty"`
	ServiceManagementEndpointBaseURI string                      `json:"ServiceManagementEndpointBaseUri,omitempty"`
	ServiceManagementEndpointSuffix  string                      `json:"ServiceManagementEndpointSuffix,omitempty"`
	SpaceID                          string                      `json:"SpaceId,omitempty"`
	SubscriptionID                   *uuid.UUID                  `json:"SubscriptionNumber,omitempty"`
	TenantedDeploymentParticipation  enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantID                         *uuid.UUID                  `json:"TenantId,omitempty"`
	TenantIDs                        []string                    `json:"TenantIds,omitempty"`
	TenantTags                       []string                    `json:"TenantTags,omitempty"`
	Token                            *SensitiveValue             `json:"Token,omitempty"`
	Username                         string                      `json:"Username,omitempty"`

	Resource
}

// Accounts defines a collection of accounts with built-in support for paged
// results.
type Accounts struct {
	Items []Account `json:"Items"`
	PagedResults
}

// NewAccount initializes an account with a name and type. If any of the input
// parameters are invalid, it will return nil and an error.
func NewAccount(name string, accountType enum.AccountType) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAccount", "name")
	}

	return &Account{
		Name:        name,
		AccountType: accountType,
	}, nil
}

// GetID returns the ID value of the Account.
func (resource Account) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Account.
func (resource Account) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Account was changed.
func (resource Account) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Account.
func (resource Account) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource Account) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource Account) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource Account) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Account and returns an error if invalid.
func (resource Account) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	switch resource.AccountType {
	case enum.AzureServicePrincipal:
		return validateAzureServicePrincipalAccount(resource)
	case enum.AzureSubscription:
		return validateAzureSubscriptionAccount(resource)
	case enum.SshKeyPair:
		return validateSSHKeyAccount(resource)
	case enum.Token:
		return validateTokenAccount(resource)
	case enum.UsernamePassword:
		return validateUsernamePasswordAccount(resource)
	}

	return nil
}

var _ ResourceInterface = &Account{}
