package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AccountResource is the embedded struct used for all accounts.
type AccountResource struct {
	AccountType            string   `json:"AccountType" validate:"required,oneof=None UsernamePassword SshKeyPair AzureSubscription AzureServicePrincipal AmazonWebServicesAccount AmazonWebServicesRoleAccount Token"`
	Description            string   `json:"Description,omitempty"`
	EnvironmentIDs         []string `json:"EnvironmentIds,omitempty"`
	Name                   string   `json:"Name" validate:"required,notblank,notall"`
	SpaceID                string   `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	TenantedDeploymentMode string   `json:"TenantedDeploymentParticipation" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs              []string `json:"TenantIds,omitempty"`
	TenantTags             []string `json:"TenantTags,omitempty"`

	resource
}

// newAccountResource creates and initializes an account resource.
func newAccountResource(name string, accountType string) *AccountResource {
	return &AccountResource{
		AccountType:            accountType,
		EnvironmentIDs:         []string{},
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
		TenantIDs:              []string{},
		TenantTags:             []string{},
		resource:               *newResource(),
	}
}

// GetAccountType returns the type of this account resource.
func (a *AccountResource) GetAccountType() string {
	return a.AccountType
}

// GetDescription returns the description of the account resource.
func (a *AccountResource) GetDescription() string {
	return a.Description
}

// GetName returns the name of the account resource.
func (a *AccountResource) GetName() string {
	return a.Name
}

// SetDescription sets the description of the account resource.
func (a *AccountResource) SetDescription(description string) {
	a.Description = description
}

// SetName sets the name of the account resource.
func (a *AccountResource) SetName(name string) {
	a.Name = name
}

// Validate checks the state of the account resource and returns an error if
// invalid.
func (a *AccountResource) Validate() error {
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

var _ IAccount = &AccountResource{}
