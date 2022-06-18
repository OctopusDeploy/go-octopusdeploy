package accounts

import (
	"encoding/json"

	resources "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

// Accounts defines a collection of accounts with built-in support for paged
// results.
type Accounts struct {
	Items []IAccount `json:"Items"`
	resources.PagedResults
}

// MarshalJSON returns an Accounts struct as its JSON encoding.
func (a *Accounts) MarshalJSON() ([]byte, error) {
	accounts := struct {
		Items []IAccount `json:"Items"`
		resources.PagedResults
	}{
		Items:        a.Items,
		PagedResults: a.PagedResults,
	}

	return json.Marshal(accounts)
}

// UnmarshalJSON sets this Accounts struct to its representation in JSON.
func (a *Accounts) UnmarshalJSON(b []byte) error {
	var accounts map[string]*json.RawMessage
	err := json.Unmarshal(b, &accounts)
	if err != nil {
		return err
	}

	var pagedResults resources.PagedResults
	err = json.Unmarshal(b, &pagedResults)
	if err != nil {
		return err
	}
	a.PagedResults = pagedResults

	var accountType AccountType
	var itemsArray []*json.RawMessage
	var items *json.RawMessage

	if accounts["Items"] != nil {
		err = json.Unmarshal(*accounts["Items"], &items)
		if err != nil {
			return err
		}

		err = json.Unmarshal(*items, &itemsArray)
		if err != nil {
			return err
		}

		for _, account := range itemsArray {
			var accountProperties map[string]*json.RawMessage
			err = json.Unmarshal(*account, &accountProperties)
			if err != nil {
				return err
			}

			if accountProperties["AccountType"] != nil {
				at := accountProperties["AccountType"]
				err := json.Unmarshal(*at, &accountType)
				if err != nil {
					return err
				}
			}

			switch accountType {
			case AccountTypeAmazonWebServicesAccount:
				var amazonWebServicesAccount *AmazonWebServicesAccount
				err := json.Unmarshal(*account, &amazonWebServicesAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, amazonWebServicesAccount)
			case AccountTypeAzureServicePrincipal:
				var azureServicePrincipalAccount *AzureServicePrincipalAccount
				err := json.Unmarshal(*account, &azureServicePrincipalAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, azureServicePrincipalAccount)
			case AccountTypeAzureSubscription:
				var azureSubscriptionAccount *AzureSubscriptionAccount
				err := json.Unmarshal(*account, &azureSubscriptionAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, azureSubscriptionAccount)
			case AccountTypeGoogleCloudPlatformAccount:
				var googleCloudAccount *GoogleCloudPlatformAccount
				err := json.Unmarshal(*account, &googleCloudAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, googleCloudAccount)
			case AccountTypeSSHKeyPair:
				var sshKeyAccount *SSHKeyAccount
				err := json.Unmarshal(*account, &sshKeyAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, sshKeyAccount)
			case AccountTypeToken:
				var tokenAccount *TokenAccount
				err := json.Unmarshal(*account, &tokenAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, tokenAccount)
			case AccountTypeUsernamePassword:
				var usernamePasswordAccount *UsernamePasswordAccount
				err := json.Unmarshal(*account, &usernamePasswordAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, usernamePasswordAccount)
			}
		}
	}

	return nil
}
