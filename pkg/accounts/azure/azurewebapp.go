package azure

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
)

type AzureWebApp struct {
	Name          string `json:"Name,omitempty"`
	Region        string `json:"Region,omitempty"`
	ResourceGroup string `json:"ResourceGroup,omitempty"`
}

type AzureWebAppSlot struct {
	Name          string `json:"Name,omitempty"`
	Site          string `json:"Site,omitempty"`
	Region        string `json:"Region,omitempty"`
	ResourceGroup string `json:"ResourceGroup,omitempty"`
}

func GetWebSites(client client.Client, account accounts.IAccount) ([]*AzureWebApp, error) {
	path := account.GetLinks()[constants.LinkWebSites]
	if path == "" {
		return nil, fmt.Errorf("cannot get websites for account '%s' (%s)", account.GetName(), account.GetID())
	}

	items := []*AzureWebApp{}

	_, err := api.ApiGet(client.Sling(), &items, path)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func GetWebSiteSlots(client client.Client, spAccount accounts.IAccount, app *AzureWebApp) ([]*AzureWebAppSlot, error) {
	path := spAccount.GetLinks()[constants.LinkWebSiteSlots]
	if path == "" {
		return nil, fmt.Errorf("cannot get websites for account '%s' (%s)", spAccount.GetName(), spAccount.GetID())
	}

	path = strings.ReplaceAll(path, "{resourceGroupName}", app.ResourceGroup)
	path = strings.ReplaceAll(path, "{webSiteName}", app.Name)

	items := []*AzureWebAppSlot{}

	_, err := api.ApiGet(client.Sling(), &items, path)
	if err != nil {
		return nil, err
	}

	return items, nil
}
