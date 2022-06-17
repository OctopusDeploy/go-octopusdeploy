package tasks

type TasksQuery struct {
	Environment             string   `uri:"environment,omitempty" url:"environment,omitempty"`
	HasPendingInterruptions bool     `uri:"hasPendingInterruptions,omitempty" url:"hasPendingInterruptions,omitempty"`
	HasWarningsOrErrors     bool     `uri:"hasWarningsOrErrors,omitempty" url:"hasWarningsOrErrors,omitempty"`
	IDs                     []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem           bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	IsActive                bool     `uri:"active,omitempty" url:"active,omitempty"`
	IsRunning               bool     `uri:"running,omitempty" url:"running,omitempty"`
	Name                    string   `uri:"name,omitempty" url:"name,omitempty"`
	Node                    string   `uri:"node,omitempty" url:"node,omitempty"`
	PartialName             string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Project                 string   `uri:"project,omitempty" url:"project,omitempty"`
	Runbook                 string   `uri:"runbook,omitempty" url:"runbook,omitempty"`
	Skip                    int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces                  []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	States                  []string `uri:"states,omitempty" url:"states,omitempty"`
	Take                    int      `uri:"take,omitempty" url:"take,omitempty"`
	Tenant                  string   `uri:"tenant,omitempty" url:"tenant,omitempty"`
}
