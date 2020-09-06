package model

import (
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Items []User `json:"Items"`
	PagedResults
}

type User struct {
	Username            string     `json:"Username,omitempty"`
	DisplayName         string     `json:"DisplayName,omitempty"`
	IsActive            bool       `json:"IsActive,omitempty"`
	IsService           bool       `json:"IsService,omitempty"`
	EmailAddress        string     `json:"EmailAddress,omitempty"`
	CanPasswordBeEdited bool       `json:"CanPasswordBeEdited,omitempty"`
	IsRequestor         bool       `json:"IsRequestor,omitempty"`
	Password            string     `json:"Password,omitempty"`
	Identities          []Identity `json:"Identities,omitempty"`
	Resource
}

func (t *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewUser(username, displayName string) *User {
	return &User{
		Username:    username,
		DisplayName: displayName,
	}
}
