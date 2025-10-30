package retention

type RetentionType string

const (
	LifecycleReleaseRetentionType = RetentionType("LifecycleRelease")
	LifecycleTentacleRetentionType = RetentionType("LifecycleTentacle")
)