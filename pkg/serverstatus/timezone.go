package serverstatus

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type Timezone struct {
	IsLocal bool   `json:"IsLocal"`
	Name    string `json:"Name,omitempty"`

	resources.Resource
}
