package client

const (
	TestURISelf                              = "/api"
	TestURIAccounts                          = "/api/Spaces-1/accounts{/id}{?skip,take,ids,partialName,accountType}"
	TestURIActionTemplateLogo                = "/api/Spaces-1/actiontemplates/{typeOrId}/logo{?cb}"
	TestURIActionTemplates                   = "/api/Spaces-1/actiontemplates{/id}{?skip,take,ids,partialName}"
	TestURIActionTemplatesCategories         = "/api/Spaces-1/actiontemplates/categories"
	TestURIActionTemplatesSearch             = "/api/Spaces-1/actiontemplates/search"
	TestURIActionTemplateVersionedLogo       = "/api/Spaces-1/actiontemplates/{typeOrId}/versions/{version}/logo{?cb}"
	TestURIAPIKeys                           = "/api/users/{userId}/apikeys"
	TestURIArtifacts                         = "/api/Spaces-1/artifacts{/id}{?skip,take,regarding,ids,partialName,order}"
	TestURIAuthenticateOctopusID             = "/users/authenticate/OctopusID{?returnUrl}"
	TestURIAuthentication                    = "/api/authentication"
	TestURIAzureDevOpsConnectivityCheck      = "/api/azuredevopsissuetracker/connectivitycheck"
	TestURIAzureEnvironments                 = "/api/accounts/azureenvironments"
	TestURIBuildInformation                  = "/api/Spaces-1/build-information{/id}{?filter,packageId,latest,skip,take,overwriteMode}"
	TestURIBuildInformationBulk              = "/api/Spaces-1/build-information/bulk{?ids}"
	TestURIBuiltInFeedStats                  = "/api/feeds/stats"
	TestURICertificateConfiguration          = "/api/configuration/certificates{/id}{?skip,take,ids,partialName}"
	TestURICertificates                      = "/api/Spaces-1/certificates{/id}{?skip,take,search,archived,tenant,firstResult,orderBy,ids,partialName}"
	TestURIChannels                          = "/api/Spaces-1/channels{/id}{?skip,take,ids,partialName}"
	TestURICloudTemplate                     = "/api/cloudtemplate/{id}/metadata{?packageId,feedId}"
	TestURICommunityActionTemplates          = "/api/communityactiontemplates{/id}{?skip,take,ids}"
	TestURIConfiguration                     = "/api/configuration{/id}"
	TestURICurrentLicense                    = "/api/licenses/licenses-current"
	TestURICurrentLicenseStatus              = "/api/licenses/licenses-current-status"
	TestURICurrentUser                       = "/api/users/me"
	TestURIDashboard                         = "/api/Spaces-1/dashboard{?projectId,releaseId,selectedTenants,selectedTags,showAll,highestLatestVersionPerProjectAndEnvironment}"
	TestURIDashboardConfiguration            = "/api/Spaces-1/dashboardconfiguration"
	TestURIDashboardDynamic                  = "/api/Spaces-1/dashboard/dynamic{?projects,environments,includePrevious}"
	TestURIDeploymentProcesses               = "/api/Spaces-1/deploymentprocesses{/id}{?skip,take,ids}"
	TestURIDeployments                       = "/api/Spaces-1/deployments{/id}{?skip,take,ids,projects,environments,tenants,channels,taskState,partialName}"
	TestURIDiscoverMachine                   = "/api/Spaces-1/machines/discover{?host,port,type,proxyId}"
	TestURIDiscoverWorker                    = "/api/Spaces-1/workers/discover{?host,port,type,proxyId}"
	TestURIDynamicExtensionsFeaturesMetadata = "/api/dynamic-extensions/features/metadata"
	TestURIDynamicExtensionsFeaturesValues   = "/api/dynamic-extensions/features/values"
	TestURIDynamicExtensionsScripts          = "/api/dynamic-extensions/scripts"
	TestURIEnvironments                      = "/api/Spaces-1/environments{/id}{?name,skip,ids,take,partialName}"
	TestURIEnvironmentSortOrder              = "/api/Spaces-1/environments/sortorder"
	TestURIEnvironmentsSummary               = "/api/Spaces-1/environments/summary{?ids,partialName,machinePartialName,roles,isDisabled,healthStatuses,commStyles,tenantIds,tenantTags,hideEmptyEnvironments,shellNames}"
	TestURIEventAgents                       = "/api/events/agents"
	TestURIEventCategories                   = "/api/events/categories{?appliesTo}"
	TestURIEventDocumentTypes                = "/api/events/documenttypes"
	TestURIEventGroups                       = "/api/events/groups{?appliesTo}"
	TestURIEvents                            = "/api/events{/id}{?skip,regarding,regardingAny,user,users,projects,projectGroups,environments,eventGroups,eventCategories,eventAgents,tags,tenants,from,to,internal,fromAutoId,toAutoId,documentTypes,asCsv,take,ids,spaces,includeSystem,excludeDifference}"
	TestURIExtensionStats                    = "/api/serverstatus/extensions"
	TestURIExternalSecurityGroupProviders    = "/api/externalsecuritygroupproviders"
	TestURIExternalUserSearch                = "/api/users/external-search{?partialName}"
	TestURIFeaturesConfiguration             = "/api/featuresconfiguration"
	TestURIFeeds                             = "/api/feeds{/id}{?skip,take,ids,partialName,feedType}"
	TestURIInterruptions                     = "/api/Spaces-1/interruptions{/id}{?skip,take,regarding,pendingOnly,ids}"
	TestURIInvitations                       = "/api/users/invitations"
	TestURIIssueTrackers                     = "/api/issuetrackers{?skip,take,ids,partialName}"
	TestURIJiraConnectAppCredentialsTest     = "/api/jiraintegration/connectivitycheck/connectapp"
	TestURIJiraCredentialsTest               = "/api/jiraintegration/connectivitycheck/jira"
	TestURILetsEncryptConfiguration          = "/api/letsencryptconfiguration"
	TestURILibraryVariables                  = "/api/Spaces-1/libraryvariablesets{/id}{?skip,contentType,take,ids,partialName}"
	TestURILifecycles                        = "/api/Spaces-1/lifecycles{/id}{?skip,take,ids,partialName}"
	TestURILoginInitiated                    = "/api/authentication/checklogininitiated"
	TestURIMachineOperatingSystems           = "/api/Spaces-1/machines/operatingsystem/names/all"
	TestURIMachinePolicies                   = "/api/Spaces-1/machinepolicies{/id}{?skip,take,ids,partialName}"
	TestURIMachinePolicyTemplate             = "/api/Spaces-1/machinepolicies/template"
	TestURIMachineRoles                      = "/api/Spaces-1/machineroles/all"
	TestURIMachines                          = "/api/Spaces-1/machines{/id}{?skip,take,name,ids,partialName,roles,isDisabled,healthStatuses,commStyles,tenantIds,tenantTags,environmentIds,thumbprint,deploymentId,shellNames}"
	TestURIMachineShells                     = "/api/Spaces-1/machines/operatingsystem/shells/all"
	TestURIMaintenanceConfiguration          = "/api/maintenanceconfiguration"
	TestURIMigrationsImport                  = "/api/migrations/import"
	TestURIMigrationsPartialExport           = "/api/migrations/partialexport"
	TestURIOctopusServerClusterSummary       = "/api/octopusservernodes/summary"
	TestURIOctopusServerNodes                = "/api/octopusservernodes{/id}{?skip,take,ids,partialName}"
	TestURIPackageDeltaSignature             = "/api/Spaces-1/packages/{packageId}/{version}/delta-signature"
	TestURIPackageDeltaUpload                = "/api/Spaces-1/packages/{packageId}/{baseVersion}/delta{?replace,overwriteMode}"
	TestURIPackageMetadata                   = "/api/Spaces-1/package-metadata{/id}{?filter,latest,skip,take,replace,overwriteMode}"
	TestURIPackageNotesList                  = "/api/Spaces-1/packages/notes{?packageIds}"
	TestURIPackages                          = "/api/Spaces-1/packages{/id}{?nuGetPackageId,filter,latest,skip,take,includeNotes}"
	TestURIPackagesBulk                      = "/api/Spaces-1/packages/bulk{?ids}"
	TestURIPackageUpload                     = "/api/Spaces-1/packages/raw{?replace,overwriteMode}"
	TestURIPerformanceConfiguration          = "/api/performanceconfiguration"
	TestURIPermissionDescriptions            = "/api/permissions/all"
	TestURIProjectGroups                     = "/api/Spaces-1/projectgroups{/id}{?skip,take,ids,partialName}"
	TestURIProjectPulse                      = "/api/Spaces-1/projects/pulse{?projectIds}"
	TestURIProjects                          = "/api/Spaces-1/projects{/id}{?name,skip,ids,clone,take,partialName,clonedFromProjectId}"
	TestURIProjectsExperimentalSummaries     = "/api/Spaces-1/projects/experimental/summaries{?ids}"
	TestURIProjectTriggers                   = "/api/Spaces-1/projecttriggers{/id}{?skip,take,ids,runbooks}"
	TestURIProxies                           = "/api/Spaces-1/proxies{/id}{?skip,take,ids,partialName}"
	TestURIRegister                          = "/api/users/register"
	TestURIReleases                          = "/api/Spaces-1/releases{/id}{?skip,ignoreChannelRules,take,ids}"
	TestURIReportingDeploymentsCountedByWeek = "/api/Spaces-1/reporting/deployments-counted-by-week{?projectIds}"
	TestURIRunbookProcesses                  = "/api/Spaces-1/runbookProcesses{/id}{?skip,take,ids}"
	TestURIRunbookRuns                       = "/api/Spaces-1/runbookRuns{/id}{?skip,take,ids,projects,environments,tenants,runbooks,taskState,partialName}"
	TestURIRunbooks                          = "/api/Spaces-1/runbooks{/id}{?skip,take,ids,partialName,clone,projectIds}"
	TestURIRunbookSnapshots                  = "/api/Spaces-1/runbookSnapshots{/id}{?skip,take,ids,publish}"
	TestURIScheduledProjectTriggers          = "/api/Spaces-1/scheduledprojecttriggers{/id}{?skip,take,ids}"
	TestURIScheduler                         = "/api/scheduler/{name}/logs{?verbose,tail}"
	TestURIScopedUserRoles                   = "/api/scopeduserroles{/id}{?skip,take,ids,partialName,spaces,includeSystem}"
	TestURIServerConfiguration               = "/api/serverconfiguration"
	TestURIServerConfigurationSettings       = "/api/serverconfiguration/settings"
	TestURIServerHealthStatus                = "/api/serverstatus/health"
	TestURIServerStatus                      = "/api/serverstatus"
	TestURISignIn                            = "/api/users/login{?returnUrl}"
	TestURISignOut                           = "/api/users/logout"
	TestURISmtpConfiguration                 = "/api/smtpconfiguration"
	TestURISmtpIsConfigured                  = "/api/smtpconfiguration/isconfigured"
	TestURISpaceHome                         = "/api/{spaceId}"
	TestURISpaces                            = "/api/spaces{/id}{?name,skip,ids,take,partialName}"
	TestURISubscriptions                     = "/api/Spaces-1/subscriptions{/id}{?skip,take,ids,partialName,spaces}"
	TestURITagSets                           = "/api/Spaces-1/tagsets{/id}{?skip,take,ids,partialName}"
	TestURITagSetSortOrder                   = "/api/Spaces-1/tagsets/sortorder"
	TestURITasks                             = "/api/tasks{/id}{?skip,active,environment,tenant,runbook,project,name,node,running,states,hasPendingInterruptions,hasWarningsOrErrors,take,ids,partialName,spaces,includeSystem}"
	TestURITaskTypes                         = "/api/tasks/taskTypes"
	TestURITeamMembership                    = "/api/teammembership{?userId,spaces,includeSystem}"
	TestURITeamMembershipPreviewTeam         = "/api/teammembership/previewteam"
	TestURITeams                             = "/api/teams{/id}{?skip,take,ids,partialName,spaces,includeSystem}"
	TestURITenants                           = "/api/Spaces-1/tenants{/id}{?skip,projectId,name,tags,take,ids,clone,partialName,clonedFromTenantId}"
	TestURITenantsMissingVariables           = "/api/Spaces-1/tenants/variables-missing{?tenantId,projectId,environmentId,includeDetails}"
	TestURITenantsStatus                     = "/api/Spaces-1/tenants/status"
	TestURITenantTagTest                     = "/api/Spaces-1/tenants/tag-test{?tenantIds,tags}"
	TestURITenantVariables                   = "/api/Spaces-1/tenantvariables/all{?projectId}"
	TestURITimezones                         = "/api/serverstatus/timezones"
	TestURIUpgradeConfiguration              = "/api/upgradeconfiguration"
	TestURIUserAuthentication                = "/api/users/authentication{/userId}"
	TestURIUserIdentityMetadata              = "/api/users/identity-metadata"
	TestURIUserOnboarding                    = "/api/Spaces-1/useronboarding"
	TestURIUserRoles                         = "/api/userroles{/id}{?skip,take,ids,partialName}"
	TestURIUsers                             = "/api/users{/id}{?skip,take,ids,filter}"
	TestURIVariableNames                     = "/api/Spaces-1/variables/names{?project,runbook,projectEnvironmentsFilter}"
	TestURIVariablePreview                   = "/api/Spaces-1/variables/preview{?project,runbook,environment,channel,tenant,action,machine,role}"
	TestURIVariables                         = "/api/Spaces-1/variables{/id}{?ids}"
	TestURIVersionControlClearCache          = "/api/configuration/versioncontrol/clear-cache"
	TestURIVersionRuleTest                   = "/api/Spaces-1/channels/rule-test{?version,versionRange,preReleaseTag,feetType}"
	TestURIWeb                               = "/app"
	TestURIWorkerOperatingSystems            = "/api/Spaces-1/workers/operatingsystem/names/all"
	TestURIWorkerPools                       = "/api/Spaces-1/workerpools{/id}{?name,skip,ids,take,partialName}"
	TestURIWorkerPoolsDynamicWorkerTypes     = "/api/Spaces-1/workerpools/dynamicworkertypes"
	TestURIWorkerPoolsSortOrder              = "/api/Spaces-1/workerpools/sortorder"
	TestURIWorkerPoolsSummary                = "/api/Spaces-1/workerpools/summary{?ids,partialName,machinePartialName,isDisabled,healthStatuses,commStyles,hideEmptyWorkerPools,shellNames}"
	TestURIWorkerPoolsSupportedTypes         = "/api/Spaces-1/workerpools/supportedtypes"
	TestURIWorkers                           = "/api/Spaces-1/workers{/id}{?skip,take,name,ids,partialName,isDisabled,healthStatuses,commStyles,workerPoolIds,thumbprint,shellNames}"
	TestURIWorkerShells                      = "/api/Spaces-1/workers/operatingsystem/shells/all"
	TestURIWorkerToolsLatestImages           = "/api/workertoolslatestimages"
)