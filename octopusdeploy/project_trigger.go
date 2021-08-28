package octopusdeploy

type ProjectTriggers struct {
	Items []*ProjectTrigger `json:"Items"`
	PagedResults
}

type ProjectTrigger struct {
	Action      *TriggerAction `json:"Action" validate:"required"`
	Description string         `json:"Description,omitempty"`
	Filter      *TriggerFilter `json:"Filter" validate:"required"`
	IsDisabled  bool           `json:"IsDisabled,omitempty"`
	Name        string         `json:"Name" validate:"required"`
	ProjectID   string         `json:"ProjectId" validate:"required"`
	SpaceID     string         `json:"SpaceId,omitempty"`

	resource
}

func (t *ProjectTrigger) AddEventGroups(eventGroups []string) {
	t.Filter.EventGroups = append(t.Filter.EventGroups, eventGroups...)
}

func (t *ProjectTrigger) AddEventCategories(eventCategories []string) {
	t.Filter.EventCategories = append(t.Filter.EventCategories, eventCategories...)
}

func NewProjectDeploymentTargetTrigger(name, projectID string, shouldRedeploy bool, roles, eventGroups, eventCategories []string) *ProjectTrigger {
	projectTrigger := NewProjectTrigger()
	projectTrigger.Name = name
	projectTrigger.ProjectID = projectID
	projectTrigger.Action = &TriggerAction{
		ActionType: "AutoDeploy",
		ShouldRedeployWhenMachineHasBeenDeployedTo: shouldRedeploy,
	}
	projectTrigger.Filter = &TriggerFilter{
		EventCategories: eventCategories,
		EventGroups:     eventGroups,
		FilterType:      "MachineFilter",
		Roles:           roles,
	}
	return projectTrigger
}

func NewProjectTrigger() *ProjectTrigger {
	return &ProjectTrigger{
		resource: *newResource(),
	}
}
