package accounts

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// Accounts defines a collection of accounts with built-in support for paged
// results.
type Accounts resources.Resources[IAccount]

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
			case AccountTypeAwsOIDC:
				var awsOIDCAccount *AwsOIDCAccount
				err := json.Unmarshal(*account, &awsOIDCAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, awsOIDCAccount)
			case AccountTypeAzureServicePrincipal:
				var azureServicePrincipalAccount *AzureServicePrincipalAccount
				err := json.Unmarshal(*account, &azureServicePrincipalAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, azureServicePrincipalAccount)
			case AccountTypeAzureOIDC:
				var azureOIDCAccount *AzureOIDCAccount
				err := json.Unmarshal(*account, &azureOIDCAccount)
				if err != nil {
					return err
				}
				a.Items = append(a.Items, azureOIDCAccount)
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

// GetNextPage retrives the next page from the links collection. If no next page
// exists it will return nill
func (r *Accounts) GetNextPage(client *sling.Sling) (*Accounts, error) {
	if r.Links.PageNext == "" {
		return nil, nil
	}
	response, err := api.ApiGet(client, new(resources.Resources[*AccountResource]), r.Links.PageNext)
	if err != nil {
		return nil, err
	}
	return ToAccounts(response.(*resources.Resources[*AccountResource])), nil
}

// GetAllPages will retrive all remaining next pages in the link collection
// and return the result as list of concatenated Items; Including the items
// from the base Resource.
func (r *Accounts) GetAllPages(client *sling.Sling) ([]IAccount, error) {
	items := make([]IAccount, 0)
	res := r
	var err error
	for res != nil {
		items = append(items, res.Items...)
		res, err = res.GetNextPage(client)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}
