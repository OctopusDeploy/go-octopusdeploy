package credentials

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const templateV2 = "/api/{spaceId}/git-credentials{/id}/v2{?skip,take,name}"

type CreateGitCredentialResponseV2 struct {
	ID string `json:"Id"`
}

type getGitCredentialByIdResponseV2 struct {
	GitCredential *Resource `json:"GitCredential"`
}

// AddV2 creates a new Git credential and returns the ID of the newly-created credential
func AddV2(client newclient.Client, gitCredential *Resource) (*CreateGitCredentialResponseV2, error) {
	if gitCredential == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterGitCredential)
	}

	return newclient.Add[CreateGitCredentialResponseV2](client, templateV2, gitCredential.SpaceID, gitCredential)
}

// GetV2 returns a page of Git credentials matching the query
func GetV2(client newclient.Client, spaceID string, query Query) (*resources.Resources[*Resource], error) {
	return newclient.GetByQuery[Resource](client, templateV2, spaceID, query)
}

// GetByIDV2 returns the Git credential or an error
func GetByIDV2(client newclient.Client, spaceID string, ID string) (*Resource, error) {
	if internal.IsEmpty(ID) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(templateV2, map[string]any{
		"spaceId": spaceID,
		"id":      ID,
	})
	if err != nil {
		return nil, err
	}

	response, err := newclient.Get[getGitCredentialByIdResponseV2](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return response.GitCredential, nil
}

// UpdateV2 modifies a Git credential
func UpdateV2(client newclient.Client, gitCredential *Resource) error {
	if gitCredential == nil {
		return internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterGitCredential)
	}

	_, err := newclient.Update[Resource](client, templateV2, gitCredential.SpaceID, gitCredential.GetID(), gitCredential)
	return err
}
