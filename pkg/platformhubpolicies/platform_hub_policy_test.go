package platformhubpolicies

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlatformHubPolicyNew_Valid(t *testing.T) {
	newPolicy, err := NewPlatformHubPolicy("Policy One", "refs/heads/main", "policy_one", "package policy_one", "package policy_one", "block")
	require.NoError(t, err)
	require.NotNil(t, newPolicy)
	require.NoError(t, newPolicy.Validate())

	require.Equal(t, "Policy One", newPolicy.GetName())
	require.Equal(t, "refs/heads/main", newPolicy.GetGitRef())
	require.Equal(t, "policy_one", newPolicy.GetSlug())
	require.Empty(t, newPolicy.GetDescription())
	require.Equal(t, "package policy_one", newPolicy.GetScopeRego())
	require.Equal(t, "package policy_one", newPolicy.GetConditionsRego())
	require.Equal(t, "block", newPolicy.GetViolationAction())
	require.Empty(t, newPolicy.GetViolationReason())
}

func TestPlatformHubPolicyNew_Invalid(t *testing.T) {
	builder := newPolicyBuilder()
	// Invalid GitRef
	policy, err := builder.GitRef("").Create()
	require.Error(t, err)
	require.Nil(t, policy)

	policy, err = builder.GitRef("   ").Create()
	require.Error(t, err)
	require.Nil(t, policy)

	// Make GitRef valid
	policy, err = builder.GitRef("refs/heads/main").Create()
	require.NoError(t, err)
	require.NotNil(t, policy)

	// Invalid Slug
	policy, err = newPolicyBuilder().Slug("").Create()
	require.Error(t, err)
	require.Nil(t, policy)

	policy, err = newPolicyBuilder().Slug("   ").Create()
	require.Error(t, err)
	require.Nil(t, policy)
}

func TestPlatformHubPolicyValidate(t *testing.T) {
	policy, err := newPolicyBuilder().Slug("valid_slug").Create()
	require.NoError(t, err)

	// Empty name
	policy.SetName("")
	require.Error(t, policy.Validate())

	policy.SetName("   ")
	require.Error(t, policy.Validate())

	policy.SetName("Valid Name")
	require.NoError(t, policy.Validate())

	// Empty description
	policy.SetDescription("")
	require.NoError(t, policy.Validate())

	// Empty scope rego
	policy.SetScopeRego("")
	require.Error(t, policy.Validate())

	policy.SetScopeRego("   ")
	require.Error(t, policy.Validate())

	policy.SetScopeRego("package valid_slug")
	require.NoError(t, policy.Validate())

	// Empty conditions rego
	policy.SetConditionsRego("")
	require.Error(t, policy.Validate())

	policy.SetConditionsRego("   ")
	require.Error(t, policy.Validate())

	policy.SetConditionsRego("package valid_slug")
	require.NoError(t, policy.Validate())

	// Empty violation reason
	policy.SetViolationReason("")
	require.NoError(t, policy.Validate())

	// Empty violation action
	policy.SetViolationAction("")
	require.Error(t, policy.Validate())

	policy.SetViolationAction("   ")
	require.Error(t, policy.Validate())

	policy.SetViolationAction("block")
	require.NoError(t, policy.Validate())
}

type policyBuilder struct {
	name            string
	gitRef          string
	slug            string
	description     string
	scopeRego       string
	conditionsRego  string
	violationReason string
	violationAction string
}

func newPolicyBuilder() *policyBuilder {
	return &policyBuilder{
		name:            "Dummy",
		gitRef:          "main",
		slug:            "dummy",
		scopeRego:       "package dummy",
		conditionsRego:  "package dummy",
		violationAction: "block",
	}
}

func (b *policyBuilder) Name(name string) *policyBuilder {
	b.name = name
	return b
}

func (b *policyBuilder) GitRef(gitRef string) *policyBuilder {
	b.gitRef = gitRef
	return b
}

func (b *policyBuilder) Slug(slug string) *policyBuilder {
	b.slug = slug
	return b
}

func (b *policyBuilder) Description(description string) *policyBuilder {
	b.description = description
	return b
}

func (b *policyBuilder) ScopeRego(scopeRego string) *policyBuilder {
	b.scopeRego = scopeRego
	return b
}

func (b *policyBuilder) ConditionsRego(conditionsRego string) *policyBuilder {
	b.conditionsRego = conditionsRego
	return b
}

func (b *policyBuilder) ViolationReason(violationReason string) *policyBuilder {
	b.violationReason = violationReason
	return b
}

func (b *policyBuilder) ViolationAction(violationAction string) *policyBuilder {
	b.violationAction = violationAction
	return b
}

func (b *policyBuilder) Create() (*PlatformHubPolicy, error) {
	policy, err := NewPlatformHubPolicy(b.name, b.gitRef, b.slug, b.scopeRego, b.conditionsRego, b.violationAction)
	if err != nil {
		return nil, err
	}

	policy.SetDescription(b.description)
	policy.SetViolationReason(b.violationReason)

	return policy, nil
}
