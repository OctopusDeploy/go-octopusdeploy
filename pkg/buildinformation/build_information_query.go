package buildinformation

type BuildInformationQuery struct {
	Filter           string `uri:"filter,omitempty" url:"filter,omitempty"`
	Latest           bool   `uri:"latest,omitempty" url:"latest,omitempty"`
	OverwriteMode    string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	PackageID        string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	IncludeWorkItems bool   `uri:"includeWorkItems,omitempty" url:"includeWorkItems,omitempty"`
	Skip             int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take             int    `uri:"take,omitempty" url:"take,omitempty"`
}
