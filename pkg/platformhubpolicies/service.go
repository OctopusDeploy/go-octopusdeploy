package platformhubpolicies

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/platformhub/{gitRef}/policies{/slug}{?skip,take,partialName}"

// Add creates and stores new policy in Platform Hub.
func Add(client newclient.Client, policy PolicyCandidate, commitMessage string) (Policy, error) {
	command, path, commandError := buildAddCommand(client, policy, commitMessage)
	if commandError != nil {
		return nil, commandError
	}

	createdPolicy, addError := newclient.Post[persistedPolicy](client.HttpSession(), path, command)
	if addError != nil {
		return nil, addError
	}

	return createdPolicy, nil
}

// List returns a paginated collection of Platform Hub policies based on the provided query.
func List(client newclient.Client, query PoliciesQuery) (*PoliciesQueryResult, error) {
	path, pathError := client.URITemplateCache().Expand(template, query)
	if pathError != nil {
		return nil, pathError
	}

	result, resultError := newclient.Get[PoliciesQueryResult](client.HttpSession(), path)
	if resultError != nil {
		return nil, resultError
	}

	return result, nil
}

// GetBySlug returns the Platform Hub policy that matches given policy key.
func GetBySlug(client newclient.Client, key PolicyKey) (Policy, error) {
	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": key.GetGitRef(), "slug": key.GetSlug()})
	if pathError != nil {
		return nil, pathError
	}

	policy, err := newclient.Get[persistedPolicy](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

// Update stores modified policy in Platform Hub.
func Update(client newclient.Client, policy Policy, commitMessage string) (Policy, error) {
	command, path, commandError := buildUpdateCommand(client, policy, commitMessage)
	if commandError != nil {
		return nil, commandError
	}

	updatedPolicy, updateError := newclient.Put[persistedPolicy](client.HttpSession(), path, command)
	if updateError != nil {
		return nil, updateError
	}

	return updatedPolicy, nil
}

// Publish publishes a Platform Hub policy version.
func Publish(client newclient.Client, policy PolicyKey, version string) (PublishedPolicy, error) {
	command, path, err := buildPublishCommand(client, policy, version)
	if err != nil {
		return nil, err
	}

	publishedVersion, postErr := newclient.Post[publishedPolicyVersion](client.HttpSession(), path, command)
	if postErr != nil {
		return nil, postErr
	}

	return publishedVersion, nil
}

// GetVersions returns published versions of a Platform Hub policy.
func GetVersions(client newclient.Client, query PublishedPoliciesQuery) (*PublishedPoliciesQueryResult, error) {
	path, pathError := buildGetVersionsPath(client, query)
	if pathError != nil {
		return nil, pathError
	}

	result, err := newclient.Get[PublishedPoliciesQueryResult](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ActivateVersion activates a published Platform Hub policy version.
func ActivateVersion(client newclient.Client, policy PublishedPolicyKey) (PublishedPolicy, error) {
	return modifyVersionStatus(client, policy, true)
}

// DeactivateVersion deactivates a published Platform Hub policy version.
func DeactivateVersion(client newclient.Client, policy PublishedPolicyKey) (PublishedPolicy, error) {
	return modifyVersionStatus(client, policy, false)
}

// PoliciesQuery represents query parameters for listing policies.
type PoliciesQuery struct {
	GitRef      string `uri:"gitRef" json:"gitRef"`
	PartialName string `uri:"partialName,omitempty" json:"partialName,omitempty"`
	Skip        int    `uri:"skip,omitempty" json:"skip,omitempty"`
	Take        int    `uri:"take,omitempty" json:"take,omitempty"`
}

// PoliciesQueryResult paginated collection of policies
type PoliciesQueryResult struct {
	Policies           []Policy
	ItemsPerPage       int
	FilteredItemsCount int
	TotalItemsCount    int
}

func (r *PoliciesQueryResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		Policies           []persistedPolicy `json:"Policies"`
		ItemsPerPage       int               `json:"ItemsPerPage"`
		FilteredItemsCount int               `json:"FilteredItemsCount"`
		TotalItemsCount    int               `json:"TotalItemsCount"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.ItemsPerPage = raw.ItemsPerPage
	r.FilteredItemsCount = raw.FilteredItemsCount
	r.TotalItemsCount = raw.TotalItemsCount
	r.Policies = make([]Policy, len(raw.Policies))

	for i := range raw.Policies {
		r.Policies[i] = &raw.Policies[i]
	}

	return nil
}

// PublishedPoliciesQuery query parameters for listing published policy versions.
type PublishedPoliciesQuery struct {
	Slug string `uri:"slug"`
	Skip int    `uri:"skip,omitempty"`
	Take int    `uri:"take,omitempty"`
}

// PublishedPoliciesQueryResult paginated collection of published policy versions
type PublishedPoliciesQueryResult struct {
	Items        []PublishedPolicy
	ItemsPerPage int
	TotalResults int
}

func (r *PublishedPoliciesQueryResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		Items        []publishedPolicyVersion `json:"Items"`
		ItemsPerPage int                      `json:"ItemsPerPage"`
		TotalResults int                      `json:"TotalResults"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.ItemsPerPage = raw.ItemsPerPage
	r.TotalResults = raw.TotalResults
	r.Items = make([]PublishedPolicy, len(raw.Items))

	for i := range raw.Items {
		r.Items[i] = &raw.Items[i]
	}

	return nil
}

func buildAddCommand(client newclient.Client, candidate PolicyCandidate, commitMessage string) (platformHubPolicyUpsertCommand, string, error) {
	policy := persistedPolicy{
		GitRef:          candidate.GitRef,
		Slug:            candidate.Slug,
		Name:            candidate.Name,
		Description:     candidate.Description,
		ScopeRego:       candidate.ScopeRego,
		ConditionsRego:  candidate.ConditionsRego,
		ViolationAction: candidate.ViolationAction,
		ViolationReason: candidate.ViolationReason,
	}

	if validationError := policy.Validate(); validationError != nil {
		return platformHubPolicyUpsertCommand{}, "", internal.CreateValidationFailureError("Add", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GitRef})
	if pathError != nil {
		return platformHubPolicyUpsertCommand{}, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		persistedPolicy:   policy,
	}
	return command, path, nil
}

func buildUpdateCommand(client newclient.Client, policy Policy, commitMessage string) (platformHubPolicyUpsertCommand, string, error) {
	if validationError := policy.Validate(); validationError != nil {
		return platformHubPolicyUpsertCommand{}, "", internal.CreateValidationFailureError("Update", validationError)
	}

	path, pathError := client.URITemplateCache().Expand(template, map[string]any{"gitRef": policy.GetGitRef(), "slug": policy.GetSlug()})
	if pathError != nil {
		return platformHubPolicyUpsertCommand{}, "", pathError
	}

	command := platformHubPolicyUpsertCommand{
		ChangeDescription: commitMessage,
		persistedPolicy: persistedPolicy{
			GitRef:          policy.GetGitRef(),
			Slug:            policy.GetSlug(),
			Name:            policy.GetName(),
			Description:     policy.GetDescription(),
			ScopeRego:       policy.GetScopeRego(),
			ConditionsRego:  policy.GetConditionsRego(),
			ViolationAction: policy.GetViolationAction(),
			ViolationReason: policy.GetViolationReason(),
		},
	}
	return command, path, nil
}

type platformHubPolicyUpsertCommand struct {
	ChangeDescription string `json:"ChangeDescription,omitempty"`

	persistedPolicy
}

type publishCommand struct {
	Version string `json:"Version"`
}

func buildPublishCommand(client newclient.Client, policy PolicyKey, version string) (publishCommand, string, error) {
	parameters := map[string]any{"gitRef": policy.GetGitRef(), "slug": policy.GetSlug()}
	path, pathError := client.URITemplateCache().Expand("/api/platformhub/{gitRef}/policies/{slug}/publish", parameters)
	if pathError != nil {
		return publishCommand{}, "", pathError
	}

	return publishCommand{Version: version}, path, nil
}

func buildGetVersionsPath(client newclient.Client, query PublishedPoliciesQuery) (string, error) {
	return client.URITemplateCache().Expand("/api/platformhub/policies/{slug}/versions/v2{?skip,take}", query)
}

type modifyVersionStatusCommand struct {
	IsActive bool `json:"IsActive"`
}

func modifyVersionStatus(client newclient.Client, policy PublishedPolicyKey, isActive bool) (PublishedPolicy, error) {
	command, path, err := buildModifyVersionStatusCommand(client, policy, isActive)
	if err != nil {
		return nil, err
	}

	modifiedVersion, postErr := newclient.Post[publishedPolicyVersion](client.HttpSession(), path, command)
	if postErr != nil {
		return nil, postErr
	}

	return modifiedVersion, nil
}

func buildModifyVersionStatusCommand(client newclient.Client, policy PublishedPolicyKey, isActive bool) (modifyVersionStatusCommand, string, error) {
	parameters := map[string]any{"slug": policy.GetSlug(), "version": policy.GetVersion()}
	path, pathError := client.URITemplateCache().Expand("/api/platformhub/policies/{slug}/versions/{version}/modify-status", parameters)
	if pathError != nil {
		return modifyVersionStatusCommand{}, "", pathError
	}

	return modifyVersionStatusCommand{IsActive: isActive}, path, nil
}
