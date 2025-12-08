package platformhubaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	validation "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// IPlatformHubAccount defines the interface for Platform Hub accounts.
type IPlatformHubAccount interface {
	GetAccountType() PlatformHubAccountType
	GetDescription() string
	SetDescription(string)

	resources.IHasName
	resources.IResource
}

// platformHubAccount is the embedded struct used for all Platform Hub accounts.
type platformHubAccount struct {
	AccountType PlatformHubAccountType `json:"AccountType" validate:"required"`
	Description string                 `json:"Description,omitempty"`
	Name        string                 `json:"Name" validate:"required,notblank"`

	resources.Resource
}

// newPlatformHubAccount creates and initializes a Platform Hub account.
func newPlatformHubAccount(name string, accountType PlatformHubAccountType) *platformHubAccount {
	return &platformHubAccount{
		AccountType: accountType,
		Name:        name,
		Resource:    *resources.NewResource(),
	}
}

// GetAccountType returns the type of this account.
func (a *platformHubAccount) GetAccountType() PlatformHubAccountType {
	return a.AccountType
}

// GetDescription returns the description of the account.
func (a *platformHubAccount) GetDescription() string {
	return a.Description
}

// GetName returns the name of the account.
func (a *platformHubAccount) GetName() string {
	return a.Name
}

// SetDescription sets the description of the account.
func (a *platformHubAccount) SetDescription(description string) {
	a.Description = description
}

// SetName sets the name of the account.
func (a *platformHubAccount) SetName(name string) {
	a.Name = name
}

// Validate checks the state of the account and returns an error if invalid.
func (a *platformHubAccount) Validate() error {
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

var _ IPlatformHubAccount = &platformHubAccount{}
