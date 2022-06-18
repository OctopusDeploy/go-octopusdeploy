package lifecycles

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/core"

type Phase struct {
	AutomaticDeploymentTargets         []string              `json:"AutomaticDeploymentTargets"`
	ID                                 string                `json:"Id,omitempty"`
	IsOptionalPhase                    bool                  `json:"IsOptionalPhase"`
	MinimumEnvironmentsBeforePromotion int32                 `json:"MinimumEnvironmentsBeforePromotion"`
	Name                               string                `json:"Name" validate:"required"`
	OptionalDeploymentTargets          []string              `json:"OptionalDeploymentTargets"`
	ReleaseRetentionPolicy             *core.RetentionPeriod `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy            *core.RetentionPeriod `json:"TentacleRetentionPolicy"`
}
