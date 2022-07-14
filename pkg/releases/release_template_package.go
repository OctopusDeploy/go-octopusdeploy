package releases

type ReleaseTemplatePackage struct {
	ActionName                 string `json:"ActionName,omitempty"`
	FeedID                     string `json:"FeedId,omitempty"`
	FeedName                   string `json:"FeedName,omitempty"`
	IsResolvable               bool   `json:"IsResolvable,omitempty"`
	NuGetFeedID                string `json:"NuGetFeedId,omitempty"`
	NuGetFeedName              string `json:"NuGetFeedName,omitempty"`
	NuGetPackageID             string `json:"NuGetPackageId,omitempty"`
	PackageID                  string `json:"PackageId,omitempty"`
	PackageReferenceName       string `json:"PackageReferenceName,omitempty"`
	ProjectName                string `json:"ProjectName,omitempty"`
	StepName                   string `json:"StepName,omitempty"`
	VersionSelectedLastRelease string `json:"VersionSelectedLastRelease,omitempty"`
}
