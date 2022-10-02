package uritemplates

// The octopus server promises not to break /api links, and doesn't give out hypermedia links for clients
// to follow, so we must maintain a big list of URLs so we can perform requests.
// NOTE: If the server *were* to ever remove an api link, we'd need to do some sort of version detection to compensate.

const (
	BuildInformation                    = "/api/{spaceId}/build-information{/id}{?packageId,filter,latest,skip,take,overwriteMode}"
	BuildInformationBulk                = "/api/{spaceId}/build-information/bulk{?ids}"
	CreateReleaseCommandV1              = "/api/{spaceId}/releases/create/v1"                                                                                                             // POST
	CreateDeploymentTenantedCommandV1   = "/api/{spaceId}/deployments/create/tenanted/v1"                                                                                                 // POST
	CreateDeploymentUntenantedCommandV1 = "/api/{spaceId}/deployments/create/untenanted/v1"                                                                                               // POST
	CreateRunRunbookCommand             = "/api/{spaceId}/runbook-runs/create/v1"                                                                                                         // POST
	DeploymentProcesses                 = "/api/{spaceId}/deploymentprocesses{/id}{?skip,take,ids}"                                                                                       // GET
	FeedSearchPackageVersions           = "/api/{spaceId}/feeds/{feedId}/packages/versions{?packageId,take,skip,includePreRelease,versionRange,preReleaseTag,filter,includeReleaseNotes}" // GET
	Packages                            = "/api/{spaceId}/packages{/id}{?nuGetPackageId,filter,latest,skip,take,includeNotes}"                                                            // GET
	PackageUpload                       = "/api/{spaceId}/packages/raw{?replace,overwriteMode}"                                                                                           // POST multipart form
	ReleaseDeploymentPreview            = "/api/{spaceId}/releases/{releaseId}/deployments/preview/{environmentId}{?includeDisabledSteps}"                                                // GET
	Releases                            = "/api/{spaceId}/releases{/id}{?skip,ignoreChannelRules,take,ids}"                                                                               // GET
	ReleasesByProject                   = "/api/{spaceId}/projects/{projectId}/releases{/version}{?skip,take,searchByVersion}"                                                            // GET
	ReleasesByProjectAndChannel         = "/api/{spaceId}/projects/{projectId}/channels/{channelId}/releases{?skip,take,searchByVersion}"                                                 // GET
	Runbooks                            = "/api/{spaceId}/runbooks{/id}{?skip,take,ids,partialName,clone,projectIds}"                                                                     // GET
	RunbooksByProject                   = "/api/{spaceId}/projects/{projectId}/runbooks{?skip,take,partialName}"                                                                          // GET
	RunbookEnvironments                 = "/api/{spaceId}/projects/{projectId}/runbooks/{runbookId}/environments"                                                                         // GET
	RunbookProcess                      = "/api/{spaceId}/projects/{projectId}/runbookProcesses/{id}"                                                                                     // GET
	RunbookRunPreview                   = "/api/{spaceId}/projects/{projectId}/runbooks/{runbookId}/runbookRuns/preview/{environment}{?includeDisabledSteps}"                             // GET
	RunbookSnapshotsByRunbook           = "/api/{spaceId}/projects/{projectId}/runbooks/{runbookId}/runbookSnapshots{/name}{?skip,take,searchByName}"                                     // GET
	RunbookSnapshotsByProject           = "/api/{spaceId}/projects/{projectId}/runbookSnapshots{/name}{?skip,take,searchByName}"                                                          // GET
	RunbookSnapshotRunPreview           = "/api/{spaceId}/runbookSnapshots/{snapshotId}/runbookRuns/preview/{environmentId}{?includeDisabledSteps}"                                       // GET
	RunbookRunTenantPreview             = "/api/{spaceId}/projects/{projectId}/runbooks/{runbookId}/runbookRuns/previews"                                                                 // POST
	Variables                           = "/api/{spaceId}/variables{/id}{?ids}"                                                                                                           // GET
)
