package octopusdeploy

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType string   `structs:"accountType,omitempty"`
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type ActionTemplateLogoQuery struct {
	CB       string `structs:"cb,omitempty"`
	TypeOrID string `structs:"typeOrId,omitempty"`
}

// ActionTemplatesQuery represents parameters to query the ActionTemplates service.
type ActionTemplatesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type ActionTemplateVersionedLogoQuery struct {
	CB       string `structs:"cb,omitempty"`
	TypeOrID string `structs:"typeOrId,omitempty"`
	Version  string `structs:"version,omitempty"`
}

type ArtifactsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	Order       string   `structs:"order,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Regarding   string   `structs:"regarding,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type BuildInformationQuery struct {
	Filter        string `structs:"filter,omitempty"`
	Latest        string `structs:"latest,omitempty"`
	OverwriteMode string `structs:"overwriteMode,omitempty"`
	PackageID     string `structs:"packageId,omitempty"`
	Skip          int    `structs:"skip,omitempty"`
	Take          int    `structs:"take,omitempty"`
}

type BuildInformationBulkQuery struct {
	IDs []string `structs:"ids,omitempty"`
}

type CertificateConfigurationQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type CertificatesQuery struct {
	Archived    string   `structs:"archived,omitempty"`
	FirstResult string   `structs:"firstResult,omitempty"`
	IDs         []string `structs:"ids,omitempty"`
	OrderBy     string   `structs:"orderBy,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Search      string   `structs:"search,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
	Tenant      string   `structs:"tenant,omitempty"`
}

type ChannelsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type CloudTemplateQuery struct {
	FeedID    string `structs:"feedId,omitempty"`
	PackageID string `structs:"packageId,omitempty"`
}

type CommunityActionTemplatesQuery struct {
	IDs  []string `structs:"ids,omitempty"`
	Skip int      `structs:"skip,omitempty"`
	Take int      `structs:"take,omitempty"`
}

type DashboardQuery struct {
	HighestLatestVersionPerProjectAndEnvironment bool     `structs:"highestLatestVersionPerProjectAndEnvironment,omitempty"`
	ProjectID                                    string   `structs:"projectId,omitempty"`
	SelectedTags                                 []string `structs:"selectedTags,omitempty"`
	SelectedTenants                              []string `structs:"selectedTenants,omitempty"`
	ShowAll                                      bool     `structs:"showAll,omitempty"`
	ReleaseID                                    string   `structs:"releaseId,omitempty"`
}

type DashboardDynamicQuery struct {
	Environments    []string `structs:"environments,omitempty"`
	IncludePrevious bool     `structs:"includePrevious,omitempty"`
	Projects        []string `structs:"projects,omitempty"`
}

type DeploymentProcessesQuery struct {
	IDs  []string `structs:"ids,omitempty"`
	Skip int      `structs:"skip,omitempty"`
	Take int      `structs:"take,omitempty"`
}

type DeploymentsQuery struct {
	Channels     string   `structs:"channels,omitempty"`
	Environments []string `structs:"environments,omitempty"`
	IDs          []string `structs:"ids,omitempty"`
	PartialName  string   `structs:"partialName,omitempty"`
	Projects     []string `structs:"projects,omitempty"`
	Skip         int      `structs:"skip,omitempty"`
	Take         int      `structs:"take,omitempty"`
	TaskState    string   `structs:"taskState,omitempty"`
	Tenants      []string `structs:"tenants,omitempty"`
}

type DiscoverMachineQuery struct {
	Host    string `structs:"host,omitempty"`
	Port    int    `structs:"port,omitempty"`
	ProxyID string `structs:"proxyId,omitempty"`
	Type    string `structs:"type,omitempty"`
}

type DiscoverWorkerQuery struct {
	Host    string `structs:"host,omitempty"`
	Port    int    `structs:"port,omitempty"`
	ProxyID string `structs:"proxyId,omitempty"`
	Type    string `structs:"type,omitempty"`
}

type EnvironmentsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	Name        string   `structs:"name,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `structs:"commStyles,omitempty"`
	HealthStatuses        []string `structs:"healthStatuses,omitempty"`
	HideEmptyEnvironments bool     `structs:"hideEmptyEnvironments,omitempty"`
	IDs                   []string `structs:"ids,omitempty"`
	IsDisabled            bool     `structs:"isDisabled,omitempty"`
	MachinePartialName    string   `structs:"machinePartialName,omitempty"`
	PartialName           string   `structs:"partialName,omitempty"`
	Roles                 []string `structs:"roles,omitempty"`
	ShellNames            []string `structs:"shellNames,omitempty"`
	TenantIDs             []string `structs:"tenantIds,omitempty"`
	TenantTags            []string `structs:"tenantTags,omitempty"`
}

type EventCategoriesQuery struct {
	AppliesTo string `structs:"appliesTo,omitempty"`
}

type EventGroupsQuery struct {
	AppliesTo string `structs:"appliesTo,omitempty"`
}

type EventsQuery struct {
	AsCSV             string   `structs:"asCsv,omitempty"`
	DocumentTypes     []string `structs:"documentTypes,omitempty"`
	Environments      []string `structs:"environments,omitempty"`
	EventAgents       []string `structs:"eventAgents,omitempty"`
	EventCategories   []string `structs:"eventCategories,omitempty"`
	EventGroups       []string `structs:"eventGroups,omitempty"`
	ExcludeDifference bool     `structs:"excludeDifference,omitempty"`
	From              string   `structs:"from,omitempty"`
	FromAutoID        string   `structs:"fromAutoId,omitempty"`
	IDs               []string `structs:"ids,omitempty"`
	IncludeSystem     bool     `structs:"includeSystem,omitempty"`
	Internal          string   `structs:"interal,omitempty"`
	Name              string   `structs:"name,omitempty"`
	PartialName       string   `structs:"partialName,omitempty"`
	ProjectGroups     []string `structs:"projectGroups,omitempty"`
	Projects          []string `structs:"projects,omitempty"`
	Regarding         string   `structs:"regarding,omitempty"`
	RegardingAny      string   `structs:"regardingAny,omitempty"`
	Skip              int      `structs:"skip,omitempty"`
	Spaces            []string `structs:"spaces,omitempty"`
	Take              int      `structs:"take,omitempty"`
	Tags              []string `structs:"tags,omitempty"`
	Tenants           []string `structs:"tenants,omitempty"`
	To                string   `structs:"to,omitempty"`
	ToAutoID          string   `structs:"toAutoId,omitempty"`
	User              string   `structs:"user,omitempty"`
	Users             []string `structs:"users,omitempty"`
}

type ExternalUserSearchQuery struct {
	PartialName string `structs:"partialName,omitempty"`
}

type FeedsQuery struct {
	FeedType    string   `structs:"feedType,omitempty"`
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type InterruptionsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PendingOnly bool     `structs:"pendingOnly,omitempty"`
	Regarding   string   `structs:"regarding,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type IssueTrackersQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type LibraryVariablesQuery struct {
	ContentType string   `structs:"contentType,omitempty"`
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type LifecyclesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type MachinePoliciesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type MachinesQuery struct {
	CommunicationStyles []string `structs:"commStyles,omitempty"`
	DeploymentID        string   `structs:"deploymentId,omitempty"`
	EnvironmentIDs      []string `structs:"environmentIds,omitempty"`
	HealthStatuses      []string `structs:"healthStatuses,omitempty"`
	IDs                 []string `structs:"ids,omitempty"`
	IsDisabled          bool     `structs:"isDisabled,omitempty"`
	Name                string   `structs:"name,omitempty"`
	PartialName         string   `structs:"partialName,omitempty"`
	Roles               []string `structs:"roles,omitempty"`
	ShellNames          []string `structs:"shellNames,omitempty"`
	Skip                int      `structs:"skip,omitempty"`
	Take                int      `structs:"take,omitempty"`
	TenantIDs           []string `structs:"tenantIds,omitempty"`
	TenantTags          []string `structs:"tenantTags,omitempty"`
	Thumbprint          string   `structs:"thumbprint,omitempty"`
}

type OctopusServerNodesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type PackageDeltaSignatureQuery struct {
	PackageID string `structs:"packageId,omitempty"`
	Version   string `structs:"version,omitempty"`
}

type PackageDeltaUploadQuery struct {
	BaseVersion   string `structs:"baseVersion,omitempty"`
	OverwriteMode string `structs:"overwriteMode,omitempty"`
	PackageID     string `structs:"packageId,omitempty"`
	Replace       bool   `structs:"replace,omitempty"`
}

type PackageMetadataQuery struct {
	Filter        string `structs:"filter,omitempty"`
	Latest        string `structs:"latest,omitempty"`
	OverwriteMode string `structs:"overwriteMode,omitempty"`
	Replace       bool   `structs:"replace,omitempty"`
	Skip          int    `structs:"skip,omitempty"`
	Take          int    `structs:"take,omitempty"`
}

type PackageNotesListQuery struct {
	PackageIDs []string `structs:"packageIds,omitempty"`
}

type PackagesQuery struct {
	Filter         string `structs:"filter,omitempty"`
	IncludeNotes   bool   `structs:"includeNotes,omitempty"`
	Latest         string `structs:"latest,omitempty"`
	NuGetPackageID string `structs:"nuGetPackageId,omitempty"`
	Skip           int    `structs:"skip,omitempty"`
	Take           int    `structs:"take,omitempty"`
}

type PackagesBulkQuery struct {
	IDs []string `structs:"ids,omitempty"`
}

type PackageUploadQuery struct {
	Replace       bool   `structs:"replace,omitempty"`
	OverwriteMode string `structs:"overwriteMode,omitempty"`
}

type ProjectGroupsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type ProjectPulseQuery struct {
	ProjectIDs []string `structs:"projectIds,omitempty"`
}

type ProjectsQuery struct {
	Clone               bool     `structs:"clone,omitempty"`
	ClonedFromProjectID string   `structs:"clonedFromProjectId,omitempty"`
	IDs                 []string `structs:"ids,omitempty"`
	Name                string   `structs:"name,omitempty"`
	PartialName         string   `structs:"partialName,omitempty"`
	Skip                int      `structs:"skip,omitempty"`
	Take                int      `structs:"take,omitempty"`
}

type ProjectsExperimentalSummariesQuery struct {
	IDs []string `structs:"ids,omitempty"`
}

type ProjectTriggersQuery struct {
	IDs      []string `structs:"ids,omitempty"`
	Runbooks []string `structs:"runbooks,omitempty"`
	Skip     int      `structs:"skip,omitempty"`
	Take     int      `structs:"take,omitempty"`
}

type ProxiesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type ReleasesQuery struct {
	IDs                []string `structs:"ids,omitempty"`
	IgnoreChannelRules bool     `structs:"ignoreChannelRules,omitempty"`
	Skip               int      `structs:"skip,omitempty"`
	Take               int      `structs:"take,omitempty"`
}

type RunbookProcessesQuery struct {
	IDs  []string `structs:"ids,omitempty"`
	Skip int      `structs:"skip,omitempty"`
	Take int      `structs:"take,omitempty"`
}

type RunbookRunsQuery struct {
	Environments []string `structs:"environments,omitempty"`
	IDs          []string `structs:"ids,omitempty"`
	PartialName  string   `structs:"partialName,omitempty"`
	Projects     []string `structs:"projects,omitempty"`
	Runbooks     []string `structs:"runbooks,omitempty"`
	Skip         int      `structs:"skip,omitempty"`
	Take         int      `structs:"take,omitempty"`
	TaskState    string   `structs:"taskState,omitempty"`
	Tenants      []string `structs:"tenants,omitempty"`
}

type RunbooksQuery struct {
	Clone       bool     `structs:"clone,omitempty"`
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	ProjectIDs  []string `structs:"projectIds,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type RunbookSnapshotsQuery struct {
	IDs     []string `structs:"ids,omitempty"`
	Publish bool     `structs:"publish,omitempty"`
	Skip    int      `structs:"skip,omitempty"`
	Take    int      `structs:"take,omitempty"`
}

type ScheduledProjectTriggersQuery struct {
	IDs  []string `structs:"ids,omitempty"`
	Skip int      `structs:"skip,omitempty"`
	Take int      `structs:"take,omitempty"`
}

type SchedulerQuery struct {
	Verbose bool   `structs:"verbose,omitempty"`
	Tail    string `structs:"tail,omitempty"`
}

type ScopedUserRolesQuery struct {
	IDs           []string `structs:"ids,omitempty"`
	IncludeSystem bool     `structs:"includeSystem,omitempty"`
	PartialName   string   `structs:"partialName,omitempty"`
	Skip          int      `structs:"skip,omitempty"`
	Spaces        []string `structs:"spaces,omitempty"`
	Take          int      `structs:"take,omitempty"`
}

type SignInQuery struct {
	ReturnURL string `structs:"returnUrl,omitempty"`
}

type SpaceHomeQuery struct {
	SpaceID string `structs:"spaceId,omitempty"`
}

type SpacesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	Name        string   `structs:"name,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type SubscriptionsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Spaces      []string `structs:"spaces,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type TagSetsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type TasksQuery struct {
	Active                  bool     `structs:"active,omityempty"`
	Environment             string   `structs:"environment,omitempty"`
	HasPendingInterruptions bool     `structs:"hasPendingInterruptions,omityempty"`
	HasWarningsOrErrors     bool     `structs:"hasWarningsOrErrors,omityempty"`
	IDs                     []string `structs:"ids,omitempty"`
	IncludeSystem           bool     `structs:"includeSystem,omitempty"`
	Name                    string   `structs:"name,omitempty"`
	Node                    string   `structs:"node,omitempty"`
	PartialName             string   `structs:"partialName,omitempty"`
	Project                 string   `structs:"project,omitempty"`
	Runbook                 string   `structs:"runbook,omitempty"`
	Running                 bool     `structs:"running,omityempty"`
	Skip                    int      `structs:"skip,omitempty"`
	Spaces                  []string `structs:"spaces,omitempty"`
	States                  []string `structs:"states,omitempty"`
	Take                    int      `structs:"take,omitempty"`
	Tenant                  string   `structs:"tenant,omitempty"`
}

type TeamMembershipQuery struct {
	IncludeSystem bool     `structs:"includeSystem,omitempty"`
	Spaces        []string `structs:"spaces,omitempty"`
	UserID        string   `structs:"userId,omitempty"`
}

type TeamsQuery struct {
	IDs           []string `structs:"ids,omitempty"`
	IncludeSystem bool     `structs:"includeSystem,omitempty"`
	PartialName   string   `structs:"partialName,omitempty"`
	Skip          int      `structs:"skip,omitempty"`
	Spaces        []string `structs:"spaces,omitempty"`
	Take          int      `structs:"take,omitempty"`
}

type TenantsQuery struct {
	Clone              bool     `structs:"clone,omitempty"`
	ClonedFromTenantID string   `structs:"clonedFromTenantId,omitempty"`
	IDs                []string `structs:"ids,omitempty"`
	Name               string   `structs:"name,omitempty"`
	PartialName        string   `structs:"partialName,omitempty"`
	ProjectID          string   `structs:"projectId,omitempty"`
	Skip               int      `structs:"skip,omitempty"`
	Tags               []string `structs:"tags,omitempty"`
	Take               int      `structs:"take,omitempty"`
}

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `structs:"environmentId,omitempty"`
	IncludeDetails bool     `structs:"includeDetails,omitempty"`
	ProjectID      string   `structs:"projectId,omitempty"`
	TenantID       string   `structs:"tenantId,omitempty"`
}

type TenantVariablesQuery struct {
	ProjectID string `structs:"projectId,omitempty"`
}

type UserRolesQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type UsersQuery struct {
	Filter string   `structs:"filter,omitempty"`
	IDs    []string `structs:"ids,omitempty"`
	Skip   int      `structs:"skip,omitempty"`
	Take   int      `structs:"take,omitempty"`
}

type VariableNamesQuery struct {
	Project                   string `structs:"project,omitempty"`
	ProjectEnvironmentsFilter string `struct:"projectEnvironmentsFilter,omitempty"`
	Runbook                   string `structs:"runbook,omitempty"`
}

type VariablePreviewQuery struct {
	Action      string `structs:"action,omitempty"`
	Channel     string `structs:"channel,omitempty"`
	Environment string `structs:"environment,omitempty"`
	Machine     string `structs:"machine,omitempty"`
	Project     string `structs:"project,omitempty"`
	Role        string `structs:"role,omitempty"`
	Runbook     string `structs:"runbook,omitempty"`
	Tenant      string `structs:"tenant,omitempty"`
}

type VariablesQuery struct {
	IDs []string `structs:"ids,omitempty"`
}

type VersionRuleTestQuery struct {
	FeetType      string `structs:"feetType,omitempty"`
	PreReleaseTag string `structs:"preReleaseTag,omitempty"`
	Version       string `structs:"version,omitempty"`
	VersionRange  string `structs:"versionRange,omitempty"`
}

type WorkerPoolsQuery struct {
	IDs         []string `structs:"ids,omitempty"`
	Name        string   `structs:"name,omitempty"`
	PartialName string   `structs:"partialName,omitempty"`
	Skip        int      `structs:"skip,omitempty"`
	Take        int      `structs:"take,omitempty"`
}

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `structs:"commStyles,omitempty"`
	HealthStatuses       []string `structs:"healthStatuses,omitempty"`
	HideEmptyWorkerPools bool     `structs:"hideEmptyWorkerPools,omitempty"`
	IDs                  []string `structs:"ids,omitempty"`
	IsDisabled           bool     `structs:"isDisabled,omitempty"`
	MachinePartialName   string   `structs:"machinePartialName,omitempty"`
	PartialName          string   `structs:"partialName,omitempty"`
	ShellNames           []string `structs:"shellNames,omitempty"`
}

type WorkersQuery struct {
	CommunicationStyles []string `structs:"commStyles,omitempty"`
	HealthStatuses      []string `structs:"healthStatuses,omitempty"`
	IDs                 []string `structs:"ids,omitempty"`
	IsDisabled          bool     `structs:"isDisabled,omitempty"`
	Name                string   `structs:"name,omitempty"`
	PartialName         string   `structs:"partialName,omitempty"`
	ShellNames          []string `structs:"shellNames,omitempty"`
	Skip                int      `structs:"skip,omitempty"`
	Take                int      `structs:"take,omitempty"`
	Thumbprint          string   `structs:"thumbprint,omitempty"`
	WorkerPoolIDs       []string `structs:"workerPoolIds,omitempty"`
}
