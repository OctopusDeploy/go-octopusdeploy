package approvalpolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// ApprovalPolicyScopingStrategy selects whether a policy is scoped by tags or by ids.
type ApprovalPolicyScopingStrategy string

const (
	ApprovalPolicyScopingStrategyTag ApprovalPolicyScopingStrategy = "Tag"
	ApprovalPolicyScopingStrategyId  ApprovalPolicyScopingStrategy = "Id"
)

// ApprovalPolicyTagScope scopes a policy by project/environment tags (ScopingStrategy = "Tag").
type ApprovalPolicyTagScope struct {
	Id              string   `json:"Id,omitempty"`
	ProjectTags     []string `json:"ProjectTags"`
	EnvironmentTags []string `json:"EnvironmentTags"`
}

// ApprovalPolicyIdScope scopes a policy by a project + specific environments (ScopingStrategy = "Id").
type ApprovalPolicyIdScope struct {
	Id             string   `json:"Id,omitempty"`
	ProjectId      string   `json:"ProjectId"`
	EnvironmentIds []string `json:"EnvironmentIds"`
}

// ApprovalPolicy is the reusable approval configuration applied to a set of
// projects/environments (by tag or by id).
type ApprovalPolicy struct {
	SpaceID     string `json:"SpaceId,omitempty"`
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description,omitempty"`

	ScopingStrategy ApprovalPolicyScopingStrategy `json:"ScopingStrategy"`
	TagScopes       []ApprovalPolicyTagScope      `json:"TagScopes"`
	IdScopes        []ApprovalPolicyIdScope       `json:"IdScopes"`

	MinimumApproversRequired int `json:"MinimumApproversRequired"`

	// AllowSelfApproval controls whether the deployment creator may approve their own
	// deployment. There is no separate "block by creator" field — this is it, inverted.
	AllowSelfApproval bool `json:"AllowSelfApproval"`

	// IsDisabled disables the policy. There is no "Enabled" field — this is it, inverted.
	IsDisabled bool `json:"IsDisabled"`

	ApprovingUserIds []string `json:"ApprovingUserIds"`
	ApprovingTeamIds []string `json:"ApprovingTeamIds"`

	resources.Resource
}

// NewApprovalPolicy creates an ApprovalPolicy with server-friendly defaults
// (Tag scoping, 2 minimum approvers, empty scope/approver slices).
func NewApprovalPolicy(name string) *ApprovalPolicy {
	return &ApprovalPolicy{
		Name:                     name,
		ScopingStrategy:          ApprovalPolicyScopingStrategyTag,
		TagScopes:                []ApprovalPolicyTagScope{},
		IdScopes:                 []ApprovalPolicyIdScope{},
		MinimumApproversRequired: 2,
		ApprovingUserIds:         []string{},
		ApprovingTeamIds:         []string{},
		Resource:                 *resources.NewResource(),
	}
}

func (p *ApprovalPolicy) GetName() string     { return p.Name }
func (p *ApprovalPolicy) SetName(name string) { p.Name = name }
