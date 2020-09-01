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
	DisplayName         string     `json:"DisplayName"`
	IsActive            bool       `json:"IsActive"`
	IsService           bool       `json:"IsService"`
	EmailAddress        string     `json:"EmailAddress,omitempty"`
	CanPasswordBeEdited bool       `json:"CanPasswordBeEdited"`
	IsRequestor         bool       `json:"IsRequestor"`
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
