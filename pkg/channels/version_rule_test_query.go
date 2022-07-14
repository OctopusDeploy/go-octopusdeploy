package channels

type VersionRuleTestQuery struct {
	FeedType      string `uri:"feedType,omitempty" url:"feedType,omitempty"`
	PreReleaseTag string `uri:"preReleaseTag,omitempty" url:"preReleaseTag,omitempty"`
	Version       string `uri:"version,omitempty" url:"version,omitempty"`
	VersionRange  string `uri:"versionRange,omitempty" url:"versionRange,omitempty"`
}
