package octopusdeploy

type ScopeValues struct {
	Environments []ScopeValue `json:"Environments"`
	Machines     []ScopeValue `json:"Machines"`
	Actions      []ScopeValue `json:"Actions"`
	Roles        []ScopeValue `json:"Roles"`
	Channels     []ScopeValue `json:"Channels"`
	TenantTags   []ScopeValue `json:"TenantTags"`
}

type ScopeValue struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}
