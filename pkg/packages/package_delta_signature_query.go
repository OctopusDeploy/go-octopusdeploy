package packages

type PackageDeltaSignatureQuery struct {
	PackageID string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	Version   string `uri:"version,omitempty" url:"version,omitempty"`
}
