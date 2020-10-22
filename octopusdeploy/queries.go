package octopusdeploy

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType string   `url:"accountType,omitempty"`
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type ActionTemplateLogoQuery struct {
	CB       string `url:"cb,omitempty"`
	TypeOrID string `url:"typeOrId,omitempty"`
}

// ActionTemplatesQuery represents parameters to query the ActionTemplates service.
type ActionTemplatesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type ActionTemplateVersionedLogoQuery struct {
	CB       string `url:"cb,omitempty"`
	TypeOrID string `url:"typeOrId,omitempty"`
	Version  string `url:"version,omitempty"`
}

type ArtifactsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	Order       string   `url:"order,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Regarding   string   `url:"regarding,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type BuildInformationQuery struct {
	Filter        string `url:"filter,omitempty"`
	Latest        string `url:"latest,omitempty"`
	OverwriteMode string `url:"overwriteMode,omitempty"`
	PackageID     string `url:"packageId,omitempty"`
	Skip          int    `url:"skip,omitempty"`
	Take          int    `url:"take,omitempty"`
}

type BuildInformationBulkQuery struct {
	IDs []string `url:"ids,omitempty,comma"`
}

type CertificateConfigurationQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type CertificatesQuery struct {
	Archived    string   `url:"archived,omitempty"`
	FirstResult string   `url:"firstResult,omitempty"`
	IDs         []string `url:"ids,omitempty,comma"`
	OrderBy     string   `url:"orderBy,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Search      string   `url:"search,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
	Tenant      string   `url:"tenant,omitempty"`
}

type ChannelsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type CloudTemplateQuery struct {
	FeedID    string `url:"feedId,omitempty"`
	PackageID string `url:"packageId,omitempty"`
}

type CommunityActionTemplatesQuery struct {
	IDs  []string `url:"ids,omitempty,comma"`
	Skip int      `url:"skip,omitempty"`
	Take int      `url:"take,omitempty"`
}

type DashboardQuery struct {
	IncludeLatest   bool     `url:"highestLatestVersionPerProjectAndEnvironment,omitempty"`
	ProjectID       string   `url:"projectId,omitempty"`
	SelectedTags    []string `url:"selectedTags,omitempty"`
	SelectedTenants []string `url:"selectedTenants,omitempty"`
	ShowAll         bool     `url:"showAll,omitempty"`
	ReleaseID       string   `url:"releaseId,omitempty"`
}

type DashboardDynamicQuery struct {
	Environments    []string `url:"environments,comma"`
	IncludePrevious bool     `url:"includePrevious,omitempty"`
	Projects        []string `url:"projects,comma"`
}

type DeploymentProcessesQuery struct {
	IDs  []string `url:"ids,omitempty,comma"`
	Skip int      `url:"skip,omitempty"`
	Take int      `url:"take,omitempty"`
}

type DeploymentsQuery struct {
	Channels     string   `url:"channels,omitempty"`
	Environments []string `url:"environments,comma"`
	IDs          []string `url:"ids,omitempty,comma"`
	PartialName  string   `url:"partialName,omitempty"`
	Projects     []string `url:"projects,comma"`
	Skip         int      `url:"skip,omitempty"`
	Take         int      `url:"take,omitempty"`
	TaskState    string   `url:"taskState,omitempty"`
	Tenants      []string `url:"tenants,comma"`
}

type DiscoverMachineQuery struct {
	Host    string `url:"host,omitempty"`
	Port    int    `url:"port,omitempty"`
	ProxyID string `url:"proxyId,omitempty"`
	Type    string `url:"type,omitempty"`
}

type DiscoverWorkerQuery struct {
	Host    string `url:"host,omitempty"`
	Port    int    `url:"port,omitempty"`
	ProxyID string `url:"proxyId,omitempty"`
	Type    string `url:"type,omitempty"`
}

type EnvironmentsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	Name        string   `url:"name,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `url:"commStyles,comma"`
	HealthStatuses        []string `url:"healthStatuses,comma"`
	HideEmptyEnvironments bool     `url:"hideEmptyEnvironments,omitempty"`
	IDs                   []string `url:"ids,omitempty,comma"`
	IsDisabled            bool     `url:"isDisabled,omitempty"`
	MachinePartialName    string   `url:"machinePartialName,omitempty"`
	PartialName           string   `url:"partialName,omitempty"`
	Roles                 []string `url:"roles,comma"`
	ShellNames            []string `url:"shellNames,comma"`
	TenantIDs             []string `url:"tenantIds,comma"`
	TenantTags            []string `url:"tenantTags,comma"`
}

type EventCategoriesQuery struct {
	AppliesTo string `url:"appliesTo,omitempty"`
}

type EventGroupsQuery struct {
	AppliesTo string `url:"appliesTo,omitempty"`
}

type EventsQuery struct {
	AsCSV             string   `url:"asCsv,omitempty"`
	DocumentTypes     []string `url:"documentTypes,comma"`
	Environments      []string `url:"environments,comma"`
	EventAgents       []string `url:"eventAgents,comma"`
	EventCategories   []string `url:"eventCategories,comma"`
	EventGroups       []string `url:"eventGroups,comma"`
	ExcludeDifference bool     `url:"excludeDifference"`
	From              string   `url:"from,omitempty"`
	FromAutoID        string   `url:"fromAutoId,omitempty"`
	IDs               []string `url:"ids,omitempty,comma"`
	IncludeSystem     bool     `url:"includeSystem"`
	Internal          string   `url:"interal,omitempty"`
	Name              string   `url:"name,omitempty"`
	PartialName       string   `url:"partialName,omitempty"`
	ProjectGroups     []string `url:"projectGroups,comma"`
	Projects          []string `url:"projects,comma"`
	Regarding         string   `url:"regarding,omitempty"`
	RegardingAny      string   `url:"regardingAny,omitempty"`
	Skip              int      `url:"skip,omitempty"`
	Spaces            []string `url:"spaces,comma"`
	Take              int      `url:"take,omitempty"`
	Tags              []string `url:"tags,comma"`
	Tenants           []string `url:"tenants,comma"`
	To                string   `url:"to,omitempty"`
	ToAutoID          string   `url:"toAutoId,omitempty"`
	User              string   `url:"user,omitempty"`
	Users             []string `url:"users,omitempty"`
}

type ExternalUserSearchQuery struct {
	PartialName string `url:"partialName,omitempty"`
}

type FeedsQuery struct {
	FeedType    string   `url:"feedType,omitempty"`
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type InterruptionsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PendingOnly bool     `url:"pendingOnly,omitempty"`
	Regarding   string   `url:"regarding,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type IssueTrackersQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type LibraryVariablesQuery struct {
	ContentType string   `url:"contentType,omitempty"`
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type LifecyclesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type MachinePoliciesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type MachinesQuery struct {
	CommunicationStyles []string `url:"commStyles,comma"`
	DeploymentID        string   `url:"deploymentId,omitempty"`
	EnvironmentIDs      []string `url:"environmentIds,comma"`
	HealthStatuses      []string `url:"healthStatuses,comma"`
	IDs                 []string `url:"ids,omitempty,comma"`
	IsDisabled          bool     `url:"isDisabled,omitempty"`
	Name                string   `url:"name,omitempty"`
	PartialName         string   `url:"partialName,omitempty"`
	Roles               []string `url:"roles,comma"`
	ShellNames          []string `url:"shellNames,comma"`
	Skip                int      `url:"skip,omitempty"`
	Take                int      `url:"take,omitempty"`
	TenantIDs           []string `url:"tenantIds,comma"`
	TenantTags          []string `url:"tenantTags,comma"`
	Thumbprint          string   `url:"thumbprint,omitempty"`
}

type OctopusServerNodesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type PackageDeltaSignatureQuery struct {
	PackageID string `url:"packageId,omitempty"`
	Version   string `url:"version,omitempty"`
}

type PackageDeltaUploadQuery struct {
	BaseVersion   string `url:"baseVersion,omitempty"`
	OverwriteMode string `url:"overwriteMode,omitempty"`
	PackageID     string `url:"packageId,omitempty"`
	Replace       bool   `url:"replace,omitempty"`
}

type PackageMetadataQuery struct {
	Filter        string `url:"filter,omitempty"`
	Latest        string `url:"latest,omitempty"`
	OverwriteMode string `url:"overwriteMode,omitempty"`
	Replace       bool   `url:"replace,omitempty"`
	Skip          int    `url:"skip,omitempty"`
	Take          int    `url:"take,omitempty"`
}

type PackageNotesListQuery struct {
	PackageIDs []string `url:"packageIds,comma"`
}

type PackagesQuery struct {
	Filter         string `url:"filter,omitempty"`
	IncludeNotes   bool   `url:"includeNotes,omitempty"`
	Latest         string `url:"latest,omitempty"`
	NuGetPackageID string `url:"nuGetPackageId,omitempty"`
	Skip           int    `url:"skip,omitempty"`
	Take           int    `url:"take,omitempty"`
}

type PackagesBulkQuery struct {
	IDs []string `url:"ids,omitempty,comma"`
}

type PackageUploadQuery struct {
	Replace       bool   `url:"replace,omitempty"`
	OverwriteMode string `url:"overwriteMode,omitempty"`
}

type ProjectGroupsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type ProjectPulseQuery struct {
	ProjectIDs []string `url:"projectIds,comma"`
}

type ProjectsQuery struct {
	ClonedFromProjectID string   `url:"clonedFromProjectId,omitempty"`
	IDs                 []string `url:"ids,omitempty,comma"`
	IsClone             bool     `url:"clone,omitempty"`
	Name                string   `url:"name,omitempty"`
	PartialName         string   `url:"partialName,omitempty"`
	Skip                int      `url:"skip,omitempty"`
	Take                int      `url:"take,omitempty"`
}

type ProjectsExperimentalSummariesQuery struct {
	IDs []string `url:"ids,omitempty,comma"`
}

type ProjectTriggersQuery struct {
	IDs      []string `url:"ids,omitempty,comma"`
	Runbooks []string `url:"runbooks,comma"`
	Skip     int      `url:"skip,omitempty"`
	Take     int      `url:"take,omitempty"`
}

type ProxiesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type ReleasesQuery struct {
	IDs                []string `url:"ids,omitempty,comma"`
	IgnoreChannelRules bool     `url:"ignoreChannelRules"`
	Skip               int      `url:"skip,omitempty"`
	Take               int      `url:"take,omitempty"`
}

type RunbookProcessesQuery struct {
	IDs  []string `url:"ids,omitempty,comma"`
	Skip int      `url:"skip,omitempty"`
	Take int      `url:"take,omitempty"`
}

type RunbookRunsQuery struct {
	Environments []string `url:"environments,comma"`
	IDs          []string `url:"ids,omitempty,comma"`
	PartialName  string   `url:"partialName,omitempty"`
	Projects     []string `url:"projects,comma"`
	Runbooks     []string `url:"runbooks,comma"`
	Skip         int      `url:"skip,omitempty"`
	Take         int      `url:"take,omitempty"`
	TaskState    string   `url:"taskState,omitempty"`
	Tenants      []string `url:"tenants,comma"`
}

type RunbooksQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	IsClone     bool     `url:"clone,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	ProjectIDs  []string `url:"projectIds,comma"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type RunbookSnapshotsQuery struct {
	IDs     []string `url:"ids,omitempty,comma"`
	Publish bool     `url:"publish,omitempty"`
	Skip    int      `url:"skip,omitempty"`
	Take    int      `url:"take,omitempty"`
}

type ScheduledProjectTriggersQuery struct {
	IDs  []string `url:"ids,omitempty,comma"`
	Skip int      `url:"skip,omitempty"`
	Take int      `url:"take,omitempty"`
}

type SchedulerQuery struct {
	Verbose bool   `url:"verbose,omitempty"`
	Tail    string `url:"tail,omitempty"`
}

type ScopedUserRolesQuery struct {
	IDs           []string `url:"ids,omitempty,comma"`
	IncludeSystem bool     `url:"includeSystem"`
	PartialName   string   `url:"partialName,omitempty"`
	Skip          int      `url:"skip,omitempty"`
	Spaces        []string `url:"spaces,comma"`
	Take          int      `url:"take,omitempty"`
}

type SignInQuery struct {
	ReturnURL string `url:"returnUrl,omitempty"`
}

type SpaceHomeQuery struct {
	SpaceID string `url:"spaceId,omitempty"`
}

type SpacesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	Name        string   `url:"name,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type SubscriptionsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Spaces      []string `url:"spaces,comma"`
	Take        int      `url:"take,omitempty"`
}

type TagSetsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type TasksQuery struct {
	Environment             string   `url:"environment,omitempty"`
	HasPendingInterruptions bool     `url:"hasPendingInterruptions"`
	HasWarningsOrErrors     bool     `url:"hasWarningsOrErrors"`
	IDs                     []string `url:"ids,omitempty,comma"`
	IncludeSystem           bool     `url:"includeSystem"`
	IsActive                bool     `url:"active"`
	IsRunning               bool     `url:"running"`
	Name                    string   `url:"name,omitempty"`
	Node                    string   `url:"node,omitempty"`
	PartialName             string   `url:"partialName,omitempty"`
	Project                 string   `url:"project,omitempty"`
	Runbook                 string   `url:"runbook,omitempty"`
	Skip                    int      `url:"skip,omitempty"`
	Spaces                  []string `url:"spaces,comma"`
	States                  []string `url:"states,comma"`
	Take                    int      `url:"take,omitempty"`
	Tenant                  string   `url:"tenant,omitempty"`
}

type TeamMembershipQuery struct {
	IncludeSystem bool     `url:"includeSystem"`
	Spaces        []string `url:"spaces,comma"`
	UserID        string   `url:"userId,omitempty"`
}

type TeamsQuery struct {
	IDs           []string `url:"ids,omitempty,comma"`
	IncludeSystem bool     `url:"includeSystem"`
	PartialName   string   `url:"partialName,omitempty"`
	Skip          int      `url:"skip,omitempty"`
	Spaces        []string `url:"spaces,comma"`
	Take          int      `url:"take,omitempty"`
}

type TenantsQuery struct {
	ClonedFromTenantID string   `url:"clonedFromTenantId,omitempty"`
	IDs                []string `url:"ids,omitempty,comma"`
	IsClone            bool     `url:"clone,omitempty"`
	Name               string   `url:"name,omitempty"`
	PartialName        string   `url:"partialName,omitempty"`
	ProjectID          string   `url:"projectId,omitempty"`
	Skip               int      `url:"skip,omitempty"`
	Tags               []string `url:"tags,comma"`
	Take               int      `url:"take,omitempty"`
}

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `url:"environmentId,omitempty"`
	IncludeDetails bool     `url:"includeDetails,omitempty"`
	ProjectID      string   `url:"projectId,omitempty"`
	TenantID       string   `url:"tenantId,omitempty"`
}

type TenantVariablesQuery struct {
	ProjectID string `url:"projectId,omitempty"`
}

type UserRolesQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type UsersQuery struct {
	Filter string   `url:"filter,omitempty"`
	IDs    []string `url:"ids,omitempty,comma"`
	Skip   int      `url:"skip,omitempty"`
	Take   int      `url:"take,omitempty"`
}

type VariableNamesQuery struct {
	Project                   string `url:"project,omitempty"`
	ProjectEnvironmentsFilter string `struct:"projectEnvironmentsFilter,omitempty"`
	Runbook                   string `url:"runbook,omitempty"`
}

type VariablePreviewQuery struct {
	Action      string `url:"action,omitempty"`
	Channel     string `url:"channel,omitempty"`
	Environment string `url:"environment,omitempty"`
	Machine     string `url:"machine,omitempty"`
	Project     string `url:"project,omitempty"`
	Role        string `url:"role,omitempty"`
	Runbook     string `url:"runbook,omitempty"`
	Tenant      string `url:"tenant,omitempty"`
}

type VariablesQuery struct {
	IDs []string `url:"ids,omitempty,comma"`
}

type VersionRuleTestQuery struct {
	FeetType      string `url:"feetType,omitempty"`
	PreReleaseTag string `url:"preReleaseTag,omitempty"`
	Version       string `url:"version,omitempty"`
	VersionRange  string `url:"versionRange,omitempty"`
}

type WorkerPoolsQuery struct {
	IDs         []string `url:"ids,omitempty,comma"`
	Name        string   `url:"name,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `url:"commStyles,comma"`
	HealthStatuses       []string `url:"healthStatuses,comma"`
	HideEmptyWorkerPools bool     `url:"hideEmptyWorkerPools,omitempty"`
	IDs                  []string `url:"ids,omitempty,comma"`
	IsDisabled           bool     `url:"isDisabled,omitempty"`
	MachinePartialName   string   `url:"machinePartialName,omitempty"`
	PartialName          string   `url:"partialName,omitempty"`
	ShellNames           []string `url:"shellNames,comma"`
}

type WorkersQuery struct {
	CommunicationStyles []string `url:"commStyles,comma"`
	HealthStatuses      []string `url:"healthStatuses,comma"`
	IDs                 []string `url:"ids,omitempty,comma"`
	IsDisabled          bool     `url:"isDisabled,omitempty"`
	Name                string   `url:"name,omitempty"`
	PartialName         string   `url:"partialName,omitempty"`
	ShellNames          []string `url:"shellNames,comma"`
	Skip                int      `url:"skip,omitempty"`
	Take                int      `url:"take,omitempty"`
	Thumbprint          string   `url:"thumbprint,omitempty"`
	WorkerPoolIDs       []string `url:"workerPoolIds,omitempty"`
}
