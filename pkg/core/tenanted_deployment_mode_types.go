package octopusdeploy

type TenantedDeploymentMode string

const (
	TenantedDeploymentModeTenanted             = TenantedDeploymentMode("Tenanted")
	TenantedDeploymentModeTenantedOrUntenanted = TenantedDeploymentMode("TenantedOrUntenanted")
	TenantedDeploymentModeUntenanted           = TenantedDeploymentMode("Untenanted")
)
