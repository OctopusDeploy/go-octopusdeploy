package serverstatus

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type SystemInfo struct {
	ClrVersion         string `json:"ClrVersion,omitempty"`
	MinThreadPoolCount int    `json:"MinThreadPoolCount"`
	OSVersion          string `json:"OSVersion,omitempty"`
	ThreadCount        int    `json:"ThreadCount"`
	// Uptime is a .NET time span, "[d.]hh:mm:ss[.fffffff]", as sent by the server.
	Uptime          string `json:"Uptime,omitempty"`
	Version         string `json:"Version,omitempty"`
	WorkingSetBytes int64  `json:"WorkingSetBytes"`

	resources.Resource
}
