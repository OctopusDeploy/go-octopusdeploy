package octopusdeploy

import (
	"fmt"
	"strings"
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
	AccountType                     AccountType            `json:"AccountType"`
	Description                     string                 `json:"Description,omitempty"`
	EnvironmentIDs                  []string               `json:"EnvironmentIds"`
	Name                            string                 `json:"Name"`
	TenantedDeploymentParticipation TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantTags                      []string               `json:"TenantTags"`
	TenantIds                       []string               `json:"TenantIds"`
	SpaceID                         string                 `json:"SpaceId,omitempty"`
	Token                           *SensitiveValue        `json:"Token,omitempty"`
	Resource
	AzureServicePrincipalResource
	AwsServicePrincipalResource

	Username string          `json:"Username,omitempty"`
	Password *SensitiveValue `json:"Password,omitempty"`
}

// NewAccount initializes an Account with a name and account type
func NewAccount(name string, accountType AccountType) (*Account, error) {
	if len(strings.Trim(name, " ")) == 0 {
		return nil, fmt.Errorf("Invalid account name")
	}

	return &Account{
		Name:        name,
		AccountType: accountType,
	}, nil
}

func (s *AccountService) Get(accountId string) (*Account, error) {
	path := fmt.Sprintf("accounts/%s", accountId)
	resp, err := apiGet(s.sling, new(Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}

func (s *AccountService) GetAll() (*[]Account, error) {
	var p []Account

	path := "accounts"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Accounts), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Accounts)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *AccountService) GetByName(accountName string) (*Account, error) {
	var foundAccount Account
	accounts, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, account := range *accounts {
		if account.Name == accountName {
			return &account, nil
		}
	}

	return &foundAccount, fmt.Errorf("no account found with account name %s", accountName)
}

func (s *AccountService) Add(account *Account) (*Account, error) {
	resp, err := apiAdd(s.sling, account, new(Account), "accounts")

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}

func (s *AccountService) Delete(accountID string) error {
	path := fmt.Sprintf("accounts/%s", accountID)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) Update(account *Account) (*Account, error) {
	path := fmt.Sprintf("accounts/%s", account.ID)
	resp, err := apiUpdate(s.sling, account, new(Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Account), nil
}
