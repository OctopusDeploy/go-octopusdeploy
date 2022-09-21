package buildinformation

type OverwriteMode string

const (
	OverwriteModeFailIfExists      = OverwriteMode("FailIfExists")
	OverwriteModeIgnoreIfExists    = OverwriteMode("IgnoreIfExists")
	OverwriteModeOverwriteExisting = OverwriteMode("OverwriteExisting")
)
