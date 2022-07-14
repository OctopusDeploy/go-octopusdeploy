package variables

type VariablePreviewQuery struct {
	Action      string `uri:"action,omitempty" url:"action,omitempty"`
	Channel     string `uri:"channel,omitempty" url:"channel,omitempty"`
	Environment string `uri:"environment,omitempty" url:"environment,omitempty"`
	Machine     string `uri:"machine,omitempty" url:"machine,omitempty"`
	Project     string `uri:"project,omitempty" url:"project,omitempty"`
	Role        string `uri:"role,omitempty" url:"role,omitempty"`
	Runbook     string `uri:"runbook,omitempty" url:"runbook,omitempty"`
	Tenant      string `uri:"tenant,omitempty" url:"tenant,omitempty"`
}
