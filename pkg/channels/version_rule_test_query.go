package channels

type VersionRuleTestQuery struct {
	FeetType      string `uri:"feetType,omitempty" url:"feetType,omitempty"`
	PreReleaseTag string `uri:"preReleaseTag,omitempty" url:"preReleaseTag,omitempty"`
	Version       string `uri:"version,omitempty" url:"version,omitempty"`
	VersionRange  string `uri:"versionRange,omitempty" url:"versionRange,omitempty"`
}
