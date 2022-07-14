package packages

type PackageNotesListQuery struct {
	PackageIDs []string `uri:"packageIds,omitempty" url:"packageIds,omitempty"`
}
