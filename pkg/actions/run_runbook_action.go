package actions

type RunRunbookAction struct {
	Environments []string `json:"EnvironmentIds"`
	Runbook      string   `json:"RunbookId"`
	Tenants      []string `json:"TenantIds"`
	TenantTags   []string `json:"TenantTags"`

	triggerAction
}

func NewRunRunbookAction() *RunRunbookAction {
	return &RunRunbookAction{
		Environments:  []string{},
		Tenants:       []string{},
		TenantTags:    []string{},
		triggerAction: *newTriggerAction(RunRunbook),
	}
}

func (a *RunRunbookAction) GetActionType() ActionType {
	return a.Type
}

func (a *RunRunbookAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &RunRunbookAction{}
