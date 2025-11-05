package retention

type SpaceDefaultRetentionPolicyQuery struct {
	RetentionType RetentionType `uri:"RetentionType"`
	SpaceID       string        `uri:"spaceId"`
}
