package octopusdeploy

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dghubble/sling"
)

// Client is an OctopusDeploy for making Octopus API requests.
type Client struct {
	sling                          *sling.Sling
	Accounts                       *accountService
	ActionTemplates                *actionTemplateService
	APIKeys                        *apiKeyService
	Artifacts                      *artifactService
	Authentication                 *authenticationService
	AzureDevOpsConnectivityCheck   *azureDevOpsConnectivityCheckService
	AzureEnvironments              *azureEnvironmentService
	BuildInformation               *buildInformationService
	CertificateConfiguration       *certificateConfigurationService
	Certificates                   *certificateService
	Channels                       *channelService
	CloudTemplate                  *cloudTemplateService
	CommunityActionTemplates       *communityActionTemplateService
	Configuration                  *configurationService
	DashboardConfigurations        *dashboardConfigurationService
	Dashboards                     *dashboardService
	DeploymentProcesses            *deploymentProcessService
	Deployments                    *deploymentService
	DynamicExtensions              *dynamicExtensionService
	Environments                   *environmentService
	Events                         *eventService
	ExternalSecurityGroupProviders *externalSecurityGroupProviderService
	FeaturesConfiguration          *featuresConfigurationService
	Feeds                          *feedService
	Interruptions                  *interruptionService
	Invitations                    *invitationService
	IssueTrackers                  *issueTrackerService
	JiraIntegration                *jiraIntegrationService
	LetsEncryptConfiguration       *letsEncryptConfigurationService
	LibraryVariableSets            *libraryVariableSetService
	Licenses                       *licenseService
	Lifecycles                     *lifecycleService
	MachinePolicies                *machinePolicyService
	MachineRoles                   *machineRoleService
	Machines                       *machineService
	MaintenanceConfiguration       *maintenanceConfigurationService
	Migrations                     *migrationService
	OctopusPackageMetadata         *octopusPackageMetadataService
	OctopusServerNodes             *octopusServerNodeService
	Packages                       *packageService
	PackageMetadata                *packageMetadataService
	PerformanceConfiguration       *performanceConfigurationService
	Permissions                    *permissionService
	ProjectGroups                  *projectGroupService
	Projects                       *projectService
	ProjectTriggers                *projectTriggerService
	Proxies                        *proxyService
	Releases                       *releaseService
	Reporting                      *reportingService
	RunbookProcesses               *runbookProcessService
	RunbookRuns                    *runbookRunService
	Runbooks                       *runbookService
	RunbookSnapshots               *runbookSnapshotService
	Root                           *rootService
	ScheduledProjectTriggers       *scheduledProjectTriggerService
	Scheduler                      *schedulerService
	ScopedUserRoles                *scopedUserRoleService
	ServerConfiguration            *serverConfigurationService
	ServerStatus                   *serverStatuService
	SMTPConfiguration              *smtpConfigurationService
	Spaces                         *spaceService
	Subscriptions                  *subscriptionService
	TagSets                        *tagSetService
	Tasks                          *taskService
	TeamMembership                 *teamMembershipService
	Teams                          *teamService
	Tenants                        *tenantService
	TenantVariables                *tenantVariableService
	UpgradeConfiguration           *upgradeConfigurationService
	UserOnboarding                 *userOnboardingService
	UserRoles                      *userRoleService
	Users                          *userService
	Variables                      *variableService
	WorkerPools                    *workerPoolService
	Workers                        *workerService
	WorkerToolsLatestImages        *workerToolsLatestImageService
}

// NewClient returns a new Octopus API client. If a nil client is provided, a
// new http.Client will be used.
func NewClient(httpClient *http.Client, apiURL *url.URL, apiKey string, spaceID string) (*Client, error) {
	if apiURL == nil {
		return nil, createInvalidParameterError(clientNewClient, ParameterOctopusURL)
	}

	if isEmpty(apiKey) {
		return nil, createInvalidParameterError(clientNewClient, ParameterAPIKey)
	}

	if !isAPIKey(apiKey) {
		return nil, createInvalidParameterError(clientNewClient, ParameterAPIKey)
	}

	baseURLWithAPI := strings.TrimRight(apiURL.String(), "/")
	baseURLWithAPI = fmt.Sprintf("%s/api", baseURLWithAPI)

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// fetch root resource and process paths
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set(clientAPIKeyHTTPHeader, apiKey)
	base.Set("User-Agent", "go-octopusdeploy")
	rootService := newRootService(base, baseURLWithAPI)

	root, err := rootService.Get()
	if err != nil {
		return nil, err
	}

	// Root with specified Space ID, if it's defined
	sroot := NewRootResource()

	if !isEmpty(spaceID) {
		baseURLWithAPI = fmt.Sprintf("%s/%s", baseURLWithAPI, spaceID)
		base = sling.New().Client(httpClient).Base(baseURLWithAPI).Set(clientAPIKeyHTTPHeader, apiKey)
		base.Set("User-Agent", "go-octopusdeploy")
		rootService = newRootService(base, baseURLWithAPI)
		sroot, err = rootService.Get()

		if err != nil {
			if err == ErrItemNotFound {
				return nil, fmt.Errorf("the space ID (%s) cannot be found", spaceID)
			}
			return nil, err
		}
	}

	rootPath := root.GetLinkPath(sroot, linkSelf)
	apiKeysPath := "/api/users"
	dynamicExtensionsPath := "/api/dynamic-extensions"
	jiraIntegrationPath := "/api/jiraintegration"
	licensesPath := "/api/licenses"
	migrationsPath := "/api/migrations"
	reportingPath := "/api/reporting"

	accountsPath := root.GetLinkPath(sroot, linkAccounts)
	actionTemplatesLogo := root.GetLinkPath(sroot, linkActionTemplateLogo)
	actionTemplatesPath := root.GetLinkPath(sroot, linkActionTemplates)
	actionTemplatesCategories := root.GetLinkPath(sroot, linkActionTemplatesCategories)
	actionTemplatesSearch := root.GetLinkPath(sroot, linkActionTemplatesSearch)
	actionTemplateVersionedLogo := root.GetLinkPath(sroot, linkActionTemplateVersionedLogo)
	artifactsPath := root.GetLinkPath(sroot, linkArtifacts)
	authenticateOctopusIDPath := root.GetLinkPath(sroot, linkAuthenticateOctopusID)
	authenticationPath := root.GetLinkPath(sroot, linkAuthentication)
	azureDevOpsConnectivityCheckPath := root.GetLinkPath(sroot, linkAzureDevOpsConnectivityCheck)
	azureEnvironmentsPath := root.GetLinkPath(sroot, linkAzureEnvironments)
	buildInformationPath := root.GetLinkPath(sroot, linkBuildInformation)
	buildInformationBulkPath := root.GetLinkPath(sroot, linkBuildInformationBulk)
	builtInFeedStatsPath := root.GetLinkPath(sroot, linkBuiltInFeedStats)
	certificateConfigurationPath := root.GetLinkPath(sroot, linkCertificateConfiguration)
	certificatesPath := root.GetLinkPath(sroot, linkCertificates)
	channelsPath := root.GetLinkPath(sroot, linkChannels)
	cloudTemplatePath := root.GetLinkPath(sroot, linkCloudTemplate)
	communityActionTemplatesPath := root.GetLinkPath(sroot, linkCommunityActionTemplates)
	configurationPath := root.GetLinkPath(sroot, linkConfiguration)
	currentLicensePath := root.GetLinkPath(sroot, linkCurrentLicense)
	currentLicenseStatusPath := root.GetLinkPath(sroot, linkCurrentLicenseStatus)
	currentUserPath := root.GetLinkPath(sroot, linkCurrentUser)
	dashboardPath := root.GetLinkPath(sroot, linkDashboard)
	dashboardConfigurationPath := root.GetLinkPath(sroot, linkDashboardConfiguration)
	dashboardDynamicPath := root.GetLinkPath(sroot, linkDashboardDynamic)
	deploymentProcessesPath := root.GetLinkPath(sroot, linkDeploymentProcesses)
	deploymentsPath := root.GetLinkPath(sroot, linkDeployments)
	discoverMachinePath := root.GetLinkPath(sroot, linkDiscoverMachine)
	discoverWorkerPath := root.GetLinkPath(sroot, linkDiscoverWorker)
	dynamicExtensionsFeaturesMetadataPath := root.GetLinkPath(sroot, linkDynamicExtensionsFeaturesMetadata)
	dynamicExtensionsFeaturesValuesPath := root.GetLinkPath(sroot, linkDynamicExtensionsFeaturesValues)
	dynamicExtensionsScriptsPath := root.GetLinkPath(sroot, linkDynamicExtensionsScripts)
	environmentsPath := root.GetLinkPath(sroot, linkEnvironments)
	environmentSortOrderPath := root.GetLinkPath(sroot, linkEnvironmentSortOrder)
	environmentsSummaryPath := root.GetLinkPath(sroot, linkEnvironmentsSummary)
	eventAgentsPath := root.GetLinkPath(sroot, linkEventAgents)
	eventCategoriesPath := root.GetLinkPath(sroot, linkEventCategories)
	eventDocumentTypesPath := root.GetLinkPath(sroot, linkEventDocumentTypes)
	eventGroupsPath := root.GetLinkPath(sroot, linkEventGroups)
	eventsPath := root.GetLinkPath(sroot, linkEvents)
	extensionStatsPath := root.GetLinkPath(sroot, linkExtensionStats)
	externalSecurityGroupProvidersPath := root.GetLinkPath(sroot, linkExternalSecurityGroupProviders)
	externalUserSearchPath := root.GetLinkPath(sroot, linkExternalUserSearch)
	featuresConfigurationPath := root.GetLinkPath(sroot, linkFeaturesConfiguration)
	feedsPath := root.GetLinkPath(sroot, linkFeeds)
	interruptionsPath := root.GetLinkPath(sroot, linkInterruptions)
	invitationsPath := root.GetLinkPath(sroot, linkInvitations)
	issueTrackersPath := root.GetLinkPath(sroot, linkIssueTrackers)
	jiraConnectAppCredentialsTestPath := root.GetLinkPath(sroot, linkJiraConnectAppCredentialsTest)
	jiraCredentialsTestPath := root.GetLinkPath(sroot, linkJiraCredentialsTest)
	letsEncryptConfigurationPath := root.GetLinkPath(sroot, linkLetsEncryptConfiguration)
	libraryVariablesPath := root.GetLinkPath(sroot, linkLibraryVariables)
	lifecyclesPath := root.GetLinkPath(sroot, linkLifecycles)
	loginInitiatedPath := root.GetLinkPath(sroot, linkLoginInitiated)
	machineOperatingSystemsPath := root.GetLinkPath(sroot, linkMachineOperatingSystems)
	machinePoliciesPath := root.GetLinkPath(sroot, linkMachinePolicies)
	machinePolicyTemplatePath := root.GetLinkPath(sroot, linkMachinePolicyTemplate)
	machineRolesPath := root.GetLinkPath(sroot, linkMachineRoles)
	machinesPath := root.GetLinkPath(sroot, linkMachines)
	machineShellsPath := root.GetLinkPath(sroot, linkMachineShells)
	maintenanceConfigurationPath := root.GetLinkPath(sroot, linkMaintenanceConfiguration)
	migrationsImportPath := root.GetLinkPath(sroot, linkMigrationsImport)
	migrationsPartialExportPath := root.GetLinkPath(sroot, linkMigrationsPartialExport)
	octopusServerClusterSummaryPath := root.GetLinkPath(sroot, linkOctopusServerClusterSummary)
	octopusServerNodesPath := root.GetLinkPath(sroot, linkOctopusServerNodes)
	packageDeltaSignaturePath := root.GetLinkPath(sroot, linkPackageDeltaSignature)
	packageDeltaUploadPath := root.GetLinkPath(sroot, linkPackageDeltaUpload)
	packageMetadataPath := root.GetLinkPath(sroot, linkPackageMetadata)
	packageNotesListPath := root.GetLinkPath(sroot, linkPackageNotesList)
	packagesPath := root.GetLinkPath(sroot, linkPackages)
	packagesBulkPath := root.GetLinkPath(sroot, linkPackagesBulk)
	packageUploadPath := root.GetLinkPath(sroot, linkPackageUpload)
	performanceConfigurationPath := root.GetLinkPath(sroot, linkPerformanceConfiguration)
	permissionsPath := root.GetLinkPath(sroot, linkPermissions)
	projectGroupsPath := root.GetLinkPath(sroot, linkProjectGroups)
	projectPulsePath := root.GetLinkPath(sroot, linkProjectPulse)
	projectsPath := root.GetLinkPath(sroot, linkProjects)
	projectsExperimentalSummariesPath := root.GetLinkPath(sroot, linkProjectsExperimentalSummaries)
	projectTriggersPath := root.GetLinkPath(sroot, linkProjectTriggers)
	proxiesPath := root.GetLinkPath(sroot, linkProxies)
	registerPath := root.GetLinkPath(sroot, linkRegister)
	releasesPath := root.GetLinkPath(sroot, linkReleases)
	reportingDeploymentsCountedByWeekPath := root.GetLinkPath(sroot, linkReportingDeploymentsCountedByWeek)
	runbookProcessesPath := root.GetLinkPath(sroot, linkRunbookProcesses)
	runbookRunsPath := root.GetLinkPath(sroot, linkRunbookRuns)
	runbooksPath := root.GetLinkPath(sroot, linkRunbooks)
	runbookSnapshotsPath := root.GetLinkPath(sroot, linkRunbookSnapshots)
	scheduledProjectTriggersPath := root.GetLinkPath(sroot, linkScheduledProjectTriggers)
	schedulerPath := root.GetLinkPath(sroot, linkScheduler)
	scopedUserRolesPath := root.GetLinkPath(sroot, linkScopedUserRoles)
	serverConfigurationPath := root.GetLinkPath(sroot, linkServerConfiguration)
	serverConfigurationSettingsPath := root.GetLinkPath(sroot, linkServerConfigurationSettings)
	serverHealthStatusPath := root.GetLinkPath(sroot, linkServerHealthStatus)
	serverStatusPath := root.GetLinkPath(sroot, linkServerStatus)
	signInPath := root.GetLinkPath(sroot, linkSignIn)
	signOutPath := root.GetLinkPath(sroot, linkSignOut)
	smtpConfigurationPath := root.GetLinkPath(sroot, linkSMTPConfiguration)
	smtpIsConfiguredPath := root.GetLinkPath(sroot, linkSMTPIsConfigured)
	spaceHomePath := root.GetLinkPath(sroot, linkSpaceHome)
	spacesPath := root.GetLinkPath(sroot, linkSpaces)
	subscriptionsPath := root.GetLinkPath(sroot, linkSubscriptions)
	tagSetsPath := root.GetLinkPath(sroot, linkTagSets)
	tagSetSortOrderPath := root.GetLinkPath(sroot, linkTagSetSortOrder)
	tasksPath := root.GetLinkPath(sroot, linkTasks)
	taskTypesPath := root.GetLinkPath(sroot, linkTaskTypes)
	teamMembershipPath := root.GetLinkPath(sroot, linkTeamMembership)
	teamMembershipPreviewTeamPath := root.GetLinkPath(sroot, linkTeamMembershipPreviewTeam)
	teamsPath := root.GetLinkPath(sroot, linkTeams)
	tenantsPath := root.GetLinkPath(sroot, linkTenants)
	tenantsMissingVariablesPath := root.GetLinkPath(sroot, linkTenantsMissingVariables)
	tenantsStatusPath := root.GetLinkPath(sroot, linkTenantsStatus)
	tenantTagTestPath := root.GetLinkPath(sroot, linkTenantTagTest)
	tenantVariablesPath := root.GetLinkPath(sroot, linkTenantVariables)
	timezonesPath := root.GetLinkPath(sroot, linkTimezones)
	upgradeConfigurationPath := root.GetLinkPath(sroot, linkUpgradeConfiguration)
	userAuthenticationPath := root.GetLinkPath(sroot, linkUserAuthentication)
	userIdentityMetadataPath := root.GetLinkPath(sroot, linkUserIdentityMetadata)
	userOnboardingPath := root.GetLinkPath(sroot, linkUserOnboarding)
	userRolesPath := root.GetLinkPath(sroot, linkUserRoles)
	usersPath := root.GetLinkPath(sroot, linkUsers)
	variableNamesPath := root.GetLinkPath(sroot, linkVariableNames)
	variablePreviewPath := root.GetLinkPath(sroot, linkVariablePreview)
	variablesPath := root.GetLinkPath(sroot, linkVariables)
	versionControlClearCachePath := root.GetLinkPath(sroot, linkVersionControlClearCache)
	versionRuleTestPath := root.GetLinkPath(sroot, linkVersionRuleTest)
	workerOperatingSystemsPath := root.GetLinkPath(sroot, linkWorkerOperatingSystems)
	workerPoolsPath := root.GetLinkPath(sroot, linkWorkerPools)
	workerPoolsDynamicWorkerTypesPath := root.GetLinkPath(sroot, linkWorkerPoolsDynamicWorkerTypes)
	workerPoolsSortOrderPath := root.GetLinkPath(sroot, linkWorkerPoolsSortOrder)
	workerPoolsSummaryPath := root.GetLinkPath(sroot, linkWorkerPoolsSummary)
	workerPoolsSupportedTypesPath := root.GetLinkPath(sroot, linkWorkerPoolsSupportedTypes)
	workersPath := root.GetLinkPath(sroot, linkWorkers)
	workerShellsPath := root.GetLinkPath(sroot, linkWorkerShells)
	workerToolsLatestImagesPath := root.GetLinkPath(sroot, linkWorkerToolsLatestImages)

	return &Client{
		sling:                          base,
		Accounts:                       newAccountService(base, accountsPath),
		ActionTemplates:                newActionTemplateService(base, actionTemplatesPath, actionTemplatesCategories, actionTemplatesLogo, actionTemplatesSearch, actionTemplateVersionedLogo),
		APIKeys:                        newAPIKeyService(base, apiKeysPath),
		Artifacts:                      newArtifactService(base, artifactsPath),
		Authentication:                 newAuthenticationService(base, authenticationPath, loginInitiatedPath),
		AzureDevOpsConnectivityCheck:   newAzureDevOpsConnectivityCheckService(base, azureDevOpsConnectivityCheckPath),
		AzureEnvironments:              newAzureEnvironmentService(base, azureEnvironmentsPath),
		BuildInformation:               newBuildInformationService(base, buildInformationPath, buildInformationBulkPath),
		CertificateConfiguration:       newCertificateConfigurationService(base, certificateConfigurationPath),
		Certificates:                   newCertificateService(base, certificatesPath),
		Channels:                       newChannelService(base, channelsPath, versionRuleTestPath),
		CloudTemplate:                  newCloudTemplateService(base, cloudTemplatePath),
		CommunityActionTemplates:       newCommunityActionTemplateService(base, communityActionTemplatesPath),
		Configuration:                  newConfigurationService(base, configurationPath, versionControlClearCachePath),
		DashboardConfigurations:        newDashboardConfigurationService(base, dashboardConfigurationPath),
		Dashboards:                     newDashboardService(base, dashboardPath, dashboardDynamicPath),
		DeploymentProcesses:            newDeploymentProcessService(base, deploymentProcessesPath),
		Deployments:                    newDeploymentService(base, deploymentsPath),
		DynamicExtensions:              newDynamicExtensionService(base, dynamicExtensionsPath, dynamicExtensionsFeaturesMetadataPath, dynamicExtensionsFeaturesValuesPath, dynamicExtensionsScriptsPath),
		Environments:                   newEnvironmentService(base, environmentsPath, environmentSortOrderPath, environmentsSummaryPath),
		Events:                         newEventService(base, eventsPath, eventAgentsPath, eventCategoriesPath, eventDocumentTypesPath, eventGroupsPath),
		ExternalSecurityGroupProviders: newExternalSecurityGroupProviderService(base, externalSecurityGroupProvidersPath),
		FeaturesConfiguration:          newFeaturesConfigurationService(base, featuresConfigurationPath),
		Feeds:                          newFeedService(base, feedsPath, builtInFeedStatsPath),
		Interruptions:                  newInterruptionService(base, interruptionsPath),
		Invitations:                    newInvitationService(base, invitationsPath),
		IssueTrackers:                  newIssueTrackerService(base, issueTrackersPath),
		JiraIntegration:                newJiraIntegrationService(base, jiraIntegrationPath, jiraConnectAppCredentialsTestPath, jiraCredentialsTestPath),
		LetsEncryptConfiguration:       newLetsEncryptConfigurationService(base, letsEncryptConfigurationPath),
		LibraryVariableSets:            newLibraryVariableSetService(base, libraryVariablesPath),
		Licenses:                       newLicenseService(base, licensesPath, currentLicensePath, currentLicenseStatusPath),
		Lifecycles:                     newLifecycleService(base, lifecyclesPath),
		MachinePolicies:                newMachinePolicyService(base, machinePoliciesPath, machinePolicyTemplatePath),
		MachineRoles:                   newMachineRoleService(base, machineRolesPath),
		Machines:                       newMachineService(base, machinesPath, discoverMachinePath, machineOperatingSystemsPath, machineShellsPath),
		MaintenanceConfiguration:       newMaintenanceConfigurationService(base, maintenanceConfigurationPath),
		Migrations:                     newMigrationService(base, migrationsPath, migrationsImportPath, migrationsPartialExportPath),
		OctopusServerNodes:             newOctopusServerNodeService(base, octopusServerNodesPath, octopusServerClusterSummaryPath),
		Packages:                       newPackageService(base, packagesPath, packageDeltaSignaturePath, packageDeltaUploadPath, packageNotesListPath, packagesBulkPath, packageUploadPath),
		PackageMetadata:                newPackageMetadataService(base, packageMetadataPath),
		PerformanceConfiguration:       newPerformanceConfigurationService(base, performanceConfigurationPath),
		Permissions:                    newPermissionService(base, permissionsPath),
		ProjectGroups:                  newProjectGroupService(base, projectGroupsPath),
		Projects:                       newProjectService(base, projectsPath, projectPulsePath, projectsExperimentalSummariesPath),
		ProjectTriggers:                newProjectTriggerService(base, projectTriggersPath),
		Proxies:                        newProxyService(base, proxiesPath),
		Releases:                       newReleaseService(base, releasesPath),
		Reporting:                      newReportingService(base, reportingPath, reportingDeploymentsCountedByWeekPath),
		RunbookProcesses:               newRunbookProcessService(base, runbookProcessesPath),
		RunbookRuns:                    newRunbookRunService(base, runbookRunsPath),
		Runbooks:                       newRunbookService(base, runbooksPath),
		RunbookSnapshots:               newRunbookSnapshotService(base, runbookSnapshotsPath),
		Root:                           newRootService(base, rootPath),
		Scheduler:                      newSchedulerService(base, schedulerPath),
		ScheduledProjectTriggers:       newScheduledProjectTriggerService(base, scheduledProjectTriggersPath),
		ScopedUserRoles:                newScopedUserRoleService(base, scopedUserRolesPath),
		ServerConfiguration:            newServerConfigurationService(base, serverConfigurationPath, serverConfigurationSettingsPath),
		ServerStatus:                   newServerStatuService(base, serverStatusPath, extensionStatsPath, serverHealthStatusPath, timezonesPath),
		SMTPConfiguration:              newSMTPConfigurationService(base, smtpConfigurationPath, smtpIsConfiguredPath),
		Spaces:                         newSpaceService(base, spacesPath, spaceHomePath),
		Subscriptions:                  newSubscriptionService(base, subscriptionsPath),
		TagSets:                        newTagSetService(base, tagSetsPath, tagSetSortOrderPath),
		Tasks:                          newTaskService(base, tasksPath, taskTypesPath),
		TeamMembership:                 newTeamMembershipService(base, teamMembershipPath, teamMembershipPreviewTeamPath),
		Teams:                          newTeamService(base, teamsPath),
		Tenants:                        newTenantService(base, tenantsPath, tenantsMissingVariablesPath, tenantsStatusPath, tenantTagTestPath),
		TenantVariables:                newTenantVariableService(base, tenantVariablesPath),
		UpgradeConfiguration:           newUpgradeConfigurationService(base, upgradeConfigurationPath),
		UserOnboarding:                 newUserOnboardingService(base, userOnboardingPath),
		UserRoles:                      newUserRoleService(base, userRolesPath),
		Users:                          newUserService(base, usersPath, apiKeysPath, authenticateOctopusIDPath, currentUserPath, externalUserSearchPath, registerPath, signInPath, signOutPath, userAuthenticationPath, userIdentityMetadataPath),
		Variables:                      newVariableService(base, variablesPath, variableNamesPath, variablePreviewPath),
		WorkerPools:                    newWorkerPoolService(base, workerPoolsPath, workerPoolsDynamicWorkerTypesPath, workerPoolsSortOrderPath, workerPoolsSummaryPath, workerPoolsSupportedTypesPath),
		Workers:                        newWorkerService(base, workersPath, discoverWorkerPath, workerOperatingSystemsPath, workerShellsPath),
		WorkerToolsLatestImages:        newWorkerToolsLatestImageService(base, workerToolsLatestImagesPath),
	}, nil
}

// APIError is a generic structure for containing errors for API operations.
type APIError struct {
	Details         string   `json:"Details,omitempty"`
	ErrorMessage    string   `json:"ErrorMessage,omitempty"`
	Errors          []string `json:"Errors,omitempty"`
	FullException   string   `json:"FullException,omitempty"`
	HelpLink        string   `json:"HelpLink,omitempty"`
	HelpText        string   `json:"HelpText,omitempty"`
	ParsedHelpLinks []string `json:"ParsedHelpLinks,omitempty"`
	StatusCode      int
}

// Error creates a predefined error for Octopus API responses.
func (e APIError) Error() string {
	return fmt.Sprintf("Octopus API error: %v %+v %v", e.ErrorMessage, e.Errors, e.FullException)
}

// APIErrorChecker is a generic error handler for the OctopusDeploy API.
func APIErrorChecker(urlPath string, resp *http.Response, wantedResponseCode int, slingError error, octopusDeployError *APIError) error {
	if octopusDeployError.Errors != nil {
		return fmt.Errorf("octopus deploy api returned an error on endpoint %s - %s", urlPath, octopusDeployError.Errors)
	}

	if slingError != nil {
		return fmt.Errorf("cannot get endpoint %s from server. failure from http client %v", urlPath, slingError)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		octopusDeployError.StatusCode = resp.StatusCode
		return octopusDeployError
	}

	if resp.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("bad request from endpoint %s. response from server %s", urlPath, resp.Status)
	}

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	if resp.StatusCode != wantedResponseCode {
		return octopusDeployError
	}

	return nil
}

// LoadNextPage checks if the next page should be loaded from the API. Returns
// the new path and a bool if the next page should be checked.
func LoadNextPage(pagedResults PagedResults) (string, bool) {
	if pagedResults.Links.PageNext != emptyString {
		return pagedResults.Links.PageNext, true
	}

	return emptyString, false
}

// Generic OctopusDeploy API Get Function.
func apiGet(sling *sling.Sling, inputStruct interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(OperationAPIGet, ParameterSling)
	}

	getClient := sling.New()
	if getClient == nil {
		return nil, createClientInitializationError(OperationAPIGet)
	}

	getClient = getClient.Get(path)
	if getClient == nil {
		return nil, createClientInitializationError(OperationAPIGet)
	}

	getClient.Set("User-Agent", "go-octopusdeploy")

	octopusDeployError := new(APIError)
	resp, err := getClient.Receive(inputStruct, &octopusDeployError)
	if err != nil {
		return nil, err
	}

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return inputStruct, nil
}

// Generic OctopusDeploy API Add Function. Expects a 201 response.
func apiAdd(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(OperationAPIAdd, ParameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(OperationAPIAdd, ParameterPath)
	}

	postClient := sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	postClient.Set("User-Agent", "go-octopusdeploy")

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIAdd)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusCreated, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// apiPost post to octopus and expect a 200 response code.
func apiPost(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(OperationAPIPost, ParameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(OperationAPIPost, ParameterPath)
	}

	postClient := sling.New()
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIPost)
	}

	postClient = postClient.Post(path)
	if postClient == nil {
		return nil, createClientInitializationError(OperationAPIPost)
	}

	postClient.Set("User-Agent", "go-octopusdeploy")

	request := postClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIPost)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Generic OctopusDeploy API Update Function.
func apiUpdate(sling *sling.Sling, inputStruct interface{}, resource interface{}, path string) (interface{}, error) {
	if sling == nil {
		return nil, createInvalidParameterError(OperationAPIUpdate, ParameterSling)
	}

	if isEmpty(path) {
		return nil, createInvalidParameterError(OperationAPIUpdate, ParameterPath)
	}

	putClient := sling.New()
	if putClient == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	putClient = putClient.Put(path)
	if putClient == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	putClient.Set("User-Agent", "go-octopusdeploy")

	request := putClient.BodyJSON(inputStruct)
	if request == nil {
		return nil, createClientInitializationError(OperationAPIUpdate)
	}

	octopusDeployError := new(APIError)
	resp, err := request.Receive(resource, &octopusDeployError)

	apiErrorCheck := APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
	if apiErrorCheck != nil {
		return nil, apiErrorCheck
	}

	return resource, nil
}

// Generic OctopusDeploy API Delete Function.
func apiDelete(sling *sling.Sling, path string) error {
	if sling == nil {
		return createInvalidParameterError(OperationAPIDelete, ParameterSling)
	}

	if isEmpty(path) {
		return createInvalidParameterError(OperationAPIDelete, ParameterPath)
	}

	deleteClient := sling.New()
	if deleteClient == nil {
		return createClientInitializationError(OperationAPIDelete)
	}

	deleteClient = deleteClient.Delete(path)
	if deleteClient == nil {
		return createClientInitializationError(OperationAPIDelete)
	}

	deleteClient.Set("User-Agent", "go-octopusdeploy")

	octopusDeployError := new(APIError)
	resp, err := deleteClient.Receive(nil, &octopusDeployError)

	return APIErrorChecker(path, resp, http.StatusOK, err, octopusDeployError)
}

// ErrItemNotFound is an OctopusDeploy error returned an item cannot be found.
var ErrItemNotFound = errors.New("cannot find the item")
