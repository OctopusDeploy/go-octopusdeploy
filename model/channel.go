package model

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Channels defines a collection of channels with built-in support for paged
// results.
type Channels struct {
	Items []*Channel `json:"Items"`
	PagedResults
}

type Channel struct {
	Description string        `json:"Description"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId"`
	Name        string        `json:"Name" validate:"required"`
	ProjectID   string        `json:"ProjectId" validate:"required"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	TenantTags  []string      `json:"TenantTags,omitempty"`

	Resource
}

func NewChannel(name string, description string, projectID string) *Channel {
	return &Channel{
		Description: strings.TrimSpace(description),
		Name:        strings.TrimSpace(name),
		ProjectID:   projectID,
		Rules:       []ChannelRule{},
		TenantTags:  []string{},
		Resource:    *newResource(),
	}
}

// Validate checks the state of the channel and returns an error if invalid.
func (c Channel) Validate() error {
	return validator.New().Struct(c)
}
