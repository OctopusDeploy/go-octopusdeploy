package octopusdeploy

type ProjectTriggers struct {
	Items []*ProjectTrigger `json:"Items"`
	PagedResults
}

type ProjectTrigger struct {
	Action     ProjectTriggerAction `json:"Action"`
	Filter     ProjectTriggerFilter `json:"Filter"`
	IsDisabled bool                 `json:"IsDisabled,omitempty"`
	Name       string               `json:"Name"`
	ProjectID  string               `json:"ProjectId,omitempty"`

	resource
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
		resource:  *newResource(),
	}
}
