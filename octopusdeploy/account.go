package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// account is the embedded struct used for all accounts.
type account struct {
	accountType            AccountType
	Description            string
	EnvironmentIDs         []string
	Name                   string `validate:"required,notblank,notall"`
	SpaceID                string `validate:"omitempty,notblank"`
	TenantedDeploymentMode TenantedDeploymentMode
	TenantIDs              []string
	TenantTags             []string

	resource
}

// newAccount creates and initializes an account.
func newAccount(name string, accountType AccountType) *account {
	return &account{
		accountType:            accountType,
		EnvironmentIDs:         []string{},
		Name:                   name,
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
		TenantIDs:              []string{},
		TenantTags:             []string{},
		resource:               *newResource(),
	}
}

// GetAccountType returns the type of this account.
func (a *account) GetAccountType() AccountType {
	return a.accountType
}

// GetDescription returns the description of the account.
func (a *account) GetDescription() string {
	return a.Description
}

// GetName returns the name of the account.
func (a *account) GetName() string {
	return a.Name
}

// GetSpaceID returns the space ID of this account.
func (a *account) GetSpaceID() string {
	return a.SpaceID
}

// SetDescription sets the description of the account.
func (a *account) SetDescription(description string) {
	a.Description = description
}

// SetName sets the name of the account.
func (a *account) SetName(name string) {
	a.Name = name
}

// SetSpaceID sets the space ID of this account.
func (a *account) SetSpaceID(spaceID string) {
	a.SpaceID = spaceID
}

// Validate checks the state of the account and returns an error if
// invalid.
func (a *account) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &account{}
