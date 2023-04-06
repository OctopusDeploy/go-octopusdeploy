package packages

import "time"

type UploadMethod string

const (
	UploadMethodStandard = UploadMethod("Standard")
	UploadMethodDelta    = UploadMethod("Delta")
)

type UploadedPackageInfo interface {
}

// ----- DeltaUploadedPackageInfo -----

type DeltaBehaviour string

const (
	DeltaBehaviourUploadedDeltaFile = DeltaBehaviour("Uploaded delta file")
	DeltaBehaviourNoPreviousFile    = DeltaBehaviour("Uploaded full file, no previous version available to delta")
	DeltaBehaviourNotEfficient      = DeltaBehaviour("Uploaded full file, delta file was not meaningfully smaller than full file")
)

type DeltaUploadedPackageInfo struct {
	FileSize                 int64
	DeltaSize                int64
	RequestSignatureDuration time.Duration
	BuildDeltaDuration       time.Duration
	UploadDuration           time.Duration  // Time taken to upload the package (whether delta or full) depending on DeltaBehaviour
	DeltaBehaviour           DeltaBehaviour // A delta package upload can result in a standard upload if there is no previous version available, or if the delta process is not efficient. This tells you what happened
}

type PackageUploadResponseV2 struct {
	CreatedNewFile bool
	UploadMethod   UploadMethod

	// Holds information about the uploaded package.
	// If UploadMethod is UploadMethodStandard, will be nil
	// If UploadMethod is UploadMethodDelta, will be a valid struct
	UploadInfo *DeltaUploadedPackageInfo

	PackageUploadResponse
}
