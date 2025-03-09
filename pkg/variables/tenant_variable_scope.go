package variables

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type TenantVariableScope struct {
	EnvironmentIds []string `json:"EnvironmentIds"`

	resources.Resource
}
