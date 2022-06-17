package packages

type PackageUploadQuery struct {
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
}
