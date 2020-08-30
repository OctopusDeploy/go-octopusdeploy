package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/go-playground/validator"
)

type UserService struct {
	sling *sling.Sling
}

func NewUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling,
	}
}

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

func (s *UserService) Get(userid string) (*User, error) {
	path := fmt.Sprintf("Users/%s", userid)
	resp, err := apiGet(s.sling, new(User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

func (s *UserService) GetAll() (*[]User, error) {
	var p []User

	path := "users"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Users), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Users)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *UserService) GetByName(username string) (*User, error) {
	var foundUser User
	Users, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, user := range *Users {
		if user.Username == username {
			return &user, nil
		}
	}

	return &foundUser, fmt.Errorf("no User found with User name %s", username)
}

func (s *UserService) Add(user *User) (*User, error) {
	resp, err := apiAdd(s.sling, user, new(User), "users")

	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

func (s *UserService) Delete(userid string) error {
	path := fmt.Sprintf("Users/%s", userid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(user *User) (*User, error) {
	path := fmt.Sprintf("Users/%s", user.ID)
	resp, err := apiUpdate(s.sling, user, new(User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}
