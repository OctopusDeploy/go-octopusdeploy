package environments

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	HealthStatuses        []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	HideEmptyEnvironments bool     `uri:"hideEmptyEnvironments,omitempty" url:"hideEmptyEnvironments,omitempty"`
	IDs                   []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled            bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	MachinePartialName    string   `uri:"machinePartialName,omitempty" url:"machinePartialName,omitempty"`
	PartialName           string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Roles                 []string `uri:"roles,omitempty" url:"roles,omitempty"`
	ShellNames            []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	TenantIDs             []string `uri:"tenantIds,omitempty" url:"tenantIds,omitempty"`
	TenantTags            []string `uri:"tenantTags,omitempty" url:"tenantTags,omitempty"`
}
