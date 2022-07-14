package packages

type PackagesBulkQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}
