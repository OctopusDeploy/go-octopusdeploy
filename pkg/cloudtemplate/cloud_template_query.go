package cloudtemplate

type CloudTemplateQuery struct {
	FeedID    string `uri:"feedId,omitempty" url:"feedId,omitempty"`
	PackageID string `uri:"packageId,omitempty" url:"packageId,omitempty"`
}
