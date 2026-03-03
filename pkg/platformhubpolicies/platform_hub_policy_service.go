package platformhubpolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/{gitRef}/policies{/slug}"

// Add creates a new Platform Hub policy.
func Add(client newclient.Client, policy *PlatformHubPolicy, commitMessage string) (*PlatformHubPolicy, error) {
	if policy == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("policy")
	}

	if validationError := policy.Validate(); validationError != nil {
		return nil, internal.CreateValidationFailureError("Add", validationError)
	}

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
func Update(client newclient.Client, policy *PlatformHubPolicy, commitMessage string) (*PlatformHubPolicy, error) {
	command, path, err := buildUpdateCommand(client, policy, commitMessage)
	if err != nil {
		return nil, err
	}

	updatedPolicy, err := newclient.Put[PlatformHubPolicy](client.HttpSession(), path, command)
	if err != nil {
		return nil, err
	}
	return updatedPolicy, nil
}

func buildAddCommand(client newclient.Client, policy *PlatformHubPolicy, commitMessage string) (*platformHubPolicyUpsertCommand, string, error) {
	if policy == nil {
		return nil, "", internal.CreateRequiredParameterIsEmptyOrNilError("policy")
	}

	if validationError := policy.Validate(); validationError != nil {
		return nil, "", internal.CreateValidationFailureError("Add", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GitRef})
	if pathError != nil {
		return nil, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		PlatformHubPolicy: *policy,
	}

	return &command, path, nil
}

func buildUpdateCommand(client newclient.Client, policy *PlatformHubPolicy, commitMessage string) (*platformHubPolicyUpsertCommand, string, error) {
	if policy == nil {
		return nil, "", internal.CreateRequiredParameterIsEmptyOrNilError("policy")
	}

	if validationError := policy.Validate(); validationError != nil {
		return nil, "", internal.CreateValidationFailureError("Update", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GitRef, "slug": policy.Slug})
	if pathError != nil {
		return nil, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		PlatformHubPolicy: *policy,
	}

	return &command, path, nil
}

type platformHubPolicyUpsertCommand struct {
	ChangeDescription string `json:"ChangeDescription,omitempty"`

	PlatformHubPolicy
}
