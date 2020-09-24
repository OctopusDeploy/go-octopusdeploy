package model

type ChannelRule struct {
	Actions []string `json:"Actions,omitempty"`
	ID      string   `json:"Id,omitempty"`
	Tag     string   `json:"Tag,omitempty"`

	//Use the NuGet or Maven versioning syntax (depending on the feed type)
	//to specify the range of versions to include
	VersionRange string `json:"VersionRange,omitempty"`
}
