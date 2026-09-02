package approvalrules

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// ApprovalRuleScopingStrategy selects whether a rule is scoped by tags or by ids.
type ApprovalRuleScopingStrategy string

const (
	ApprovalRuleScopingStrategyTag ApprovalRuleScopingStrategy = "Tag"
	ApprovalRuleScopingStrategyId  ApprovalRuleScopingStrategy = "Id"
)

// TenantApprovalStrategy determines how tenanted deployments are gated.
type TenantApprovalStrategy string

const (
	// TenantApprovalStrategyPerRelease creates a single change request that covers
	// every tenant of a release and environment.
	TenantApprovalStrategyPerRelease TenantApprovalStrategy = "PerRelease"

	// TenantApprovalStrategyPerTenant creates a separate change request, requiring
	// its own approvals, for each tenant.
	TenantApprovalStrategyPerTenant TenantApprovalStrategy = "PerTenant"
)

// ApprovalRuleTagScope scopes a rule by project/environment tags (ScopingStrategy = "Tag").
// Each scope requires at least one project tag and at least one environment tag.
type ApprovalRuleTagScope struct {
	Id              string   `json:"Id,omitempty"`
	ProjectTags     []string `json:"ProjectTags"`
	EnvironmentTags []string `json:"EnvironmentTags"`
}

// ApprovalRuleIdScope scopes a rule by a project + specific environments
// (ScopingStrategy = "Id"). Each scope requires at least one environment id, and
// every environment must be reachable from the project's lifecycles.
type ApprovalRuleIdScope struct {
	Id             string   `json:"Id,omitempty"`
	ProjectId      string   `json:"ProjectId"`
	EnvironmentIds []string `json:"EnvironmentIds"`
}

// ApprovalRule is the reusable approval configuration applied to a set of
// projects/environments (by tag or by id).
//
// Update replaces the whole resource: Name, Description, MinimumApproversRequired
// and the approver lists are always taken from the request, so omitting
// Description clears it and omitting MinimumApproversRequired resets it to 2.
// Always send a fully populated rule, ideally one returned by GetByID.
//
// The stored ScopingStrategy follows whichever scope list is sent: supplying
// TagScopes forces "Tag" and clears IdScopes, and supplying IdScopes forces "Id"
// and clears TagScopes.
type ApprovalRule struct {
	SpaceID     string `json:"SpaceId,omitempty"`
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description,omitempty"`

	ScopingStrategy ApprovalRuleScopingStrategy `json:"ScopingStrategy"`

	// TenantApprovalStrategy determines how tenanted deployments are gated. When
	// several rules match a deployment the most restrictive wins, so any matching
	// rule set to "PerTenant" results in per-tenant change requests.
	TenantApprovalStrategy TenantApprovalStrategy `json:"TenantApprovalStrategy"`

	TagScopes []ApprovalRuleTagScope `json:"TagScopes"`
	IdScopes  []ApprovalRuleIdScope  `json:"IdScopes"`

	// MinimumApproversRequired is accepted in the range 1-99. It is omitted when
	// zero, which is never a valid value, so that the server applies its default
	// of 2 rather than rejecting the request.
	MinimumApproversRequired int `json:"MinimumApproversRequired,omitempty"`

	// AllowSelfApproval controls whether the deployment creator may approve their own
	// deployment. There is no separate "block by creator" field — this is it, inverted.
	AllowSelfApproval bool `json:"AllowSelfApproval"`

	// IsDisabled disables the rule. There is no "Enabled" field — this is it, inverted.
	// An enabled rule must have at least one scope.
	IsDisabled bool `json:"IsDisabled"`

	// ApprovingUserIds and ApprovingTeamIds must not both be empty.
	ApprovingUserIds []string `json:"ApprovingUserIds"`
	ApprovingTeamIds []string `json:"ApprovingTeamIds"`

	resources.Resource
}

// NewApprovalRule creates an ApprovalRule with server-friendly defaults
// (Tag scoping, per-release tenant approvals, 2 minimum approvers, and empty
// scope/approver slices).
func NewApprovalRule(name string) *ApprovalRule {
	return &ApprovalRule{
		Name:                     name,
		ScopingStrategy:          ApprovalRuleScopingStrategyTag,
		TenantApprovalStrategy:   TenantApprovalStrategyPerRelease,
		TagScopes:                []ApprovalRuleTagScope{},
		IdScopes:                 []ApprovalRuleIdScope{},
		MinimumApproversRequired: 2,
		ApprovingUserIds:         []string{},
		ApprovingTeamIds:         []string{},
		Resource:                 *resources.NewResource(),
	}
}

func (r *ApprovalRule) GetName() string     { return r.Name }
func (r *ApprovalRule) SetName(name string) { r.Name = name }
