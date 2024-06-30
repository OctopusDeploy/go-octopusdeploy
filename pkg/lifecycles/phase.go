package lifecycles

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type Phase struct {
	AutomaticDeploymentTargets         []string              `json:"AutomaticDeploymentTargets"`
	ID                                 string                `json:"Id,omitempty"`
	IsOptionalPhase                    bool                  `json:"IsOptionalPhase"`
	IsPriorityPhase                    bool                  `json:"IsPriorityPhase"`
	MinimumEnvironmentsBeforePromotion int32                 `json:"MinimumEnvironmentsBeforePromotion"`
	Name                               string                `json:"Name" validate:"required"`
	OptionalDeploymentTargets          []string              `json:"OptionalDeploymentTargets"`
	ReleaseRetentionPolicy             *core.RetentionPeriod `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy            *core.RetentionPeriod `json:"TentacleRetentionPolicy"`
}

func NewPhase(name string) *Phase {
	return &Phase{
		AutomaticDeploymentTargets: []string{},
		Name:                       name,
		OptionalDeploymentTargets:  []string{},
		ReleaseRetentionPolicy:     core.NewRetentionPeriod(30, "Days", false),
		TentacleRetentionPolicy:    core.NewRetentionPeriod(30, "Days", false),
	}
}
