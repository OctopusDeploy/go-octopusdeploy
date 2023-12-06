package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// IAccount defines the interface for accounts.
type IAccount interface {
	GetAccountType() AccountType
	GetDescription() string
	GetEnvironmentIDs() []string
	GetSlug() string
	GetTenantedDeploymentMode() core.TenantedDeploymentMode
	GetTenantIDs() []string
	GetTenantTags() []string
	SetDescription(string)
	SetEnvironmentIDs([]string)
	SetSlug(string)
	SetTenantedDeploymentMode(core.TenantedDeploymentMode)
	SetTenantIDs([]string)
	SetTenantTags([]string)

	resources.IHasName
	resources.IHasSpace
	resources.IResource
}

// account is the embedded struct used for all accounts.
type account struct {
	AccountType            AccountType                 `json:"AccountType" validate:"required,oneof=None UsernamePassword SshKeyPair AzureSubscription AzureServicePrincipal AzureOidc AmazonWebServicesAccount AwsOidc AmazonWebServicesRoleAccount GoogleCloudAccount Token"`
	Description            string                      `json:"Description,omitempty"`
	EnvironmentIDs         []string                    `json:"EnvironmentIds,omitempty"`
	Name                   string                      `json:"Name" validate:"required,notblank,notall"`
	Slug                   string                      `json:"Slug,omitempty"`
	SpaceID                string                      `json:"SpaceId,omitempty"`
	TenantedDeploymentMode core.TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIDs              []string                    `json:"TenantIds,omitempty"`
	TenantTags             []string                    `json:"TenantTags,omitempty"`

	resources.Resource
}

// newAccount creates and initializes an account.
func newAccount(name string, accountType AccountType) *account {
	return &account{
		AccountType:            accountType,
		EnvironmentIDs:         []string{},
		Name:                   name,
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
		TenantIDs:              []string{},
		TenantTags:             []string{},
		Resource:               *resources.NewResource(),
	}
}

// GetAccountType returns the type of this account.
func (a *account) GetAccountType() AccountType {
	return a.AccountType
}

// GetDescription returns the description of the account.
func (a *account) GetDescription() string {
	return a.Description
}

// GetEnvironmentIDs returns the environment IDs associated with this account.
func (a *account) GetEnvironmentIDs() []string {
	return a.EnvironmentIDs
}

// GetName returns the name of the account.
func (a *account) GetName() string {
	return a.Name
}

// GetSlug returns the slug to this account.
func (a *account) GetSlug() string {
	return a.Slug
}

// GetSpaceID returns the space ID of this account.
func (a *account) GetSpaceID() string {
	return a.SpaceID
}

// GetTenantedDeploymentMode returns the tenanted deployment mode of this account.
func (a *account) GetTenantedDeploymentMode() core.TenantedDeploymentMode {
	return a.TenantedDeploymentMode
}

// GetTenantIDs returns the tenant IDs associated with this account.
func (a *account) GetTenantIDs() []string {
	return a.TenantIDs
}

// GetTenantTags returns the tenant tags assigned to this account.
func (a *account) GetTenantTags() []string {
	return a.TenantTags
}

// SetDescription sets the description of the account.
func (a *account) SetDescription(description string) {
	a.Description = description
}

// SetEnvironmentIDs sets the associated environment IDs of the account.
func (a *account) SetEnvironmentIDs(environmentIds []string) {
	a.EnvironmentIDs = environmentIds
}

// SetName sets the name of the account.
func (a *account) SetName(name string) {
	a.Name = name
}

// SetSlug sets the slug of this account.
func (a *account) SetSlug(slug string) {
	a.Slug = slug
}

// SetSpaceID sets the space ID of this account.
func (a *account) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
}

// SetTenantedDeploymentMode sets the tenanted deployment mode of this account.
func (a *account) SetTenantedDeploymentMode(mode core.TenantedDeploymentMode) {
	a.TenantedDeploymentMode = mode
}

// SetTenantIDs sets the tenant IDs associated with this account.
func (a *account) SetTenantIDs(tenantIds []string) {
	a.TenantIDs = tenantIds
}

// SetTenantTags sets the tenant tags associated with this account.
func (a *account) SetTenantTags(tenantTags []string) {
	a.TenantTags = tenantTags
}

// Validate checks the state of the account and returns an error if
// invalid.
func (a *account) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &account{}
