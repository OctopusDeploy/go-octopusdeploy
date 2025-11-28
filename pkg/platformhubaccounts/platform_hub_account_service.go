package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/accounts{/id}"

// Add creates a new Platform Hub account.
func Add(client newclient.Client, platformHubAccount IPlatformHubAccount) (IPlatformHubAccount, error) {
	if platformHubAccount == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccount")
	}

	if platformHubAccount.GetName() == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccount.Name")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[PlatformHubAccountResource](client.HttpSession(), path, platformHubAccount)
	if err != nil {
		return nil, err
	}

	if res.ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccountID")
	}

	//create only returns the id of the new resource, so we need to get the full object
	return GetByID(client, res.ID)
}

// GetByID returns the Platform Hub account that matches the input ID.
func GetByID(client newclient.Client, id string) (IPlatformHubAccount, error) {
	if id == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("id")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}

	// Get the resource and convert to the appropriate concrete type
	resource, err := newclient.Get[PlatformHubAccountResource](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return resource.ToPlatformHubAccount()
}

// Update modifies a Platform Hub account.
func Update(client newclient.Client, platformHubAccount IPlatformHubAccount) (IPlatformHubAccount, error) {
	if platformHubAccount == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccount")
	}

	if platformHubAccount.GetID() == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccount.ID")
	}

	if platformHubAccount.GetName() == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubAccount.Name")
	}

	accountResource, err := ToPlatformHubAccountResource(platformHubAccount)
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"id": platformHubAccount.GetID()})
	if err != nil {
		return nil, err
	}

	_, err = newclient.Put[PlatformHubAccountResource](client.HttpSession(), path, accountResource)
	if err != nil {
		return nil, err
	}

	//modify doesn't return the updated object, so we return the input object after a successful update
	return platformHubAccount, nil
}

// Delete removes a Platform Hub account with the specified ID.
func Delete(client newclient.Client, id string) error {
	if id == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("id")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"id": id})
	if err != nil {
		return err
	}

	err = newclient.Delete(client.HttpSession(), path)
	if err != nil {
		return err
	}
	return nil
}
