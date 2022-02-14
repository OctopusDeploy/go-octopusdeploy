package resources

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type ProjectTrigger struct {
	Action      ITriggerAction `json:"Action"`
	Description string         `json:"Description,omitempty"`
	Filter      ITriggerFilter `json:"Filter"`
	IsDisabled  bool           `json:"IsDisabled,omitempty"`
	Name        string         `json:"Name"`
	ProjectID   string         `json:"ProjectId,omitempty"`
	SpaceID     string         `json:"SpaceId,omitempty"`

	Resource
}

func NewProjectTrigger(name string, description string, isDisabled bool, projectID string, action ITriggerAction, filter ITriggerFilter) *ProjectTrigger {
	return &ProjectTrigger{
		Action:     action,
		Filter:     filter,
		IsDisabled: isDisabled,
		Name:       name,
		ProjectID:  projectID,
		Resource:   *NewResource(),
	}
}

// UnmarshalJSON sets this trigger to its representation in JSON.
func (projectTrigger *ProjectTrigger) UnmarshalJSON(b []byte) error {
	var rawMessage map[string]*json.RawMessage
	if err := json.Unmarshal(b, &rawMessage); err != nil {
		return err
	}

	var r Resource
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	projectTrigger.Resource = r

	for k, v := range rawMessage {
		switch k {
		case "Action":
			var action triggerAction
			if err := json.Unmarshal(*v, &action); err != nil {
				return err
			}
			switch action.Type {
			case AutoDeploy:
				var action *AutoDeployAction
				if err := json.Unmarshal(*v, &action); err != nil {
					return err
				}
				projectTrigger.Action = action
			case DeployLatestRelease:
				var action *DeployLatestReleaseAction
				if err := json.Unmarshal(*v, &action); err != nil {
					return err
				}
				projectTrigger.Action = action
			case DeployNewRelease:
				var action *DeployNewReleaseAction
				if err := json.Unmarshal(*v, &action); err != nil {
					return err
				}
				projectTrigger.Action = action
			case RunRunbook:
				var action *RunRunbookAction
				if err := json.Unmarshal(*v, &action); err != nil {
					return err
				}
				projectTrigger.Action = action
			}
		case "Description":
			if v != nil {
				if err := json.Unmarshal(*v, &projectTrigger.Description); err != nil {
					return err
				}
			}
		case "Filter":
			var filter triggerFilter
			if err := json.Unmarshal(*v, &filter); err != nil {
				return err
			}
			switch filter.Type {
			case ContinuousDailySchedule:
				var filter *ContinuousDailyScheduledTriggerFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			case CronExpressionSchedule:
				var filter *CronScheduledTriggerFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			case DailySchedule:
				var filter *DailyScheduledTriggerFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			case DaysPerMonthSchedule:
				var filter *monthlyScheduledTriggerFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			case DaysPerWeekSchedule:
				// TODO: sort this out
			case MachineFilter:
				var filter *DeploymentTargetFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			case OnceDailySchedule:
				var filter *OnceDailyScheduledTriggerFilter
				if err := json.Unmarshal(*v, &filter); err != nil {
					return err
				}
				projectTrigger.Filter = filter
			}
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
