package feeds

type SearchPackageVersionsQuery struct {
	FeedID              string `uri:"id,omitempty"`
	Filter              string `uri:"filter,omitempty"`
	IncludePreRelease   bool   `uri:"includePreRelease,omitempty"`
	IncludeReleaseNotes bool   `uri:"includeReleaseNotes,omitempty"`
	PackageID           string `uri:"packageId,omitempty"`
	PreReleaseTag       string `uri:"preReleaseTag,omitempty"`
	Skip                int    `uri:"skip,omitempty"`
	Take                int    `uri:"take,omitempty"`
	VersionRange        string `uri:"versionRange,omitempty"`
}
