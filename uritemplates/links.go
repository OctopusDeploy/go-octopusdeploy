package uritemplates

// The octopus server promises not to break /api links, and doesn't give out hypermedia links for clients
// to follow, so we must maintain a big list of URLs so we can perform requests.
// NOTE: If the server *were* to ever remove an api link, we'd need to do some sort of version detection to compensate.

const (
	CreateReleaseCommandV1              = "/api/{spaceId}/releases/create/v1"
	CreateDeploymentTenantedCommandV1   = "/api/{spaceId}/deployments/create/tenanted/v1"
	CreateDeploymentUntenantedCommandV1 = "/api/{spaceId}/deployments/create/untenanted/v1"
	CreateRunRunbookCommand             = "/api/{spaceId}/runbook-runs/create/v1"
	DeploymentProcesses                 = "/api/{spaceId}/deploymentprocesses{/id}{?skip,take,ids}"
	ReleaseDeploymentPreview            = "/api/{spaceId}/releases/{releaseId}/deployments/preview/{environmentId}{?includeDisabledSteps}"
	Releases                            = "/api/{spaceId}/releases{/id}{?skip,ignoreChannelRules,take,ids}"
	ReleasesByProject                   = "/api/{spaceId}/projects/{projectId}/releases{/version}{?skip,take,searchByVersion}"
	ReleasesByProjectAndChannel         = "/api/{spaceId}/projects/{projectId}/channels/{channelId}/releases{?skip,take,searchByVersion}"
	Variables                           = "/api/{spaceId}/variables{/id}{?ids}"
)
