package machines

type WorkersQuery struct {
	CommunicationStyles   []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	DeploymentTargetTypes []string `uri:"deploymentTargetTypes,omitempty" url:"deploymentTargetType,omitempty"`
	HealthStatuses        []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	IDs                   []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled            bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	Name                  string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName           string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ShellNames            []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	Skip                  int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                  int      `uri:"take,omitempty" url:"take,omitempty"`
	Thumbprint            string   `uri:"thumbprint,omitempty" url:"thumbprint,omitempty"`
	WorkerPoolIDs         []string `uri:"workerPoolIds" url:"workerPoolIds,omitempty"`
}
