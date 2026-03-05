package platformhubpolicies

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PlatformHubPolicy represents a Platform Hub policy resource.
type PlatformHubPolicy struct {
	Name            string `json:"Name" validate:"required,notblank"`
	GitRef          string `json:"GitRef" validate:"required,notblank"`
	Slug            string `json:"Slug" validate:"required,notblank"`
	Description     string `json:"Description,omitempty"`
	ScopeRego       string `json:"ScopeRego" validate:"required,notblank"`
	ConditionsRego  string `json:"ConditionsRego" validate:"required,notblank"`
	ViolationReason string `json:"ViolationReason,omitempty"`
	ViolationAction string `json:"ViolationAction" validate:"required,notblank"`
}

// NewPlatformHubPolicy creates and initializes a Platform Hub policy.
func NewPlatformHubPolicy(name, gitRef, slug, scopeRego, conditionsRego, violationAction string) (*PlatformHubPolicy, error) {
	policy := PlatformHubPolicy{
		Name:            name,
		GitRef:          gitRef,
		Slug:            slug,
		ScopeRego:       scopeRego,
		ConditionsRego:  conditionsRego,
		ViolationAction: violationAction,
	}

	if validationError := policy.Validate(); validationError != nil {
		return nil, validationError
	}

	return &policy, nil
}

// GetName returns the name of the policy.
func (p *PlatformHubPolicy) GetName() string {
	return p.Name
}

// SetName sets the name of the policy.
func (p *PlatformHubPolicy) SetName(name string) {
	p.Name = name
}

// GetGitRef returns the git ref of the policy.
func (p *PlatformHubPolicy) GetGitRef() string {
	return p.GitRef
}

// GetSlug returns the slug of the policy.
func (p *PlatformHubPolicy) GetSlug() string {
	return p.Slug
}

// GetDescription returns the description of the policy.
func (p *PlatformHubPolicy) GetDescription() string {
	return p.Description
}

// SetDescription sets the description of the policy.
func (p *PlatformHubPolicy) SetDescription(description string) {
	p.Description = description
}

// GetScopeRego returns the scope Rego of the policy.
func (p *PlatformHubPolicy) GetScopeRego() string {
	return p.ScopeRego
}

// SetScopeRego sets the scope Rego of the policy.
func (p *PlatformHubPolicy) SetScopeRego(scopeRego string) {
	p.ScopeRego = scopeRego
}

// GetConditionsRego returns the conditions rego of the policy.
func (p *PlatformHubPolicy) GetConditionsRego() string {
	return p.ConditionsRego
}

// SetConditionsRego sets the conditions rego of the policy.
func (p *PlatformHubPolicy) SetConditionsRego(conditionsRego string) {
	p.ConditionsRego = conditionsRego
}

// GetViolationReason returns the violation reason of the policy.
func (p *PlatformHubPolicy) GetViolationReason() string {
	return p.ViolationReason
}

// SetViolationReason sets the violation reason of the policy.
func (p *PlatformHubPolicy) SetViolationReason(violationReason string) {
	p.ViolationReason = violationReason
}

// GetViolationAction returns the violation action of the policy.
func (p *PlatformHubPolicy) GetViolationAction() string {
	return p.ViolationAction
}

// SetViolationAction sets the violation action of the policy.
func (p *PlatformHubPolicy) SetViolationAction(violationAction string) {
	p.ViolationAction = violationAction
}

// Validate checks the state of the policy and returns an error if invalid.
func (p *PlatformHubPolicy) Validate() error {
	validate, err := getValidator()
	if err != nil {
		return err
	}

	return validate.Struct(p)
}

var getValidator = sync.OnceValues(buildValidator)

func buildValidator() (*validator.Validate, error) {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return nil, err
	}

	return v, nil
}
