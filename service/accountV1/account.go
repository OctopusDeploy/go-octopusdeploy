package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// Account is the embedded struct used for all accounts.
type Account struct {
	AccountType            AccountType `validate:"required"`
	Description            string
	EnvironmentIDs         []string
	Name                   string `validate:"required,notblank,notall"`
	SpaceID                string `validate:"omitempty,notblank"`
	TenantedDeploymentMode resources.TenantedDeploymentMode
	TenantIDs              []string
	TenantTags             []string

	resources.Resource
}

// IAccount defines the interface for accounts.
type IAccount interface {
	GetAccountType() AccountType
	GetDescription() string
	GetEnvironmentIDs() []string
	GetTenantedDeploymentMode() resources.TenantedDeploymentMode
	GetTenantIDs() []string
	GetTenantTags() []string
	SetDescription(string)
	SetEnvironmentIDs([]string)
	SetTenantedDeploymentMode(resources.TenantedDeploymentMode)
	SetTenantIDs([]string)
	SetTenantTags([]string)

	resources.IHasName
	resources.IHasSpace
	resources.IResource
}

// NewAccount creates and initializes an accountV1.
func NewAccount(name string, accountType AccountType) *Account {
	return &Account{
		AccountType:            accountType,
		EnvironmentIDs:         []string{},
		Name:                   name,
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
		TenantIDs:              []string{},
		TenantTags:             []string{},
		Resource:               *resources.NewResource(),
	}
}

// GetAccountType returns the type of this accountV1.
func (a *Account) GetAccountType() AccountType {
	return a.AccountType
}

// GetDescription returns the description of the accountV1.
func (a *Account) GetDescription() string {
	return a.Description
}

// GetEnvironmentIDs returns the environment IDs associated with this accountV1.
func (a *Account) GetEnvironmentIDs() []string {
	return a.EnvironmentIDs
}

// GetName returns the name of the accountV1.
func (a *Account) GetName() string {
	return a.Name
}

// GetSpaceID returns the space ID of this accountV1.
func (a *Account) GetSpaceID() string {
	return a.SpaceID
}

// GetTenantedDeploymentMode returns the tenanted deployment mode of this accountV1.
func (a *Account) GetTenantedDeploymentMode() resources.TenantedDeploymentMode {
	return a.TenantedDeploymentMode
}

// GetTenantIDs returns the tenant IDs associated with this accountV1.
func (a *Account) GetTenantIDs() []string {
	return a.TenantIDs
}

// GetTenantTags returns the tenant tags assigned to this accountV1.
func (a *Account) GetTenantTags() []string {
	return a.TenantTags
}

// SetDescription sets the description of the accountV1.
func (a *Account) SetDescription(description string) {
	a.Description = description
}

// SetEnvironmentIDs sets the associated environment IDs of the accountV1.
func (a *Account) SetEnvironmentIDs(environmentIds []string) {
	a.EnvironmentIDs = environmentIds
}

// SetName sets the name of the accountV1.
func (a *Account) SetName(name string) {
	a.Name = name
}

// SetSpaceID sets the space ID of this accountV1.
func (a *Account) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
}

// SetTenantedDeploymentMode sets the tenanted deployment mode of this accountV1.
func (a *Account) SetTenantedDeploymentMode(mode resources.TenantedDeploymentMode) {
	a.TenantedDeploymentMode = mode
}

// SetTenantIDs sets the tenant IDs associated with this accountV1.
func (a *Account) SetTenantIDs(tenantIds []string) {
	a.TenantIDs = tenantIds
}

// SetTenantTags sets the tenant tags associated with this accountV1.
func (a *Account) SetTenantTags(tenantTags []string) {
	a.TenantTags = tenantTags
}

// Validate checks the state of the accountV1 and returns an error if
// invalid.
func (a *Account) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &Account{}
