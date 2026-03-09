package platformhubpolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/{gitRef}/policies"

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

type platformHubPolicyUpsertCommand struct {
	ChangeDescription string `json:"ChangeDescription,omitempty"`

	PlatformHubPolicy
}
