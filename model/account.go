package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

// Accounts defines a collection of accounts with built-in support for paged
// results.
type Accounts struct {
	Items []Account `json:"Items"`
	PagedResults
}

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
	ResourceManagementEndpointBase   string                      `json:"ResourceManagementEndpointBaseUri,omitempty"`

	Resource
}

// NewAccount initializes an account with a name and account type.
func NewAccount(name string, accountType enum.AccountType) (*Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAccount", "name")
	}

	return &Account{
		Name:        name,
		AccountType: accountType,
	}, nil
}

func (a *Account) GetID() string {
	return a.ID
}

func (a *Account) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	switch a.AccountType {
	case enum.AzureServicePrincipal:
		return validateAzureServicePrincipalAccount(a)
	case enum.AzureSubscription:
		return validateAzureSubscriptionAccount(a)
	case enum.SshKeyPair:
		return validateSSHKeyAccount(a)
	case enum.Token:
		return validateTokenAccount(a)
	case enum.UsernamePassword:
		return validateUsernamePasswordAccount(a)
	}

	return nil
}

var _ ResourceInterface = &Account{}
