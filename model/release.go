package model

import "time"

type Release struct {
	Assembled                          time.Time                                `json:"Assembled,omitempty"`
	BuildInformation                   []*ReleasePackageVersionBuildInformation `json:"BuildInformation"`
	ChannelID                          string                                   `json:"ChannelId,omitempty"`
	IgnoreChannelRules                 bool                                     `json:"IgnoreChannelRules,omitempty"`
	LibraryVariableSetSnapshotIds      []string                                 `json:"LibraryVariableSetSnapshotIds"`
	ProjectDeploymentProcessSnapshotID string                                   `json:"ProjectDeploymentProcessSnapshotId,omitempty"`
	ProjectID                          string                                   `json:"ProjectId,omitempty"`
	ProjectVariableSetSnapshotID       string                                   `json:"ProjectVariableSetSnapshotId,omitempty"`
	ReleaseNotes                       string                                   `json:"ReleaseNotes,omitempty"`
	SelectedPackages                   []*SelectedPackage                       `json:"SelectedPackages"`
	SpaceID                            string                                   `json:"SpaceId,omitempty"`
	Version                            *string                                  `json:"Version"`

	Resource
}

// Releases defines a collection of Release instance with built-in support for
// paged results from the API.
type Releases struct {
	Items []Release `json:"Items"`

	PagedResults
}
