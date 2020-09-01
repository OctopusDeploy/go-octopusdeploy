package model

type ChannelRule struct {
	// name of Package step(s) this rule applies to
	Actions []string `json:"Actions,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// Pre-release tag
	Tag string `json:"Tag,omitempty"`

	//Use the NuGet or Maven versioning syntax (depending on the feed type)
	//to specify the range of versions to include
	VersionRange string `json:"VersionRange,omitempty"`
}
