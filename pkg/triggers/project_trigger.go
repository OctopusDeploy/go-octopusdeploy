package triggers

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/filters"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type ProjectTrigger struct {
	Action      actions.ITriggerAction `json:"Action"`
	Description string                 `json:"Description,omitempty"`
	Filter      filters.ITriggerFilter `json:"Filter"`
	IsDisabled  bool                   `json:"IsDisabled"`
	Name        string                 `json:"Name"`
	ProjectID   string                 `json:"ProjectId,omitempty"`
	SpaceID     string                 `json:"SpaceId,omitempty"`

	resources.Resource
}

func NewProjectTrigger(name string, description string, isDisabled bool, projectID string, action actions.ITriggerAction, filter filters.ITriggerFilter) *ProjectTrigger {
	return &ProjectTrigger{
		Action:     action,
		Filter:     filter,
		IsDisabled: isDisabled,
		Name:       name,
		ProjectID:  projectID,
		Resource:   *resources.NewResource(),
	}
}

// UnmarshalJSON sets this trigger to its representation in JSON.
func (projectTrigger *ProjectTrigger) UnmarshalJSON(b []byte) error {
	var rawMessage map[string]*json.RawMessage
	if err := json.Unmarshal(b, &rawMessage); err != nil {
		return err
	}

	var r resources.Resource
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	projectTrigger.Resource = r

	for k, v := range rawMessage {
		switch k {
		case "Action":
			action, err := actions.FromJson(v)
			if err != nil {
				return err
			}
			projectTrigger.Action = action
		case "Description":
			if v != nil {
				if err := json.Unmarshal(*v, &projectTrigger.Description); err != nil {
					return err
				}
			}
		case "Filter":
			filter, err := filters.FromJson(v)
			if err != nil {
				return err
			}
			projectTrigger.Filter = filter
		case "IsDisabled":
			if err := json.Unmarshal(*v, &projectTrigger.IsDisabled); err != nil {
				return err
			}
		case "Name":
			if err := json.Unmarshal(*v, &projectTrigger.Name); err != nil {
				return err
			}
		case "ProjectId":
			if err := json.Unmarshal(*v, &projectTrigger.ProjectID); err != nil {
				return err
			}
		case "SpaceId":
			if err := json.Unmarshal(*v, &projectTrigger.SpaceID); err != nil {
				return err
			}
		}
	}

	return nil
}

// Validate checks the state of the deployment target and returns an error if
// invalid.
func (t *ProjectTrigger) Validate() error {
	return validator.New().Struct(t)
}
