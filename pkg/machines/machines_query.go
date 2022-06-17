package machines

type MachinesQuery struct {
	CommunicationStyles []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	DeploymentID        string   `uri:"deploymentId,omitempty" url:"deploymentId,omitempty"`
	EnvironmentIDs      []string `uri:"environmentIds,omitempty" url:"environmentIds,omitempty"`
	HealthStatuses      []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	IDs                 []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled          bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	Name                string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName         string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Roles               []string `uri:"roles,omitempty" url:"roles,omitempty"`
	ShellNames          []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	Skip                int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                int      `uri:"take,omitempty" url:"take,omitempty"`
	TenantIDs           []string `uri:"tenantIds,omitempty" url:"tenantIds,omitempty"`
	TenantTags          []string `uri:"tenantTags,omitempty" url:"tenantTags,omitempty"`
	Thumbprint          string   `uri:"thumbprint,omitempty" url:"thumbprint,omitempty"`
}
