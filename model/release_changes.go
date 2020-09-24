package model

type ReleaseChanges struct {
	BuildInformation []*ReleasePackageVersionBuildInformation `json:"BuildInformation"`
	Commits          []*CommitDetails                         `json:"Commits"`
	ReleaseNotes     string                                   `json:"ReleaseNotes,omitempty"`
	Version          string                                   `json:"Version,omitempty"`
	WorkItems        []*WorkItemLink                          `json:"WorkItems"`
}
