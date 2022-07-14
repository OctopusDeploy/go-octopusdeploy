package packages

type PackageDeltaUploadQuery struct {
	BaseVersion   string `uri:"baseVersion,omitempty" url:"baseVersion,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	PackageID     string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
}
