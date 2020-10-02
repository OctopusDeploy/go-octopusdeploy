package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Release struct {
	Assembled                          time.Time                                `json:"Assembled,omitempty"`
	BuildInformation                   []*ReleasePackageVersionBuildInformation `json:"BuildInformation"`
	ChannelID                          string                                   `json:"ChannelId,omitempty"`
	IgnoreChannelRules                 bool                                     `json:"IgnoreChannelRules,omitempty"`
	LibraryVariableSetSnapshotIDs      []string                                 `json:"LibraryVariableSetSnapshotIds"`
	ProjectDeploymentProcessSnapshotID string                                   `json:"ProjectDeploymentProcessSnapshotId,omitempty"`
	ProjectID                          string                                   `json:"ProjectId,omitempty"`
	ProjectVariableSetSnapshotID       string                                   `json:"ProjectVariableSetSnapshotId,omitempty"`
	ReleaseNotes                       string                                   `json:"ReleaseNotes,omitempty"`
	SelectedPackages                   []*SelectedPackage                       `json:"SelectedPackages"`
	SpaceID                            string                                   `json:"SpaceId,omitempty"`
	Version                            *string                                  `json:"Version"`

	Resource
}

// Releases defines a collection of Release instance with built-in support for paged results from the API.
type Releases struct {
	Items []Release `json:"Items"`
	PagedResults
}

// GetID returns the ID value of the Release.
func (resource Release) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Release.
func (resource Release) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Release was changed.
func (resource Release) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Release.
func (resource Release) GetLinks() map[string]string {
	return resource.Links
}

func (resource Release) SetID(id string) {
	resource.ID = id
}

func (resource Release) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource Release) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Release and returns an error if invalid.
func (resource Release) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &Release{}
