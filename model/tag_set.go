package model

import (
	"github.com/go-playground/validator/v10"
)

type TagSets struct {
	Items []TagSet `json:"Items"`
	PagedResults
}

type TagSet struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Tags []Tag  `json:"Tags,omitempty"`
}

func (t *TagSet) GetID() string {
	return t.ID
}

func (t *TagSet) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name: name,
	}
}

var _ ResourceInterface = &TagSet{}
