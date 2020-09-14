package model

import (
	"github.com/go-playground/validator/v10"
)

type ProjectGroups struct {
	Items []ProjectGroup `json:"Items"`
	PagedResults
}

type ProjectGroup struct {
	Description       string   `json:"Description,omitempty"`
	EnvironmentIds    []string `json:"EnvironmentIds"`
	Name              string   `json:"Name,omitempty" validate:"required"`
	RetentionPolicyID string   `json:"RetentionPolicyId,omitempty"`

	Resource
}

func (p *ProjectGroup) GetID() string {
	return p.ID
}

func (p *ProjectGroup) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)

	if err != nil {
		return err
	}

	return nil
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name: name,
	}
}

var _ ResourceInterface = &Project{}
