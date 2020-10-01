package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Users defines a collection of users with built-in support for paged results.
type Users struct {
	Items []User `json:"Items"`
	PagedResults
}

// User represents a user in Octopus.
type User struct {
	CanPasswordBeEdited bool       `json:"CanPasswordBeEdited,omitempty"`
	DisplayName         string     `json:"DisplayName,omitempty"`
	EmailAddress        string     `json:"EmailAddress,omitempty"`
	Identities          []Identity `json:"Identities,omitempty"`
	IsActive            bool       `json:"IsActive,omitempty"`
	IsRequestor         bool       `json:"IsRequestor,omitempty"`
	IsService           bool       `json:"IsService,omitempty"`
	Password            string     `json:"Password,omitempty"`
	Username            string     `json:"Username,omitempty"`

	Resource
}

// NewUser initializes a user with an username and a display name.
func NewUser(username string, displayName string) *User {
	return &User{
		Username:    username,
		DisplayName: displayName,
	}
}

// GetID returns the ID value of the User.
func (resource User) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of
// this user.
func (resource User) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this user was changed.
func (resource User) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this user.
func (resource User) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the user and returns an error if invalid.
func (resource User) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &User{}
