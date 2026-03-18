package platformhubpolicies

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPolicyValidate_Name(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// Valid
	policy.SetName("Valid Name")
	require.NoError(t, policy.Validate())

	// Invalid
	policy.SetName("")
	require.ErrorContains(t, policy.Validate(), "Name")

	policy.SetName("   ")
	require.ErrorContains(t, policy.Validate(), "Name")
}

func TestPolicyValidate_GitRef(t *testing.T) {
	policy := newPolicyBuilder().WithGitRef("main").Build()

	// Valid
	require.NoError(t, policy.Validate())

	// Invalid
	policy = newPolicyBuilder().WithGitRef("").Build()
	require.ErrorContains(t, policy.Validate(), "GitRef")

	policy = newPolicyBuilder().WithGitRef("   ").Build()
	require.ErrorContains(t, policy.Validate(), "GitRef")
}

func TestPolicyValidate_Slug(t *testing.T) {
	policy := newPolicyBuilder().WithSlug("valid_slug").Build()

	// Valid
	require.NoError(t, policy.Validate())

	// Invalid
	policy = newPolicyBuilder().WithSlug("").Build()
	require.ErrorContains(t, policy.Validate(), "Slug")

	policy = newPolicyBuilder().WithSlug("   ").Build()
	require.ErrorContains(t, policy.Validate(), "Slug")
}

func TestPolicyValidate_Description(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// Description is optional
	policy.SetDescription("Description")
	require.NoError(t, policy.Validate())

	policy.SetDescription("")
	require.NoError(t, policy.Validate())
}

func TestPolicyValidate_ScopeRego(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// Valid
	policy.SetScopeRego("package scope")
	require.NoError(t, policy.Validate())

	// Invalid
	policy.SetScopeRego("")
	require.ErrorContains(t, policy.Validate(), "ScopeRego")

	policy.SetScopeRego("   ")
	require.ErrorContains(t, policy.Validate(), "ScopeRego")
}

func TestPolicyValidate_ConditionsRego(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// Valid
	policy.SetConditionsRego("package conditions")
	require.NoError(t, policy.Validate())

	// Invalid
	policy.SetConditionsRego("")
	require.ErrorContains(t, policy.Validate(), "ConditionsRego")

	policy.SetConditionsRego("   ")
	require.ErrorContains(t, policy.Validate(), "ConditionsRego")
}

func TestPolicyValidate_ViolationAction(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// Valid
	policy.SetViolationAction("block")
	require.NoError(t, policy.Validate())

	// Invalid
	policy.SetViolationAction("")
	require.ErrorContains(t, policy.Validate(), "ViolationAction")

	policy.SetViolationAction("   ")
	require.ErrorContains(t, policy.Validate(), "ViolationAction")
}

func TestPolicyValidate_ViolationReason(t *testing.T) {
	policy := newPolicyBuilder().Build()

	// ViolationReason is optional
	policy.SetViolationReason("Some reason")
	require.NoError(t, policy.Validate())

	policy.SetViolationReason("")
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

func (b *policyBuilder) WithName(name string) *policyBuilder {
	b.name = name
	return b
}

func (b *policyBuilder) WithGitRef(gitRef string) *policyBuilder {
	b.gitRef = gitRef
	return b
}

func (b *policyBuilder) WithSlug(slug string) *policyBuilder {
	b.slug = slug
	return b
}

func (b *policyBuilder) WithDescription(description string) *policyBuilder {
	b.description = description
	return b
}

func (b *policyBuilder) WithScopeRego(scopeRego string) *policyBuilder {
	b.scopeRego = scopeRego
	return b
}

func (b *policyBuilder) WithConditionsRego(conditionsRego string) *policyBuilder {
	b.conditionsRego = conditionsRego
	return b
}

func (b *policyBuilder) WithViolationReason(violationReason string) *policyBuilder {
	b.violationReason = violationReason
	return b
}

func (b *policyBuilder) WithViolationAction(violationAction string) *policyBuilder {
	b.violationAction = violationAction
	return b
}

func (b *policyBuilder) Build() Policy {
	return &persistedPolicy{
		Name:            b.name,
		GitRef:          b.gitRef,
		Slug:            b.slug,
		ScopeRego:       b.scopeRego,
		ConditionsRego:  b.conditionsRego,
		ViolationAction: b.violationAction,
		Description:     b.description,
		ViolationReason: b.violationReason,
	}
}

func (b *policyBuilder) BuildCandidate() *PolicyCandidate {
	return &PolicyCandidate{
		Name:            b.name,
		GitRef:          b.gitRef,
		Slug:            b.slug,
		ScopeRego:       b.scopeRego,
		ConditionsRego:  b.conditionsRego,
		ViolationAction: b.violationAction,
		Description:     b.description,
		ViolationReason: b.violationReason,
	}
}
