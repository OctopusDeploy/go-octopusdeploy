package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// PlatformHubAccountResource represents details for all Platform Hub accounts.
type PlatformHubAccountResource struct {
	AccountType PlatformHubAccountType `json:"AccountType" validate:"required,oneof=AmazonWebServicesAccount AmazonWebServicesOidcAccount"`
	Description string                 `json:"Description,omitempty"`
	Name        string                 `json:"Name" validate:"required,notall"`
	AccessKey   string                 `json:"AccessKey,omitempty"`
	SecretKey   *core.SensitiveValue   `json:"SecretKey,omitempty"`

	// OIDC-specific fields
	RoleArn                string   `json:"RoleArn,omitempty"`
	SessionDuration        string   `json:"SessionDuration,omitempty"`
	DeploymentSubjectKeys  []string `json:"DeploymentSubjectKeys,omitempty"`
	HealthCheckSubjectKeys []string `json:"HealthCheckSubjectKeys,omitempty"`
	AccountTestSubjectKeys []string `json:"AccountTestSubjectKeys,omitempty"`

	resources.Resource
}

// NewPlatformHubAccountResource creates and initializes a Platform Hub account resource.
func NewPlatformHubAccountResource(name string, accountType PlatformHubAccountType) *PlatformHubAccountResource {
	return &PlatformHubAccountResource{
		AccountType: accountType,
		Name:        name,
		Resource:    *resources.NewResource(),
	}
}

// GetAccountType returns the type of this account resource.
func (r *PlatformHubAccountResource) GetAccountType() PlatformHubAccountType {
	return r.AccountType
}

// GetDescription returns the description of this account resource.
func (r *PlatformHubAccountResource) GetDescription() string {
	return r.Description
}

// GetName returns the name of this account resource.
func (r *PlatformHubAccountResource) GetName() string {
	return r.Name
}

// SetDescription sets the description of the account resource.
func (r *PlatformHubAccountResource) SetDescription(description string) {
	r.Description = description
}

// SetName sets the name of this account resource.
func (r *PlatformHubAccountResource) SetName(name string) {
	r.Name = name
}

// ToPlatformHubAccount converts a PlatformHubAccountResource to the appropriate concrete account type.
func (r *PlatformHubAccountResource) ToPlatformHubAccount() (IPlatformHubAccount, error) {
	if r == nil {
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccount", "PlatformHubAccountResource")
	}

	if err := r.Validate(); err != nil {
		return nil, err
	}

	var account IPlatformHubAccount

	switch r.AccountType {
	case AccountTypePlatformHubAwsAccount:
		awsAccount, err := NewPlatformHubAwsAccount(r.GetName(), r.AccessKey, r.SecretKey)
		if err != nil {
			return nil, err
		}
		account = awsAccount
	case AccountTypePlatformHubAwsOIDCAccount:
		oidcAccount, err := NewPlatformHubAwsOIDCAccount(r.GetName(), r.RoleArn)
		if err != nil {
			return nil, err
		}
		oidcAccount.SessionDuration = r.SessionDuration
		oidcAccount.DeploymentSubjectKeys = r.DeploymentSubjectKeys
		oidcAccount.HealthCheckSubjectKeys = r.HealthCheckSubjectKeys
		oidcAccount.AccountTestSubjectKeys = r.AccountTestSubjectKeys
		account = oidcAccount
	default:
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccount", "AccountType")
	}

	account.SetDescription(r.Description)
	account.SetID(r.GetID())
	account.SetLinks(r.GetLinks())
	account.SetModifiedBy(r.GetModifiedBy())
	account.SetModifiedOn(r.GetModifiedOn())

	return account, nil
}

// ToPlatformHubAccountResource converts a concrete account type to PlatformHubAccountResource.
func ToPlatformHubAccountResource(account IPlatformHubAccount) (*PlatformHubAccountResource, error) {
	if account == nil {
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccountResource", "PlatformHubAccount")
	}

	// conversion unnecessary if input account is *PlatformHubAccountResource
	if v, ok := account.(*PlatformHubAccountResource); ok {
		return v, nil
	}

	resource := NewPlatformHubAccountResource(account.GetName(), account.GetAccountType())

	switch resource.AccountType {
	case AccountTypePlatformHubAwsAccount:
		awsAccount := account.(*PlatformHubAwsAccount)
		resource.AccessKey = awsAccount.AccessKey
		resource.SecretKey = awsAccount.SecretKey
	case AccountTypePlatformHubAwsOIDCAccount:
		oidcAccount := account.(*PlatformHubAwsOIDCAccount)
		resource.RoleArn = oidcAccount.RoleArn
		resource.SessionDuration = oidcAccount.SessionDuration
		resource.DeploymentSubjectKeys = oidcAccount.DeploymentSubjectKeys
		resource.HealthCheckSubjectKeys = oidcAccount.HealthCheckSubjectKeys
		resource.AccountTestSubjectKeys = oidcAccount.AccountTestSubjectKeys
	default:
		return nil, internal.CreateInvalidParameterError("ToPlatformHubAccountResource", "AccountType")
	}

	resource.SetDescription(account.GetDescription())
	resource.SetID(account.GetID())
	resource.SetLinks(account.GetLinks())
	resource.SetModifiedBy(account.GetModifiedBy())
	resource.SetModifiedOn(account.GetModifiedOn())

	return resource, nil
}

var _ IPlatformHubAccount = &PlatformHubAccountResource{}
