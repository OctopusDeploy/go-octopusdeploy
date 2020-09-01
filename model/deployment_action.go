package model

type DeploymentAction struct {
	ID                            string             `json:"Id,omitempty"`
	Name                          string             `json:"Name"`
	ActionType                    string             `json:"ActionType"`
	IsDisabled                    bool               `json:"IsDisabled"`
	IsRequired                    bool               `json:"IsRequired"`
	WorkerPoolID                  string             `json:"WorkerPoolId,omitempty"`
	CanBeUsedForProjectVersioning bool               `json:"CanBeUsedForProjectVersioning"`
	Environments                  []string           `json:"Environments,omitempty"`
	ExcludedEnvironments          []string           `json:"ExcludedEnvironments,omitempty"`
	Channels                      []string           `json:"Channels,omitempty"`
	TenantTags                    []string           `json:"TenantTags,omitempty"`
	Properties                    map[string]string  `json:"Properties"` // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Packages                      []PackageReference `json:"Packages,omitempty"`
}
