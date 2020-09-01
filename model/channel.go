package model

import (
	"github.com/go-playground/validator/v10"
)

type Channels struct {
	Items []Channel `json:"Items"`
	PagedResults
}

type Channel struct {
	Description string        `json:"Description"`
	ID          string        `json:"Id,omitempty"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId"`
	Name        string        `json:"Name"`
	ProjectID   string        `json:"ProjectId"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	TenantTags  []string      `json:"TenantedDeploymentMode,omitempty"`
}

func (d *Channels) Validate() error {
	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		return err
	}

	return nil
}

// ValidateChannelValues checks the values of a Channel object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating channels.
func ValidateChannelValues(Channel *Channel) error {
	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", Channel.Name),
		ValidateRequiredPropertyValue("ProjectID", Channel.ProjectID),
	})
}

func NewChannel(name, description, projectID string) *Channel {
	return &Channel{
		Name:        name,
		ProjectID:   projectID,
		Description: description,
	}
}
