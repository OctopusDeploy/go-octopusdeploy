package core

type TenantedDeploymentMode string

const (
	TenantedDeploymentModeTenanted             = TenantedDeploymentMode("Tenanted")
	TenantedDeploymentModeTenantedOrUntenanted = TenantedDeploymentMode("TenantedOrUntenanted")
	TenantedDeploymentModeUntenanted           = TenantedDeploymentMode("Untenanted")
)

type GuidedFailureMode string

const (
	GuidedFailureModeEnvironmentDefault = GuidedFailureMode("EnvironmentDefault")
	GuidedFailureModeOff                = GuidedFailureMode("Off")
	GuidedFailureModeOn                 = GuidedFailureMode("On")
)
