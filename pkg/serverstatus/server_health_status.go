package serverstatus

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type ServerHealthStatus struct {
	Description                  string `json:"Description,omitempty"`
	IsCompliantWithLicense       bool   `json:"IsCompliantWithLicense"`
	IsEntireClusterDrainingTasks bool   `json:"IsEntireClusterDrainingTasks"`
	IsEntireClusterReadOnly      bool   `json:"IsEntireClusterReadOnly"`
	IsOperatingNormally          bool   `json:"IsOperatingNormally"`

	resources.Resource
}
