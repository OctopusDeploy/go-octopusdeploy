package retention

type RetentionType string

const (
	LifecycleReleaseRetentionType  = RetentionType("LifecycleRelease")
	LifecycleTentacleRetentionType = RetentionType("LifecycleTentacle")
	RunbookRetentionType           = RetentionType("RunbookRetention")
)

const (
	RetentionStrategyForever string = "Forever"
	RetentionStrategyCount   string = "Count"
)

const (
	RetentionUnitDays  string = "Days"
	RetentionUnitItems string = "Items"
)
