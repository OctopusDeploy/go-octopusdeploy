package platformhubpolicies

import (
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// PolicyCandidate represents a set of information to create new Hub policy.
type PolicyCandidate struct {
	GitRef          string
	Slug            string
	Name            string
	Description     string
	ScopeRego       string
	ConditionsRego  string
	ViolationAction string
	ViolationReason string
}

// Policy represents a set of information to create new Hub policy.
type Policy interface {
	GetName() string
	SetName(string)

	GetDescription() string
	SetDescription(string)

	GetScopeRego() string
	SetScopeRego(string)

	GetConditionsRego() string
	SetConditionsRego(string)

	GetViolationAction() string
	SetViolationAction(string)

	GetViolationReason() string
	SetViolationReason(string)

	Validate() error

	PolicyKey
}

type PolicyKey interface {
	GetGitRef() string
	GetSlug() string
}

// PublishedPolicy represents a read-only view of a published Platform Hub policy version.
type PublishedPolicy interface {
	GetID() string
	GetPublishedDate() time.Time
	GetGitRef() string
	GetGitCommit() string
	GetName() string
	GetDescription() string
	GetViolationReason() string
	GetViolationAction() string
	GetScopeRego() string
	GetConditionsRego() string
	IsActivated() bool

	PublishedPolicyKey
}

type PublishedPolicyKey interface {
	GetSlug() string
	GetVersion() string
}

type persistedPolicy struct {
	GitRef          string `json:"GitRef" validate:"required,notblank"`
	Slug            string `json:"Slug" validate:"required,notblank"`
	Name            string `json:"Name" validate:"required,notblank"`
	Description     string `json:"Description,omitempty"`
	ScopeRego       string `json:"ScopeRego" validate:"required,notblank"`
	ConditionsRego  string `json:"ConditionsRego" validate:"required,notblank"`
	ViolationAction string `json:"ViolationAction" validate:"required,notblank"`
	ViolationReason string `json:"ViolationReason,omitempty"`
}

func (p *persistedPolicy) GetName() string             { return p.Name }
func (p *persistedPolicy) SetName(name string)         { p.Name = name }
func (p *persistedPolicy) GetGitRef() string           { return p.GitRef }
func (p *persistedPolicy) GetSlug() string             { return p.Slug }
func (p *persistedPolicy) GetDescription() string      { return p.Description }
func (p *persistedPolicy) SetDescription(d string)     { p.Description = d }
func (p *persistedPolicy) GetScopeRego() string        { return p.ScopeRego }
func (p *persistedPolicy) SetScopeRego(s string)       { p.ScopeRego = s }
func (p *persistedPolicy) GetConditionsRego() string   { return p.ConditionsRego }
func (p *persistedPolicy) SetConditionsRego(c string)  { p.ConditionsRego = c }
func (p *persistedPolicy) GetViolationReason() string  { return p.ViolationReason }
func (p *persistedPolicy) SetViolationReason(r string) { p.ViolationReason = r }
func (p *persistedPolicy) GetViolationAction() string  { return p.ViolationAction }
func (p *persistedPolicy) SetViolationAction(a string) { p.ViolationAction = a }

// Validate checks the state of the policy and returns an error if invalid.
func (p *persistedPolicy) Validate() error {
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

type publishedPolicyVersion struct {
	ID              string    `json:"Id"`
	Slug            string    `json:"Slug"`
	Version         string    `json:"Version"`
	PublishedDate   time.Time `json:"PublishedDate"`
	GitRef          string    `json:"GitRef"`
	GitCommit       string    `json:"GitCommit"`
	Name            string    `json:"Name"`
	Description     string    `json:"Description,omitempty"`
	ViolationReason string    `json:"ViolationReason,omitempty"`
	ViolationAction string    `json:"ViolationAction"`
	RegoScope       string    `json:"RegoScope"`
	RegoConditions  string    `json:"RegoConditions"`
	IsActive        bool      `json:"IsActive"`
}

func (v *publishedPolicyVersion) GetID() string               { return v.ID }
func (v *publishedPolicyVersion) GetSlug() string             { return v.Slug }
func (v *publishedPolicyVersion) GetVersion() string          { return v.Version }
func (v *publishedPolicyVersion) GetPublishedDate() time.Time { return v.PublishedDate }
func (v *publishedPolicyVersion) GetGitRef() string           { return v.GitRef }
func (v *publishedPolicyVersion) GetGitCommit() string        { return v.GitCommit }
func (v *publishedPolicyVersion) GetName() string             { return v.Name }
func (v *publishedPolicyVersion) GetDescription() string      { return v.Description }
func (v *publishedPolicyVersion) GetViolationReason() string  { return v.ViolationReason }
func (v *publishedPolicyVersion) GetViolationAction() string  { return v.ViolationAction }
func (v *publishedPolicyVersion) GetScopeRego() string        { return v.RegoScope }
func (v *publishedPolicyVersion) GetConditionsRego() string   { return v.RegoConditions }
func (v *publishedPolicyVersion) IsActivated() bool           { return v.IsActive }
