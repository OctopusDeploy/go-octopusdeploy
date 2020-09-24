package model

type VariableScopeValues struct {
	Actions      []*ReferenceDataItem        `json:"Actions"`
	Channels     []*ReferenceDataItem        `json:"Channels"`
	Environments []*ReferenceDataItem        `json:"Environments"`
	Machines     []*ReferenceDataItem        `json:"Machines"`
	Processes    []*ProcessReferenceDataItem `json:"Processes"`
	Roles        []*ReferenceDataItem        `json:"Roles"`
	TenantTags   []*ReferenceDataItem        `json:"TenantTags"`
}
