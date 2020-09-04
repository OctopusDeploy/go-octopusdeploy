package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
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
	AccountType                     enum.AccountType            `json:"AccountType" validate:"required"`
	Description                     string                      `json:"Description,omitempty"`
	EnvironmentIDs                  []string                    `json:"EnvironmentIds,omitempty"`
	Name                            string                      `json:"Name" validate:"required"`
	TenantedDeploymentParticipation enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantTags                      []string                    `json:"TenantTags,omitempty"`
	TenantIDs                       []string                    `json:"TenantIds,omitempty"`
	SpaceID                         string                      `json:"SpaceId,omitempty"`
	Token                           *SensitiveValue             `json:"Token,omitempty"`
	Username                        string                      `json:"Username,omitempty"`
	Password                        *SensitiveValue             `json:"Password,omitempty"`
	AwsServicePrincipalResource
	AzureServicePrincipalResource
	Resource
}

// NewAccount initializes an account with a name and account type.
func NewAccount(name string, accountType enum.AccountType) (*Account, error) {
	if len(strings.Trim(name, " ")) == 0 {
		return nil, errors.New("client: invalid account name")
	}

	return &Account{
		Name:        name,
		AccountType: accountType,
	}, nil
}

func (account *Account) Validate() error {
	validate := validator.New()
	err := validate.Struct(account)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return err
	}

	switch account.AccountType {
	case enum.UsernamePassword:
		return validateUsernamePasswordAccount(account)
	case enum.AzureSubscription:
		return validateAzureSubscriptionAccount(account)
	case enum.AzureServicePrincipal:
		return validateAzureServicePrincipalAccount(account)
	case enum.SshKeyPair:
		return validateSSHKeyAccount(account)
	}

	return nil
}

func validateUsernamePasswordAccount(account *Account) error {
	validations := []error{
		ValidateRequiredPropertyValue("username", account.Username),
	}

	return ValidateMultipleProperties(validations)
}

func validateSSHKeyAccount(account *Account) error {
	validations := []error{
		ValidateRequiredPropertyValue("name", account.Name),
	}

	return ValidateMultipleProperties(validations)
}

func validateAzureServicePrincipalAccount(account *Account) error {
	validations := []error{
		ValidateRequiredUUID("ClientID", account.ClientID),
		ValidateRequiredUUID("SubscriptionNumber", account.SubscriptionNumber),
		ValidateRequiredUUID("TenantID", account.TenantID),
	}

	return ValidateMultipleProperties(validations)
}

func validateAzureSubscriptionAccount(account *Account) error {
	validations := []error{
		ValidateRequiredUUID("ClientID", account.ClientID),
		ValidateRequiredUUID("SubscriptionNumber", account.SubscriptionNumber),
	}
	return ValidateMultipleProperties(validations)
}

var _ ResourceInterface = &Account{}
