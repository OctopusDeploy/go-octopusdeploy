package variables

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type VariableScopeValues struct {
	Actions      []*resources.ReferenceDataItem        `json:"Actions"`
	Channels     []*resources.ReferenceDataItem        `json:"Channels"`
	Environments []*resources.ReferenceDataItem        `json:"Environments"`
	Machines     []*resources.ReferenceDataItem        `json:"Machines"`
	Processes    []*resources.ProcessReferenceDataItem `json:"Processes"`
	Roles        []*resources.ReferenceDataItem        `json:"Roles"`
	TenantTags   []*resources.ReferenceDataItem        `json:"TenantTags"`
}
