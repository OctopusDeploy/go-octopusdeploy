package octopusdeploy

type Phase struct {
	ID                                 string           `json:"Id,omitempty"`
	Name                               string           `json:"Name" validate:"required"`
	MinimumEnvironmentsBeforePromotion int32            `json:"MinimumEnvironmentsBeforePromotion"`
	IsOptionalPhase                    bool             `json:"IsOptionalPhase"`
	ReleaseRetentionPolicy             *RetentionPeriod `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy            *RetentionPeriod `json:"TentacleRetentionPolicy"`
	AutomaticDeploymentTargets         []string         `json:"AutomaticDeploymentTargets"`
	OptionalDeploymentTargets          []string         `json:"OptionalDeploymentTargets"`
}
