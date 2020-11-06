package octopusdeploy

type Phase struct {
	AutomaticDeploymentTargets         []string         `json:"AutomaticDeploymentTargets"`
	ID                                 string           `json:"Id,omitempty"`
	IsOptionalPhase                    bool             `json:"IsOptionalPhase"`
	MinimumEnvironmentsBeforePromotion int32            `json:"MinimumEnvironmentsBeforePromotion"`
	Name                               string           `json:"Name" validate:"required"`
	OptionalDeploymentTargets          []string         `json:"OptionalDeploymentTargets"`
	ReleaseRetentionPolicy             *RetentionPeriod `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy            *RetentionPeriod `json:"TentacleRetentionPolicy"`
}
