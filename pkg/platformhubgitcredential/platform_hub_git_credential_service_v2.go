package platformhubgitcredential

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

const templateV2 = "/api/platformhub/git-credentials{/id}/v2{?skip,take,name}"

type CreatePlatformHubGitCredentialResponseV2 struct {
	ID string `json:"Id"`
}

// AddV2 creates a new Platform Hub git credential and returns the ID of the newly-created credential
func AddV2(client newclient.Client, platformHubGitCredential *PlatformHubGitCredential) (*CreatePlatformHubGitCredentialResponseV2, error) {
	if platformHubGitCredential == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential")
	}
	if platformHubGitCredential.Name == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.Name")
	}

	path, err := client.URITemplateCache().Expand(templateV2, map[string]any{})
	if err != nil {
		return nil, err
	}

	return newclient.Post[CreatePlatformHubGitCredentialResponseV2](client.HttpSession(), path, platformHubGitCredential)
}

// GetV2 returns a page of Platform Hub git credentials matching the query
func GetV2(client newclient.Client, query credentials.Query) (*resources.Resources[*PlatformHubGitCredential], error) {
	values, _ := uritemplates.Struct2map(query)
	if values == nil {
		values = map[string]any{}
	}

	path, err := client.URITemplateCache().Expand(templateV2, values)
	if err != nil {
		return nil, err
	}

	return newclient.Get[resources.Resources[*PlatformHubGitCredential]](client.HttpSession(), path)
}

// GetByIDV2 returns the Platform Hub git credential or an error
func GetByIDV2(client newclient.Client, id string) (*PlatformHubGitCredential, error) {
	if id == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("id")
	}

	path, err := client.URITemplateCache().Expand(templateV2, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}

	return newclient.Get[PlatformHubGitCredential](client.HttpSession(), path)
}

// UpdateV2 modifies a Platform Hub git credential
func UpdateV2(client newclient.Client, platformHubGitCredential *PlatformHubGitCredential) error {
	if platformHubGitCredential == nil {
		return internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential")
	}
	if platformHubGitCredential.ID == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.ID")
	}
	if platformHubGitCredential.Name == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.Name")
	}

	path, err := client.URITemplateCache().Expand(templateV2, map[string]any{"id": platformHubGitCredential.ID})
	if err != nil {
		return err
	}

	_, err = newclient.Put[PlatformHubGitCredential](client.HttpSession(), path, platformHubGitCredential)
	return err
}
