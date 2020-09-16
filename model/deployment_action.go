package model

type DeploymentAction struct {
	ActionType                    string             `json:"ActionType"`
	CanBeUsedForProjectVersioning bool               `json:"CanBeUsedForProjectVersioning"`
	Channels                      []string           `json:"Channels,omitempty"`
	Environments                  []string           `json:"Environments,omitempty"`
	ExcludedEnvironments          []string           `json:"ExcludedEnvironments,omitempty"`
	ID                            string             `json:"Id,omitempty"`
	IsDisabled                    bool               `json:"IsDisabled"`
	IsRequired                    bool               `json:"IsRequired"`
	Name                          string             `json:"Name"`
	Packages                      []PackageReference `json:"Packages,omitempty"`
	Properties                    map[string]string  `json:"Properties"` // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	TenantTags                    []string           `json:"TenantTags,omitempty"`
	WorkerPoolID                  string             `json:"WorkerPoolId,omitempty"`
}

// NewDeploymentAction initializes a DeploymentAction with a name.
func NewDeploymentAction(name string) (*DeploymentAction, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewDeploymentAction", "name")
	}

	return &DeploymentAction{
		Name: name,
	}, nil
}
