package platformhubgitcredential

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/git-credentials{/id}"

// Add creates a new Platform Hub git credential.
func Add(client newclient.Client, platformHubGitCredential *PlatformHubGitCredential) (*PlatformHubGitCredential, error) {
	if platformHubGitCredential == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential")
	}

	if platformHubGitCredential.Name == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.Name")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[PlatformHubGitCredential](client.HttpSession(), path, platformHubGitCredential)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetByID returns the Platform Hub git credential that matches the input ID.
func GetByID(client newclient.Client, id string) (*PlatformHubGitCredential, error) {
	if id == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("id")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[PlatformHubGitCredential](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a Platform Hub git credential.
func Update(client newclient.Client, platformHubGitCredential *PlatformHubGitCredential) (*PlatformHubGitCredential, error) {
	if platformHubGitCredential == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential")
	}

	if platformHubGitCredential.ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.ID")
	}

	if platformHubGitCredential.Name == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential.Name")
	}

	path, err := client.URITemplateCache().Expand(template, map[string]any{"id": platformHubGitCredential.ID})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Put[PlatformHubGitCredential](client.HttpSession(), path, platformHubGitCredential)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete removes a Platform Hub git credential with the specified ID.
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
