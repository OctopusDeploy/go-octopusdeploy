package platformhubpolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/{gitRef}/policies{/slug}{?skip,take,partialName}"

// Add creates a new Platform Hub policy.
func Add(client newclient.Client, policy PlatformHubPolicy, commitMessage string) (*PlatformHubPolicy, error) {
	command, path, commandError := buildAddCommand(client, policy, commitMessage)
	if commandError != nil {
		return nil, commandError
	}

	createdPolicy, addError := newclient.Post[PlatformHubPolicy](client.HttpSession(), path, command)
	if addError != nil {
		return nil, addError
	}

	return createdPolicy, nil
}

// PoliciesQuery represents query parameters for listing policies.
type PoliciesQuery struct {
	GitRef      string `uri:"gitRef" json:"gitRef"`
	PartialName string `uri:"partialName,omitempty" json:"partialName,omitempty"`
	Skip        int    `uri:"skip,omitempty" json:"skip,omitempty"`
	Take        int    `uri:"take,omitempty" json:"take,omitempty"`
}

// PoliciesResponse represents the response from listing Platform Hub policies.
type PoliciesResponse struct {
	Policies           []PlatformHubPolicy `json:"Policies"`
	ItemsPerPage       int                 `json:"ItemsPerPage"`
	FilteredItemsCount int                 `json:"FilteredItemsCount"`
	TotalItemsCount    int                 `json:"TotalItemsCount"`
}

// Get returns a collection of Platform Hub policies based on the provided query.
func Get(client newclient.Client, query PoliciesQuery) (*PoliciesResponse, error) {
	path, pathError := client.URITemplateCache().Expand(template, query)
	if pathError != nil {
		return nil, pathError
	}

	response, responseError := newclient.Get[PoliciesResponse](client.HttpSession(), path)
	if responseError != nil {
		return nil, responseError
	}

	return response, nil
}

// GetBySlug returns the Platform Hub policy that matches given GitRef and Slug.
func GetBySlug(client newclient.Client, gitRef string, slug string) (*PlatformHubPolicy, error) {
	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": gitRef, "slug": slug})
	if pathError != nil {
		return nil, pathError
	}

	policy, err := newclient.Get[PlatformHubPolicy](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

// Update modifies an existing Platform Hub policy.
func Update(client newclient.Client, policy PlatformHubPolicy, commitMessage string) (*PlatformHubPolicy, error) {
	command, path, commandError := buildUpdateCommand(client, policy, commitMessage)
	if commandError != nil {
		return nil, commandError
	}

	updatedPolicy, updateError := newclient.Put[PlatformHubPolicy](client.HttpSession(), path, command)
	if updateError != nil {
		return nil, updateError
	}

	return updatedPolicy, nil
}

func buildAddCommand(client newclient.Client, policy PlatformHubPolicy, commitMessage string) (platformHubPolicyUpsertCommand, string, error) {
	if validationError := policy.Validate(); validationError != nil {
		return platformHubPolicyUpsertCommand{}, "", internal.CreateValidationFailureError("Add", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GitRef})
	if pathError != nil {
		return platformHubPolicyUpsertCommand{}, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		PlatformHubPolicy: policy,
	}

	return command, path, nil
}

func buildUpdateCommand(client newclient.Client, policy PlatformHubPolicy, commitMessage string) (platformHubPolicyUpsertCommand, string, error) {
	if validationError := policy.Validate(); validationError != nil {
		return platformHubPolicyUpsertCommand{}, "", internal.CreateValidationFailureError("Update", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GitRef, "slug": policy.Slug})
	if pathError != nil {
		return platformHubPolicyUpsertCommand{}, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		PlatformHubPolicy: policy,
	}

	return command, path, nil
}

type platformHubPolicyUpsertCommand struct {
	ChangeDescription string `json:"ChangeDescription,omitempty"`

	PlatformHubPolicy
}
