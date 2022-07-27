package channels

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/validation"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type Channel struct {
	Description string        `json:"Description,omitempty"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId,omitempty"`
	Name        string        `json:"Name" validate:"required,notblank,notall"`
	ProjectID   string        `json:"ProjectId" validate:"required,notblank"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	SpaceID     string        `json:"SpaceId,omitempty"`
	TenantTags  []string      `json:"TenantTags,omitempty"`

	resources.Resource
}

func NewChannel(name string, projectID string) *Channel {
	return &Channel{
		Name:       strings.TrimSpace(name),
		ProjectID:  projectID,
		Rules:      []ChannelRule{},
		TenantTags: []string{},
		Resource:   *resources.NewResource(),
	}
}

// Validate checks the state of the channel and returns an error if invalid.
func (c Channel) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", validation.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(c)
}
