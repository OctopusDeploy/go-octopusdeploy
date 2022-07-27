package userroles

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

// UserRole represents a user role in Octopus.
type UserRole struct {
	CanBeDeleted                 bool     `json:"CanBeDeleted,omitempty"`
	Description                  string   `json:"Description,omitempty"`
	GrantedSpacePermissions      []string `json:"GrantedSpacePermissions"`
	GrantedSystemPermissions     []string `json:"GrantedSystemPermissions"`
	Name                         string   `json:"Name,omitempty"`
	SpacePermissionDescriptions  []string `json:"SpacePermissionDescriptions"`
	SupportedRestrictions        []string `json:"SupportedRestrictions"`
	SystemPermissionDescriptions []string `json:"SystemPermissionDescriptions"`

	resources.Resource
}

// NewUserRole initializes a user role with a name.
func NewUserRole(name string) *UserRole {
	return &UserRole{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}
