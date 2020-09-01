package model

import (
	"errors"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
)

// Accounts defines a collection of Account types with built-in support for
// paged results from the API
type Accounts struct {
	Items []Account `json:"Items"`
	PagedResults
}

// Account represents account details used for deployments, including
// username/password, tokens, Azure and AWS credentials, and SSH key pairs
type Account struct {
	AccountType                     enum.AccountType            `json:"AccountType"`
	Description                     string                      `json:"Description,omitempty"`
	EnvironmentIDs                  []string                    `json:"EnvironmentIds"`
	Name                            string                      `json:"Name"`
	TenantedDeploymentParticipation enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantTags                      []string                    `json:"TenantTags"`
	TenantIds                       []string                    `json:"TenantIds"`
	SpaceID                         string                      `json:"SpaceId,omitempty"`
	Token                           *SensitiveValue             `json:"Token,omitempty"`
	Username                        string                      `json:"Username,omitempty"`
	Password                        *SensitiveValue             `json:"Password,omitempty"`
	AwsServicePrincipalResource
	AzureServicePrincipalResource
	Resource
}

// NewAccount initializes an Account with a name and account type
func NewAccount(name string, accountType enum.AccountType) (*Account, error) {
	if len(strings.Trim(name, " ")) == 0 {
		return nil, errors.New("client: invalid account name")
	}

	return &Account{
		Name:        name,
		AccountType: accountType,
	}, nil
}
