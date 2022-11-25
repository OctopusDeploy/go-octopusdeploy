package client

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/artifacts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/authentication"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/azure"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/azure/devops"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/buildinformation"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/certificates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/cloudtemplate"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/configuration"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/dashboard"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/events"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/externalsecuritygroupproviders"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/interruptions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/invitations"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/issuetrackers"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/jira"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/licenses"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/lifecycles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/migrations"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/octopusservernodes"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/permissions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projectgroups"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/proxies"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/reporting"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/runbooks"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/scheduler"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/serverstatus"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/subscriptions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tagsets"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tasks"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/teammembership"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/teams"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tenants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/triggers"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/useronboarding"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/userroles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/users"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/workerpools"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/workertoolslatestimages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

// Client is an OctopusDeploy for making Octopus API requests.
type Client struct {
	httpSession *newclient.HttpSession
	sling       *sling.Sling // note sling will be removed in v3. The sling instance is wired up to use the same underlying http.Client as the httpSession

	Accounts                       *accounts.AccountService
	ActionTemplates                *actiontemplates.ActionTemplateService
	APIKeys                        *users.ApiKeyService
	Artifacts                      *artifacts.ArtifactService
	Authentication                 *authentication.AuthenticationService
	AzureDevOpsConnectivityCheck   *devops.AzureDevOpsConnectivityCheckService
	AzureEnvironments              *azure.AzureEnvironmentService
	BuildInformation               *buildinformation.BuildInformationService
	CertificateConfiguration       *configuration.CertificateConfigurationService
	Certificates                   *certificates.CertificateService
	Channels                       *channels.ChannelService
	CloudTemplate                  *cloudtemplate.CloudTemplateService
	CommunityActionTemplates       *actions.CommunityActionTemplateService
	Configuration                  *configuration.ConfigurationService
	GitCredentials                 *credentials.Service
	DashboardConfigurations        *dashboard.DashboardConfigurationService
	Dashboards                     *dashboard.DashboardService
	DeploymentProcesses            *deployments.DeploymentProcessService
	Deployments                    *deployments.DeploymentService
	DynamicExtensions              *extensions.DynamicExtensionService
	Environments                   *environments.EnvironmentService
	Events                         *events.EventService
	ExternalSecurityGroupProviders *externalsecuritygroupproviders.ExternalSecurityGroupProviderService
	FeaturesConfiguration          *configuration.FeaturesConfigurationService
	Feeds                          *feeds.FeedService
	Interruptions                  *interruptions.InterruptionService
	Invitations                    *invitations.InvitationService
	IssueTrackers                  *issuetrackers.IssueTrackerService
	JiraIntegration                *jira.JiraIntegrationService
	LetsEncryptConfiguration       *configuration.LetsEncryptConfigurationService
	LibraryVariableSets            *variables.LibraryVariableSetService
	Licenses                       *licenses.LicenseService
	Lifecycles                     *lifecycles.LifecycleService
	MachinePolicies                *machines.MachinePolicyService
	MachineRoles                   *machines.MachineRoleService
	Machines                       *machines.MachineService
	MaintenanceConfiguration       *configuration.MaintenanceConfigurationService
	Migrations                     *migrations.MigrationService
	OctopusPackageMetadata         *packages.OctopusPackageMetadataService
	OctopusServerNodes             *octopusservernodes.OctopusServerNodeService
	Packages                       *packages.PackageService
	PackageMetadata                *packages.PackageMetadataService
	PerformanceConfiguration       *configuration.PerformanceConfigurationService
	Permissions                    *permissions.PermissionService
	ProjectGroups                  *projectgroups.ProjectGroupService
	Projects                       *projects.ProjectService
	ProjectTriggers                *triggers.ProjectTriggerService
	Proxies                        *proxies.ProxyService
	Releases                       *releases.ReleaseService
	Reporting                      *reporting.ReportingService
	RunbookProcesses               *runbooks.RunbookProcessService
	RunbookRuns                    *runbooks.RunbookRunService
	Runbooks                       *runbooks.RunbookService
	RunbookSnapshots               *runbooks.RunbookSnapshotService
	Root                           *RootService
	ScheduledProjectTriggers       *triggers.ScheduledProjectTriggerService
	Scheduler                      *scheduler.SchedulerService
	ScopedUserRoles                *userroles.ScopedUserRoleService
	ScriptModules                  *variables.ScriptModuleService
	ServerConfiguration            *configuration.ServerConfigurationService
	ServerStatus                   *serverstatus.ServerStatusService
	SmtpConfiguration              *configuration.SmtpConfigurationService
	Spaces                         *spaces.SpaceService
	Subscriptions                  *subscriptions.SubscriptionService
	TagSets                        *tagsets.TagSetService
	Tasks                          *tasks.TaskService
	TeamMembership                 *teammembership.TeamMembershipService
	Teams                          *teams.TeamService
	Tenants                        *tenants.TenantService
	TenantVariables                *variables.TenantVariableService
	UpgradeConfiguration           *configuration.UpgradeConfigurationService
	UserOnboarding                 *useronboarding.UserOnboardingService
	UserRoles                      *userroles.UserRoleService
	Users                          *users.UserService
	Variables                      *variables.VariableService
	WorkerPools                    *workerpools.WorkerPoolService
	Workers                        *machines.WorkerService
	WorkerToolsLatestImages        *workertoolslatestimages.WorkerToolsLatestImageService

	// conform to newclient.Client temporarily until this class goes away
	uriTemplateCache *uritemplates.URITemplateCache
}

func IsAPIKey(apiKey string) bool {
	if len(apiKey) < 5 {
		return false
	}

	var expression = regexp.MustCompile(`^(API-)([A-Z\d])+$`)
	return expression.MatchString(apiKey)
}

// NewClient returns a new Octopus API client. If a nil client is provided, a
// new http.Client will be used.
func NewClient(httpClient *http.Client, apiURL *url.URL, apiKey string, spaceID string) (*Client, error) {
	return NewClientForTool(httpClient, apiURL, apiKey, spaceID, "")
}

// NewClientForTool returns a new Octopus API client with a tool reference in the useragent string.
// If a nil client is provided, a new http.Client will be used.
func NewClientForTool(httpClient *http.Client, apiURL *url.URL, apiKey string, spaceID string, requestingTool string) (*Client, error) {
	if apiURL == nil {
		return nil, internal.CreateInvalidParameterError("NewClient", "apiURL")
	}

	if internal.IsEmpty(apiKey) {
		return nil, internal.CreateInvalidParameterError("NewClient", "apiKey")
	}

	if !IsAPIKey(apiKey) {
		return nil, internal.CreateInvalidParameterError("NewClient", "apiKey")
	}

	baseURLWithAPI := strings.TrimRight(apiURL.String(), "/")
	baseURLWithAPI = fmt.Sprintf("%s/api", baseURLWithAPI)

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	// fetch root resource and process paths
	base := sling.New().Client(httpClient).Base(baseURLWithAPI).Set(constants.ClientAPIKeyHTTPHeader, apiKey)
	base.Set("User-Agent", api.GetUserAgentString(requestingTool))
	rootService := NewRootService(base, baseURLWithAPI)

	root, err := rootService.Get()
	if err != nil {
		return nil, err
	}

	// Root with specified Space ID, if it's defined
	sroot := NewRootResource()

	if !internal.IsEmpty(spaceID) {
		baseURLWithAPI = fmt.Sprintf("%s/%s", baseURLWithAPI, spaceID)
		base = sling.New().Client(httpClient).Base(baseURLWithAPI).Set(constants.ClientAPIKeyHTTPHeader, apiKey)
		base.Set("User-Agent", api.GetUserAgentString(requestingTool))
		rootService = NewRootService(base, baseURLWithAPI)
		sroot, err = rootService.Get()

		if err != nil {
			if err == services.ErrItemNotFound {
				return nil, fmt.Errorf("the space ID (%s) cannot be found", spaceID)
			}
			return nil, err
		}
	}

	baseURLWithAPIParsed, fatalErr := url.Parse(baseURLWithAPI)
	if fatalErr != nil { // should never fail because baseURLWithAPI is entirely constructed out of things that are known to be parseable
		panic("failure parsing baseURL " + fatalErr.Error())
	}

	httpSession := &newclient.HttpSession{
		HttpClient: httpClient,
		BaseURL:    baseURLWithAPIParsed,
		DefaultHeaders: map[string]string{
			constants.ClientAPIKeyHTTPHeader: apiKey,
			"User-Agent":                     api.GetUserAgentString(requestingTool),
		},
	}

	rootPath := root.GetLinkPath(sroot, constants.LinkSelf)
	apiKeysPath := "/api/users"
	dynamicExtensionsPath := "/api/dynamic-extensions"
	jiraIntegrationPath := "/api/jiraintegration"
	licensesPath := "/api/licenses"
	migrationsPath := "/api/migrations"
	reportingPath := "/api/reporting"

	accountsPath := root.GetLinkPath(sroot, constants.LinkAccounts)
	actionTemplatesLogo := root.GetLinkPath(sroot, constants.LinkActionTemplateLogo)
	actionTemplatesPath := root.GetLinkPath(sroot, constants.LinkActionTemplates)
	actionTemplatesCategories := root.GetLinkPath(sroot, constants.LinkActionTemplatesCategories)
	actionTemplatesSearch := root.GetLinkPath(sroot, constants.LinkActionTemplatesSearch)
	actionTemplateVersionedLogo := root.GetLinkPath(sroot, constants.LinkActionTemplateVersionedLogo)
	artifactsPath := root.GetLinkPath(sroot, constants.LinkArtifacts)
	authenticateOctopusIDPath := root.GetLinkPath(sroot, constants.LinkAuthenticateOctopusID)
	authenticationPath := root.GetLinkPath(sroot, constants.LinkAuthentication)
	azureDevOpsConnectivityCheckPath := root.GetLinkPath(sroot, constants.LinkAzureDevOpsConnectivityCheck)
	azureEnvironmentsPath := root.GetLinkPath(sroot, constants.LinkAzureEnvironments)
	buildInformationPath := root.GetLinkPath(sroot, constants.LinkBuildInformation)
	buildInformationBulkPath := root.GetLinkPath(sroot, constants.LinkBuildInformationBulk)
	builtInFeedStatsPath := root.GetLinkPath(sroot, constants.LinkBuiltInFeedStats)
	certificateConfigurationPath := root.GetLinkPath(sroot, constants.LinkCertificateConfiguration)
	certificatesPath := root.GetLinkPath(sroot, constants.LinkCertificates)
	channelsPath := root.GetLinkPath(sroot, constants.LinkChannels)
	cloudTemplatePath := root.GetLinkPath(sroot, constants.LinkCloudTemplate)
	communityActionTemplatesPath := root.GetLinkPath(sroot, constants.LinkCommunityActionTemplates)
	configurationPath := root.GetLinkPath(sroot, constants.LinkConfiguration)
	currentLicensePath := root.GetLinkPath(sroot, constants.LinkCurrentLicense)
	currentLicenseStatusPath := root.GetLinkPath(sroot, constants.LinkCurrentLicenseStatus)
	currentUserPath := root.GetLinkPath(sroot, constants.LinkCurrentUser)
	dashboardPath := root.GetLinkPath(sroot, constants.LinkDashboard)
	dashboardConfigurationPath := root.GetLinkPath(sroot, constants.LinkDashboardConfiguration)
	dashboardDynamicPath := root.GetLinkPath(sroot, constants.LinkDashboardDynamic)
	deploymentProcessesPath := root.GetLinkPath(sroot, constants.LinkDeploymentProcesses)
	deploymentsPath := root.GetLinkPath(sroot, constants.LinkDeployments)
	discoverMachinePath := root.GetLinkPath(sroot, constants.LinkDiscoverMachine)
	discoverWorkerPath := root.GetLinkPath(sroot, constants.LinkDiscoverWorker)
	dynamicExtensionsFeaturesMetadataPath := root.GetLinkPath(sroot, constants.LinkDynamicExtensionsFeaturesMetadata)
	dynamicExtensionsFeaturesValuesPath := root.GetLinkPath(sroot, constants.LinkDynamicExtensionsFeaturesValues)
	dynamicExtensionsScriptsPath := root.GetLinkPath(sroot, constants.LinkDynamicExtensionsScripts)
	environmentsPath := root.GetLinkPath(sroot, constants.LinkEnvironments)
	environmentSortOrderPath := root.GetLinkPath(sroot, constants.LinkEnvironmentSortOrder)
	environmentsSummaryPath := root.GetLinkPath(sroot, constants.LinkEnvironmentsSummary)
	eventAgentsPath := root.GetLinkPath(sroot, constants.LinkEventAgents)
	eventCategoriesPath := root.GetLinkPath(sroot, constants.LinkEventCategories)
	eventDocumentTypesPath := root.GetLinkPath(sroot, constants.LinkEventDocumentTypes)
	eventGroupsPath := root.GetLinkPath(sroot, constants.LinkEventGroups)
	eventsPath := root.GetLinkPath(sroot, constants.LinkEvents)
	extensionStatsPath := root.GetLinkPath(sroot, constants.LinkExtensionStats)
	externalSecurityGroupProvidersPath := root.GetLinkPath(sroot, constants.LinkExternalSecurityGroupProviders)
	externalUserSearchPath := root.GetLinkPath(sroot, constants.LinkExternalUserSearch)
	featuresConfigurationPath := root.GetLinkPath(sroot, constants.LinkFeaturesConfiguration)
	feedsPath := root.GetLinkPath(sroot, constants.LinkFeeds)
	gitCredentialsPath := root.GetLinkPath(sroot, constants.LinkGitCredentials)
	interruptionsPath := root.GetLinkPath(sroot, constants.LinkInterruptions)
	invitationsPath := root.GetLinkPath(sroot, constants.LinkInvitations)
	issueTrackersPath := root.GetLinkPath(sroot, constants.LinkIssueTrackers)
	jiraConnectAppCredentialsTestPath := root.GetLinkPath(sroot, constants.LinkJiraConnectAppCredentialsTest)
	jiraCredentialsTestPath := root.GetLinkPath(sroot, constants.LinkJiraCredentialsTest)
	letsEncryptConfigurationPath := root.GetLinkPath(sroot, constants.LinkLetsEncryptConfiguration)
	libraryVariablesPath := root.GetLinkPath(sroot, constants.LinkLibraryVariables)
	lifecyclesPath := root.GetLinkPath(sroot, constants.LinkLifecycles)
	loginInitiatedPath := root.GetLinkPath(sroot, constants.LinkLoginInitiated)
	machineOperatingSystemsPath := root.GetLinkPath(sroot, constants.LinkMachineOperatingSystems)
	machinePoliciesPath := root.GetLinkPath(sroot, constants.LinkMachinePolicies)
	machinePolicyTemplatePath := root.GetLinkPath(sroot, constants.LinkMachinePolicyTemplate)
	machineRolesPath := root.GetLinkPath(sroot, constants.LinkMachineRoles)
	machinesPath := root.GetLinkPath(sroot, constants.LinkMachines)
	machineShellsPath := root.GetLinkPath(sroot, constants.LinkMachineShells)
	maintenanceConfigurationPath := root.GetLinkPath(sroot, constants.LinkMaintenanceConfiguration)
	migrationsImportPath := root.GetLinkPath(sroot, constants.LinkMigrationsImport)
	migrationsPartialExportPath := root.GetLinkPath(sroot, constants.LinkMigrationsPartialExport)
	octopusServerClusterSummaryPath := root.GetLinkPath(sroot, constants.LinkOctopusServerClusterSummary)
	octopusServerNodesPath := root.GetLinkPath(sroot, constants.LinkOctopusServerNodes)
	packageDeltaSignaturePath := root.GetLinkPath(sroot, constants.LinkPackageDeltaSignature)
	packageDeltaUploadPath := root.GetLinkPath(sroot, constants.LinkPackageDeltaUpload)
	packageMetadataPath := root.GetLinkPath(sroot, constants.LinkPackageMetadata)
	packageNotesListPath := root.GetLinkPath(sroot, constants.LinkPackageNotesList)
	packagesPath := root.GetLinkPath(sroot, constants.LinkPackages)
	packagesBulkPath := root.GetLinkPath(sroot, constants.LinkPackagesBulk)
	packageUploadPath := root.GetLinkPath(sroot, constants.LinkPackageUpload)
	performanceConfigurationPath := root.GetLinkPath(sroot, constants.LinkPerformanceConfiguration)
	permissionsPath := root.GetLinkPath(sroot, constants.LinkPermissions)
	projectGroupsPath := root.GetLinkPath(sroot, constants.LinkProjectGroups)
	projectPulsePath := root.GetLinkPath(sroot, constants.LinkProjectPulse)
	projectsExperimentalSummariesPath := root.GetLinkPath(sroot, constants.LinkProjectsExperimentalSummaries)
	projectsExportProjectsPath := root.GetLinkPath(sroot, constants.LinkExportProjects)
	projectsImportProjectsPath := root.GetLinkPath(sroot, constants.LinkImportProjects)
	projectsPath := root.GetLinkPath(sroot, constants.LinkProjects)
	projectTriggersPath := root.GetLinkPath(sroot, constants.LinkProjectTriggers)
	proxiesPath := root.GetLinkPath(sroot, constants.LinkProxies)
	registerPath := root.GetLinkPath(sroot, constants.LinkRegister)
	releasesPath := root.GetLinkPath(sroot, constants.LinkReleases)
	reportingDeploymentsCountedByWeekPath := root.GetLinkPath(sroot, constants.LinkReportingDeploymentsCountedByWeek)
	runbookProcessesPath := root.GetLinkPath(sroot, constants.LinkRunbookProcesses)
	runbookRunsPath := root.GetLinkPath(sroot, constants.LinkRunbookRuns)
	runbooksPath := root.GetLinkPath(sroot, constants.LinkRunbooks)
	runbookSnapshotsPath := root.GetLinkPath(sroot, constants.LinkRunbookSnapshots)
	scheduledProjectTriggersPath := root.GetLinkPath(sroot, constants.LinkScheduledProjectTriggers)
	schedulerPath := root.GetLinkPath(sroot, constants.LinkScheduler)
	scopedUserRolesPath := root.GetLinkPath(sroot, constants.LinkScopedUserRoles)
	serverConfigurationPath := root.GetLinkPath(sroot, constants.LinkServerConfiguration)
	serverConfigurationSettingsPath := root.GetLinkPath(sroot, constants.LinkServerConfigurationSettings)
	serverHealthStatusPath := root.GetLinkPath(sroot, constants.LinkServerHealthStatus)
	serverStatusPath := root.GetLinkPath(sroot, constants.LinkServerStatus)
	signInPath := root.GetLinkPath(sroot, constants.LinkSignIn)
	signOutPath := root.GetLinkPath(sroot, constants.LinkSignOut)
	smtpConfigurationPath := root.GetLinkPath(sroot, constants.LinkSMTPConfiguration)
	smtpIsConfiguredPath := root.GetLinkPath(sroot, constants.LinkSMTPIsConfigured)
	spaceHomePath := root.GetLinkPath(sroot, constants.LinkSpaceHome)
	spacesPath := root.GetLinkPath(sroot, constants.LinkSpaces)
	subscriptionsPath := root.GetLinkPath(sroot, constants.LinkSubscriptions)
	tagSetsPath := root.GetLinkPath(sroot, constants.LinkTagSets)
	tagSetSortOrderPath := root.GetLinkPath(sroot, constants.LinkTagSetSortOrder)
	tasksPath := root.GetLinkPath(sroot, constants.LinkTasks)
	taskTypesPath := root.GetLinkPath(sroot, constants.LinkTaskTypes)
	teamMembershipPath := root.GetLinkPath(sroot, constants.LinkTeamMembership)
	teamMembershipPreviewTeamPath := root.GetLinkPath(sroot, constants.LinkTeamMembershipPreviewTeam)
	teamsPath := root.GetLinkPath(sroot, constants.LinkTeams)
	tenantsPath := root.GetLinkPath(sroot, constants.LinkTenants)
	tenantsMissingVariablesPath := root.GetLinkPath(sroot, constants.LinkTenantsMissingVariables)
	tenantsStatusPath := root.GetLinkPath(sroot, constants.LinkTenantsStatus)
	tenantTagTestPath := root.GetLinkPath(sroot, constants.LinkTenantTagTest)
	tenantVariablesPath := root.GetLinkPath(sroot, constants.LinkTenantVariables)
	timezonesPath := root.GetLinkPath(sroot, constants.LinkTimezones)
	upgradeConfigurationPath := root.GetLinkPath(sroot, constants.LinkUpgradeConfiguration)
	userAuthenticationPath := root.GetLinkPath(sroot, constants.LinkUserAuthentication)
	userIdentityMetadataPath := root.GetLinkPath(sroot, constants.LinkUserIdentityMetadata)
	userOnboardingPath := root.GetLinkPath(sroot, constants.LinkUserOnboarding)
	userRolesPath := root.GetLinkPath(sroot, constants.LinkUserRoles)
	usersPath := root.GetLinkPath(sroot, constants.LinkUsers)
	variableNamesPath := root.GetLinkPath(sroot, constants.LinkVariableNames)
	variablePreviewPath := root.GetLinkPath(sroot, constants.LinkVariablePreview)
	variablesPath := root.GetLinkPath(sroot, constants.LinkVariables)
	versionControlClearCachePath := root.GetLinkPath(sroot, constants.LinkVersionControlClearCache)
	versionRuleTestPath := root.GetLinkPath(sroot, constants.LinkVersionRuleTest)
	workerOperatingSystemsPath := root.GetLinkPath(sroot, constants.LinkWorkerOperatingSystems)
	workerPoolsPath := root.GetLinkPath(sroot, constants.LinkWorkerPools)
	workerPoolsDynamicWorkerTypesPath := root.GetLinkPath(sroot, constants.LinkWorkerPoolsDynamicWorkerTypes)
	workerPoolsSortOrderPath := root.GetLinkPath(sroot, constants.LinkWorkerPoolsSortOrder)
	workerPoolsSummaryPath := root.GetLinkPath(sroot, constants.LinkWorkerPoolsSummary)
	workerPoolsSupportedTypesPath := root.GetLinkPath(sroot, constants.LinkWorkerPoolsSupportedTypes)
	workersPath := root.GetLinkPath(sroot, constants.LinkWorkers)
	workerShellsPath := root.GetLinkPath(sroot, constants.LinkWorkerShells)
	workerToolsLatestImagesPath := root.GetLinkPath(sroot, constants.LinkWorkerToolsLatestImages)

	return &Client{
		httpSession:                    httpSession,
		sling:                          base,
		Accounts:                       accounts.NewAccountService(base, accountsPath),
		ActionTemplates:                actiontemplates.NewActionTemplateService(base, actionTemplatesPath, actionTemplatesCategories, actionTemplatesLogo, actionTemplatesSearch, actionTemplateVersionedLogo),
		APIKeys:                        users.NewAPIKeyService(base, apiKeysPath),
		Artifacts:                      artifacts.NewArtifactService(base, artifactsPath),
		Authentication:                 authentication.NewAuthenticationService(base, authenticationPath, loginInitiatedPath),
		AzureDevOpsConnectivityCheck:   devops.NewAzureDevOpsConnectivityCheckService(base, azureDevOpsConnectivityCheckPath),
		AzureEnvironments:              azure.NewAzureEnvironmentService(base, azureEnvironmentsPath),
		BuildInformation:               buildinformation.NewBuildInformationService(base, buildInformationPath, buildInformationBulkPath),
		CertificateConfiguration:       configuration.NewCertificateConfigurationService(base, certificateConfigurationPath),
		Certificates:                   certificates.NewCertificateService(base, certificatesPath),
		Channels:                       channels.NewChannelService(base, channelsPath, versionRuleTestPath),
		CloudTemplate:                  cloudtemplate.NewCloudTemplateService(base, cloudTemplatePath),
		CommunityActionTemplates:       actions.NewCommunityActionTemplateService(base, communityActionTemplatesPath),
		Configuration:                  configuration.NewConfigurationService(base, configurationPath, versionControlClearCachePath),
		DashboardConfigurations:        dashboard.NewDashboardConfigurationService(base, dashboardConfigurationPath),
		Dashboards:                     dashboard.NewDashboardService(base, dashboardPath, dashboardDynamicPath),
		DeploymentProcesses:            deployments.NewDeploymentProcessService(base, deploymentProcessesPath),
		Deployments:                    deployments.NewDeploymentService(base, deploymentsPath),
		DynamicExtensions:              extensions.NewDynamicExtensionService(base, dynamicExtensionsPath, dynamicExtensionsFeaturesMetadataPath, dynamicExtensionsFeaturesValuesPath, dynamicExtensionsScriptsPath),
		Environments:                   environments.NewEnvironmentService(base, environmentsPath, environmentSortOrderPath, environmentsSummaryPath),
		Events:                         events.NewEventService(base, eventsPath, eventAgentsPath, eventCategoriesPath, eventDocumentTypesPath, eventGroupsPath),
		ExternalSecurityGroupProviders: externalsecuritygroupproviders.NewExternalSecurityGroupProviderService(base, externalSecurityGroupProvidersPath),
		FeaturesConfiguration:          configuration.NewFeaturesConfigurationService(base, featuresConfigurationPath),
		Feeds:                          feeds.NewFeedService(base, feedsPath, builtInFeedStatsPath),
		GitCredentials:                 credentials.NewService(base, gitCredentialsPath),
		Interruptions:                  interruptions.NewInterruptionService(base, interruptionsPath),
		Invitations:                    invitations.NewInvitationService(base, invitationsPath),
		IssueTrackers:                  issuetrackers.NewIssueTrackerService(base, issueTrackersPath),
		JiraIntegration:                jira.NewJiraIntegrationService(base, jiraIntegrationPath, jiraConnectAppCredentialsTestPath, jiraCredentialsTestPath),
		LetsEncryptConfiguration:       configuration.NewLetsEncryptConfigurationService(base, letsEncryptConfigurationPath),
		LibraryVariableSets:            variables.NewLibraryVariableSetService(base, libraryVariablesPath),
		Licenses:                       licenses.NewLicenseService(base, licensesPath, currentLicensePath, currentLicenseStatusPath),
		Lifecycles:                     lifecycles.NewLifecycleService(base, lifecyclesPath),
		MachinePolicies:                machines.NewMachinePolicyService(base, machinePoliciesPath, machinePolicyTemplatePath),
		MachineRoles:                   machines.NewMachineRoleService(base, machineRolesPath),
		Machines:                       machines.NewMachineService(base, machinesPath, discoverMachinePath, machineOperatingSystemsPath, machineShellsPath),
		MaintenanceConfiguration:       configuration.NewMaintenanceConfigurationService(base, maintenanceConfigurationPath),
		Migrations:                     migrations.NewMigrationService(base, migrationsPath, migrationsImportPath, migrationsPartialExportPath),
		OctopusServerNodes:             octopusservernodes.NewOctopusServerNodeService(base, octopusServerNodesPath, octopusServerClusterSummaryPath),
		Packages:                       packages.NewPackageService(base, packagesPath, packageDeltaSignaturePath, packageDeltaUploadPath, packageNotesListPath, packagesBulkPath, packageUploadPath),
		PackageMetadata:                packages.NewPackageMetadataService(base, packageMetadataPath),
		PerformanceConfiguration:       configuration.NewPerformanceConfigurationService(base, performanceConfigurationPath),
		Permissions:                    permissions.NewPermissionService(base, permissionsPath),
		ProjectGroups:                  projectgroups.NewProjectGroupService(base, projectGroupsPath),
		Projects:                       projects.NewProjectService(base, projectsPath, projectPulsePath, projectsExperimentalSummariesPath, projectsImportProjectsPath, projectsExportProjectsPath),
		ProjectTriggers:                triggers.NewProjectTriggerService(base, projectTriggersPath),
		Proxies:                        proxies.NewProxyService(base, proxiesPath),
		Releases:                       releases.NewReleaseService(base, releasesPath),
		Reporting:                      reporting.NewReportingService(base, reportingPath, reportingDeploymentsCountedByWeekPath),
		RunbookProcesses:               runbooks.NewRunbookProcessService(base, runbookProcessesPath),
		RunbookRuns:                    runbooks.NewRunbookRunService(base, runbookRunsPath),
		Runbooks:                       runbooks.NewRunbookService(base, runbooksPath),
		RunbookSnapshots:               runbooks.NewRunbookSnapshotService(base, runbookSnapshotsPath),
		Root:                           NewRootService(base, rootPath),
		Scheduler:                      scheduler.NewSchedulerService(base, schedulerPath),
		ScheduledProjectTriggers:       triggers.NewScheduledProjectTriggerService(base, scheduledProjectTriggersPath),
		ScopedUserRoles:                userroles.NewScopedUserRoleService(base, scopedUserRolesPath),
		ScriptModules:                  variables.NewScriptModuleService(base, libraryVariablesPath),
		ServerConfiguration:            configuration.NewServerConfigurationService(base, serverConfigurationPath, serverConfigurationSettingsPath),
		ServerStatus:                   serverstatus.NewServerStatusService(base, serverStatusPath, extensionStatsPath, serverHealthStatusPath, timezonesPath),
		SmtpConfiguration:              configuration.NewSmtpConfigurationService(base, smtpConfigurationPath, smtpIsConfiguredPath),
		Spaces:                         spaces.NewSpaceService(base, spacesPath, spaceHomePath),
		Subscriptions:                  subscriptions.NewSubscriptionService(base, subscriptionsPath),
		TagSets:                        tagsets.NewTagSetService(base, tagSetsPath, tagSetSortOrderPath),
		Tasks:                          tasks.NewTaskService(base, tasksPath, taskTypesPath),
		TeamMembership:                 teammembership.NewTeamMembershipService(base, teamMembershipPath, teamMembershipPreviewTeamPath),
		Teams:                          teams.NewTeamService(base, teamsPath),
		Tenants:                        tenants.NewTenantService(base, tenantsPath, tenantsMissingVariablesPath, tenantsStatusPath, tenantTagTestPath),
		TenantVariables:                variables.NewTenantVariableService(base, tenantVariablesPath),
		UpgradeConfiguration:           configuration.NewUpgradeConfigurationService(base, upgradeConfigurationPath),
		UserOnboarding:                 useronboarding.NewUserOnboardingService(base, userOnboardingPath),
		UserRoles:                      userroles.NewUserRoleService(base, userRolesPath),
		Users:                          users.NewUserService(base, usersPath, apiKeysPath, authenticateOctopusIDPath, currentUserPath, externalUserSearchPath, registerPath, signInPath, signOutPath, userAuthenticationPath, userIdentityMetadataPath),
		Variables:                      variables.NewVariableService(base, variablesPath, variableNamesPath, variablePreviewPath),
		WorkerPools:                    workerpools.NewWorkerPoolService(base, workerPoolsPath, workerPoolsDynamicWorkerTypesPath, workerPoolsSortOrderPath, workerPoolsSummaryPath, workerPoolsSupportedTypesPath),
		Workers:                        machines.NewWorkerService(base, workersPath, discoverWorkerPath, workerOperatingSystemsPath, workerShellsPath),
		WorkerToolsLatestImages:        workertoolslatestimages.NewWorkerToolsLatestImageService(base, workerToolsLatestImagesPath),

		uriTemplateCache: uritemplates.NewUriTemplateCache(),
	}, nil
}

// confirm to newclient.Client interface for compatibility
func (n *Client) HttpSession() *newclient.HttpSession {
	return n.httpSession
}

func (n *Client) Sling() *sling.Sling {
	return n.sling
}

func (n *Client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.uriTemplateCache
}
