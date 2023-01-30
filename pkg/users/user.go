package users

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

// User represents a user in Octopus.
type User struct {
	CanPasswordBeEdited bool       `json:"CanPasswordBeEdited,omitempty"`
	DisplayName         string     `json:"DisplayName,omitempty"`
	EmailAddress        string     `json:"EmailAddress,omitempty"`
	Identities          []Identity `json:"Identities,omitempty"`
	IsActive            bool       `json:"IsActive"`
	IsRequestor         bool       `json:"IsRequestor,omitempty"`
	IsService           bool       `json:"IsService,omitempty"`
	Password            string     `json:"Password,omitempty" validate:"max=20"`
	Username            string     `json:"Username,omitempty"`

	resources.Resource
}

// NewUser initializes a user with an username and a display name.
func NewUser(username string, displayName string) *User {
	return &User{
		DisplayName: displayName,
		Username:    username,
		Resource:    *resources.NewResource(),
		IsActive:    true,
	}
}
