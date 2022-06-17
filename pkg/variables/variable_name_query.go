package variables

type VariableNamesQuery struct {
	Project                   string `uri:"project,omitempty" url:"project,omitempty"`
	ProjectEnvironmentsFilter string `uri:"projectEnvironmentsFilter,omitempty" url:"projectEnvironmentsFilter,omitempty"`
	Runbook                   string `uri:"runbook,omitempty" url:"runbook,omitempty"`
}
