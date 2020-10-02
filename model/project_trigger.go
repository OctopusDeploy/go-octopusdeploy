package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ProjectTriggers struct {
	Items []ProjectTrigger `json:"Items"`
	PagedResults
}

type ProjectTrigger struct {
	Action     ProjectTriggerAction `json:"Action"`
	Filter     ProjectTriggerFilter `json:"Filter"`
	IsDisabled bool                 `json:"IsDisabled,omitempty"`
	Name       string               `json:"Name"`
	ProjectID  string               `json:"ProjectId,omitempty"`

	Resource
}

func (t *ProjectTrigger) AddEventGroups(eventGroups []string) {
	t.Filter.EventGroups = append(t.Filter.EventGroups, eventGroups...)
}

func (t *ProjectTrigger) AddEventCategories(eventCategories []string) {
	t.Filter.EventCategories = append(t.Filter.EventCategories, eventCategories...)
}

func NewProjectDeploymentTargetTrigger(name, projectID string, shouldRedeploy bool, roles, eventGroups, eventCategories []string) *ProjectTrigger {
	return &ProjectTrigger{
		Action: ProjectTriggerAction{
			ActionType: "AutoDeploy",
			ShouldRedeployWhenMachineHasBeenDeployedTo: shouldRedeploy,
		},
		Filter: ProjectTriggerFilter{
			EventCategories: eventCategories,
			EventGroups:     eventGroups,
			FilterType:      "MachineFilter",
			Roles:           roles,
		},
		Name:      name,
		ProjectID: projectID,
	}
}

// GetID returns the ID value of the ProjectTrigger.
func (resource ProjectTrigger) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this ProjectTrigger.
func (resource ProjectTrigger) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this ProjectTrigger was changed.
func (resource ProjectTrigger) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this ProjectTrigger.
func (resource ProjectTrigger) GetLinks() map[string]string {
	return resource.Links
}

func (resource ProjectTrigger) SetID(id string) {
	resource.ID = id
}

func (resource ProjectTrigger) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource ProjectTrigger) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the ProjectTrigger and returns an error if invalid.
func (resource ProjectTrigger) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &ProjectTrigger{}
