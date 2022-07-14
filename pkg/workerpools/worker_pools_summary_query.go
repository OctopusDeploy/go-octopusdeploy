package workerpools

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	HealthStatuses       []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	HideEmptyWorkerPools bool     `uri:"hideEmptyWorkerPools,omitempty" url:"hideEmptyWorkerPools,omitempty"`
	IDs                  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled           bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	MachinePartialName   string   `uri:"machinePartialName,omitempty" url:"machinePartialName,omitempty"`
	PartialName          string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ShellNames           []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
}
