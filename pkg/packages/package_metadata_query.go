package packages

type PackageMetadataQuery struct {
	Filter        string `uri:"filter,omitempty" url:"filter,omitempty"`
	Latest        string `uri:"latest,omitempty" url:"latest,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
	Skip          int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take          int    `uri:"take,omitempty" url:"take,omitempty"`
}
