package packages

type PackagesQuery struct {
	Filter         string `uri:"filter,omitempty" url:"filter,omitempty"`
	IncludeNotes   bool   `uri:"includeNotes,omitempty" url:"includeNotes,omitempty"`
	Latest         string `uri:"latest,omitempty" url:"latest,omitempty"`
	NuGetPackageID string `uri:"nuGetPackageId,omitempty" url:"nuGetPackageId,omitempty"`
	Skip           int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take           int    `uri:"take,omitempty" url:"take,omitempty"`
}
