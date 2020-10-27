package octopusdeploy

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType AccountType `url:"accountType,omitempty"`
	IDs         []string    `url:"ids,omitempty"`
	PartialName string      `url:"partialName,omitempty"`
	Skip        int         `uri:"skip" url:"skip,omitempty"`
	Take        int         `uri:"take" url:"take,omitempty"`
}

type ActionTemplateLogoQuery struct {
	CB       string `uri:"cb" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId" url:"typeOrId,omitempty"`
}

// ActionTemplatesQuery represents parameters to query the ActionTemplates service.
type ActionTemplatesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type ActionTemplateVersionedLogoQuery struct {
	CB       string `uri:"cb" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId" url:"typeOrId,omitempty"`
	Version  string `uri:"version" url:"version,omitempty"`
}

type APIQuery struct {
	Skip int `uri:"skip" url:"skip,omitempty"`
	Take int `uri:"take" url:"take,omitempty"`
}

type ArtifactsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	Order       string   `uri:"order" url:"order,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Regarding   string   `uri:"regarding" url:"regarding,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type BuildInformationQuery struct {
	Filter        string `uri:"filter" url:"filter,omitempty"`
	Latest        string `uri:"latest" url:"latest,omitempty"`
	OverwriteMode string `uri:"overwriteMode" url:"overwriteMode,omitempty"`
	PackageID     string `uri:"packageId" url:"packageId,omitempty"`
	Skip          int    `uri:"skip" url:"skip,omitempty"`
	Take          int    `uri:"take" url:"take,omitempty"`
}

type BuildInformationBulkQuery struct {
	IDs []string `uri:"ids" url:"ids,omitempty"`
}

type CertificateConfigurationQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type CertificatesQuery struct {
	Archived    string   `uri:"archived" url:"archived,omitempty"`
	FirstResult string   `uri:"firstResult" url:"firstResult,omitempty"`
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	OrderBy     string   `uri:"orderBy" url:"orderBy,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Search      string   `uri:"search" url:"search,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
	Tenant      string   `uri:"tenant" url:"tenant,omitempty"`
}

type ChannelsQuery struct {
	IDs         []string `url:"ids,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type CloudTemplateQuery struct {
	FeedID    string `uri:"feedId" url:"feedId,omitempty"`
	PackageID string `uri:"packageId" url:"packageId,omitempty"`
}

type CommunityActionTemplatesQuery struct {
	IDs  []string `url:"ids,omitempty"`
	Skip int      `uri:"skip" url:"skip,omitempty"`
	Take int      `uri:"take" url:"take,omitempty"`
}

type DashboardQuery struct {
	IncludeLatest   bool     `url:"highestLatestVersionPerProjectAndEnvironment"`
	ProjectID       string   `uri:"projectId" url:"projectId,omitempty"`
	SelectedTags    []string `uri:"selectedTags" url:"selectedTags,omitempty"`
	SelectedTenants []string `uri:"selectedTenants" url:"selectedTenants,omitempty"`
	ShowAll         bool     `uri:"showAll" url:"showAll,omitempty"`
	ReleaseID       string   `uri:"releaseId" url:"releaseId,omitempty"`
}

type DashboardDynamicQuery struct {
	Environments    []string `uri:"environments" url:"environments,omitempty"`
	IncludePrevious bool     `uri:"includePrevious" url:"includePrevious,omitempty"`
	Projects        []string `uri:"projects" url:"projects,omitempty"`
}

type DeploymentProcessesQuery struct {
	IDs  []string `uri:"ids" url:"ids,omitempty"`
	Skip int      `uri:"skip" url:"skip,omitempty"`
	Take int      `uri:"take" url:"take,omitempty"`
}

type DeploymentQuery struct {
	Skip int `uri:"skip" url:"skip,omitempty"`
	Take int `uri:"take" url:"take,omitempty"`
}

type DeploymentsQuery struct {
	Channels     string   `uri:"channels" url:"channels,omitempty"`
	Environments []string `uri:"environments" url:"environments,omitempty"`
	IDs          []string `uri:"ids" url:"ids,omitempty"`
	PartialName  string   `uri:"partialName" url:"partialName,omitempty"`
	Projects     []string `uri:"projects" url:"projects,omitempty"`
	Skip         int      `uri:"skip" url:"skip,omitempty"`
	Take         int      `uri:"take" url:"take,omitempty"`
	TaskState    string   `uri:"taskState" url:"taskState,omitempty"`
	Tenants      []string `uri:"tenants" url:"tenants,omitempty"`
}

type DiscoverMachineQuery struct {
	Host    string `uri:"host" url:"host,omitempty"`
	Port    int    `uri:"port" url:"port,omitempty"`
	ProxyID string `uri:"proxyId" url:"proxyId,omitempty"`
	Type    string `uri:"type" url:"type,omitempty"`
}

type DiscoverWorkerQuery struct {
	Host    string `uri:"host" url:"host,omitempty"`
	Port    int    `uri:"port" url:"port,omitempty"`
	ProxyID string `uri:"proxyId" url:"proxyId,omitempty"`
	Type    string `uri:"type" url:"type,omitempty"`
}

type EnvironmentsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	Name        string   `uri:"name" url:"name,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `uri:"commStyles" url:"commStyles,omitempty"`
	HealthStatuses        []string `uri:"healthStatuses" url:"healthStatuses,omitempty"`
	HideEmptyEnvironments bool     `uri:"hideEmptyEnvironments" url:"hideEmptyEnvironments,omitempty"`
	IDs                   []string `uri:"ids" url:"ids,omitempty"`
	IsDisabled            bool     `uri:"isDisabled" url:"isDisabled,omitempty"`
	MachinePartialName    string   `uri:"machinePartialName" url:"machinePartialName,omitempty"`
	PartialName           string   `uri:"partialName" url:"partialName,omitempty"`
	Roles                 []string `uri:"roles" url:"roles,omitempty"`
	ShellNames            []string `uri:"shellNames" url:"shellNames,omitempty"`
	TenantIDs             []string `uri:"tenantIds" url:"tenantIds,omitempty"`
	TenantTags            []string `uri:"tenantTags" url:"tenantTags,omitempty"`
}

type EventCategoriesQuery struct {
	AppliesTo string `uri:"appliesTo" url:"appliesTo,omitempty"`
}

type EventGroupsQuery struct {
	AppliesTo string `uri:"appliesTo" url:"appliesTo,omitempty"`
}

type EventsQuery struct {
	AsCSV             string   `uri:"asCsv" url:"asCsv,omitempty"`
	DocumentTypes     []string `uri:"documentTypes" url:"documentTypes,omitempty"`
	Environments      []string `uri:"environments" url:"environments,omitempty"`
	EventAgents       []string `uri:"eventAgents" url:"eventAgents,omitempty"`
	EventCategories   []string `uri:"eventCategories" url:"eventCategories,omitempty"`
	EventGroups       []string `uri:"eventGroups" url:"eventGroups,omitempty"`
	ExcludeDifference bool     `uri:"excludeDifference" url:"excludeDifference,omitempty"`
	From              string   `uri:"from" url:"from,omitempty"`
	FromAutoID        string   `uri:"fromAutoId" url:"fromAutoId,omitempty"`
	IDs               []string `uri:"ids" url:"ids,omitempty"`
	IncludeSystem     bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	Internal          string   `uri:"interal" url:"interal,omitempty"`
	Name              string   `uri:"name" url:"name,omitempty"`
	PartialName       string   `uri:"partialName" url:"partialName,omitempty"`
	ProjectGroups     []string `uri:"projectGroups" url:"projectGroups,omitempty"`
	Projects          []string `uri:"projects" url:"projects,omitempty"`
	Regarding         string   `uri:"regarding" url:"regarding,omitempty"`
	RegardingAny      string   `uri:"regardingAny" url:"regardingAny,omitempty"`
	Skip              int      `uri:"skip" url:"skip,omitempty"`
	Spaces            []string `uri:"spaces" url:"spaces,omitempty"`
	Take              int      `uri:"take" url:"take,omitempty"`
	Tags              []string `uri:"tags" url:"tags,omitempty"`
	Tenants           []string `uri:"tenants" url:"tenants,omitempty"`
	To                string   `uri:"to" url:"to,omitempty"`
	ToAutoID          string   `uri:"toAutoId" url:"toAutoId,omitempty"`
	User              string   `uri:"user" url:"user,omitempty"`
	Users             []string `uri:"users" url:"users,omitempty"`
}

type ExternalUserSearchQuery struct {
	PartialName string `uri:"partialName" url:"partialName,omitempty"`
}

type FeedsQuery struct {
	FeedType    string   `uri:"feedType" url:"feedType,omitempty"`
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type InterruptionsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PendingOnly bool     `uri:"pendingOnly" url:"pendingOnly,omitempty"`
	Regarding   string   `uri:"regarding" url:"regarding,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type IssueTrackersQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type LibraryVariablesQuery struct {
	ContentType string   `uri:"contentType" url:"contentType,omitempty"`
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type LifecyclesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type MachinePoliciesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type MachinesQuery struct {
	CommunicationStyles []string `uri:"commStyles" url:"commStyles,omitempty"`
	DeploymentID        string   `uri:"deploymentId" url:"deploymentId,omitempty"`
	EnvironmentIDs      []string `uri:"environmentIds" url:"environmentIds,omitempty"`
	HealthStatuses      []string `uri:"healthStatuses" url:"healthStatuses,omitempty"`
	IDs                 []string `uri:"ids" url:"ids,omitempty"`
	IsDisabled          bool     `uri:"isDisabled" url:"isDisabled,omitempty"`
	Name                string   `uri:"name" url:"name,omitempty"`
	PartialName         string   `uri:"partialName" url:"partialName,omitempty"`
	Roles               []string `uri:"roles" url:"roles,omitempty"`
	ShellNames          []string `uri:"shellNames" url:"shellNames,omitempty"`
	Skip                int      `uri:"skip" url:"skip,omitempty"`
	Take                int      `uri:"take" url:"take,omitempty"`
	TenantIDs           []string `uri:"tenantIds" url:"tenantIds,omitempty"`
	TenantTags          []string `uri:"tenantTags" url:"tenantTags,omitempty"`
	Thumbprint          string   `uri:"thumbprint" url:"thumbprint,omitempty"`
}

type OctopusServerNodesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type PackageDeltaSignatureQuery struct {
	PackageID string `uri:"packageId" url:"packageId,omitempty"`
	Version   string `uri:"version" url:"version,omitempty"`
}

type PackageDeltaUploadQuery struct {
	BaseVersion   string `uri:"baseVersion" url:"baseVersion,omitempty"`
	OverwriteMode string `uri:"overwriteMode" url:"overwriteMode,omitempty"`
	PackageID     string `uri:"packageId" url:"packageId,omitempty"`
	Replace       bool   `uri:"replace" url:"replace,omitempty"`
}

type PackageMetadataQuery struct {
	Filter        string `uri:"filter" url:"filter,omitempty"`
	Latest        string `uri:"latest" url:"latest,omitempty"`
	OverwriteMode string `uri:"overwriteMode" url:"overwriteMode,omitempty"`
	Replace       bool   `uri:"replace" url:"replace,omitempty"`
	Skip          int    `uri:"skip" url:"skip,omitempty"`
	Take          int    `uri:"take" url:"take,omitempty"`
}

type PackageNotesListQuery struct {
	PackageIDs []string `uri:"packageIds" url:"packageIds,omitempty"`
}

type PackagesQuery struct {
	Filter         string `uri:"filter" url:"filter,omitempty"`
	IncludeNotes   bool   `uri:"includeNotes" url:"includeNotes,omitempty"`
	Latest         string `uri:"latest" url:"latest,omitempty"`
	NuGetPackageID string `uri:"nuGetPackageId" url:"nuGetPackageId,omitempty"`
	Skip           int    `uri:"skip" url:"skip,omitempty"`
	Take           int    `uri:"take" url:"take,omitempty"`
}

type PackagesBulkQuery struct {
	IDs []string `uri:"ids" url:"ids,omitempty"`
}

type PackageUploadQuery struct {
	Replace       bool   `uri:"replace" url:"replace,omitempty"`
	OverwriteMode string `uri:"overwriteMode" url:"overwriteMode,omitempty"`
}

type UserQuery struct {
	IncludeSystem bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces" url:"spaces,omitempty"`
}

type ProjectGroupsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type ProjectPulseQuery struct {
	ProjectIDs []string `uri:"projectIds" url:"projectIds,omitempty"`
}

type ProjectsQuery struct {
	ClonedFromProjectID string   `url:"clonedFromProjectId"`
	IDs                 []string `uri:"ids" url:"ids,omitempty"`
	IsClone             bool     `uri:"clone" url:"clone,omitempty"`
	Name                string   `uri:"name" url:"name,omitempty"`
	PartialName         string   `uri:"partialName" url:"partialName,omitempty"`
	Skip                int      `uri:"skip" url:"skip,omitempty"`
	Take                int      `uri:"take" url:"take,omitempty"`
}

type ProjectsExperimentalSummariesQuery struct {
	IDs []string `uri:"ids" url:"ids,omitempty"`
}

type ProjectTriggersQuery struct {
	IDs      []string `uri:"ids" url:"ids,omitempty"`
	Runbooks []string `uri:"runbooks" url:"runbooks,omitempty"`
	Skip     int      `uri:"skip" url:"skip,omitempty"`
	Take     int      `uri:"take" url:"take,omitempty"`
}

type ProxiesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type ReleaseQuery struct {
	SearchByVersion string `uri:"searchByVersion" url:"searchByVersion,omitempty"`
	Skip            int    `uri:"skip" url:"skip,omitempty"`
	Take            int    `uri:"take" url:"take,omitempty"`
}

type ReleasesQuery struct {
	IDs                []string `uri:"ids" url:"ids,omitempty"`
	IgnoreChannelRules bool     `url:"ignoreChannelRules"`
	Skip               int      `uri:"skip" url:"skip,omitempty"`
	Take               int      `uri:"take" url:"take,omitempty"`
}

type RunbookProcessesQuery struct {
	IDs  []string `uri:"ids" url:"ids,omitempty"`
	Skip int      `uri:"skip" url:"skip,omitempty"`
	Take int      `uri:"take" url:"take,omitempty"`
}

type RunbookRunsQuery struct {
	Environments []string `uri:"environments" url:"environments,omitempty"`
	IDs          []string `uri:"ids" url:"ids,omitempty"`
	PartialName  string   `uri:"partialName" url:"partialName,omitempty"`
	Projects     []string `uri:"projects" url:"projects,omitempty"`
	Runbooks     []string `uri:"runbooks" url:"runbooks,omitempty"`
	Skip         int      `uri:"skip" url:"skip,omitempty"`
	Take         int      `uri:"take" url:"take,omitempty"`
	TaskState    string   `uri:"taskState" url:"taskState,omitempty"`
	Tenants      []string `uri:"tenants" url:"tenants,omitempty"`
}

type RunbooksQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	IsClone     bool     `uri:"clone" url:"clone,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	ProjectIDs  []string `uri:"projectIds" url:"projectIds,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type RunbookSnapshotsQuery struct {
	IDs     []string `uri:"ids" url:"ids,omitempty"`
	Publish bool     `uri:"publish" url:"publish,omitempty"`
	Skip    int      `uri:"skip" url:"skip,omitempty"`
	Take    int      `uri:"take" url:"take,omitempty"`
}

type ScheduledProjectTriggersQuery struct {
	IDs  []string `uri:"ids" url:"ids,omitempty"`
	Skip int      `uri:"skip" url:"skip,omitempty"`
	Take int      `uri:"take" url:"take,omitempty"`
}

type SchedulerQuery struct {
	Verbose bool   `uri:"verbose" url:"verbose,omitempty"`
	Tail    string `uri:"tail" url:"tail,omitempty"`
}

type ScopedUserRolesQuery struct {
	IDs           []string `uri:"ids" url:"ids,omitempty"`
	IncludeSystem bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	PartialName   string   `uri:"partialName" url:"partialName,omitempty"`
	Skip          int      `uri:"skip" url:"skip,omitempty"`
	Spaces        []string `uri:"spaces" url:"spaces,omitempty"`
	Take          int      `uri:"take" url:"take,omitempty"`
}

type SignInQuery struct {
	ReturnURL string `uri:"returnUrl" url:"returnUrl,omitempty"`
}

type SpaceHomeQuery struct {
	SpaceID string `uri:"spaceId" url:"spaceId,omitempty"`
}

type SpacesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	Name        string   `uri:"name" url:"name,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type SubscriptionsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Spaces      []string `uri:"spaces" url:"spaces,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type TagSetsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type TasksQuery struct {
	Environment             string   `uri:"environment" url:"environment,omitempty"`
	HasPendingInterruptions bool     `url:"hasPendingInterruptions"`
	HasWarningsOrErrors     bool     `url:"hasWarningsOrErrors"`
	IDs                     []string `uri:"ids" url:"ids,omitempty"`
	IncludeSystem           bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	IsActive                bool     `uri:"active" url:"active,omitempty"`
	IsRunning               bool     `uri:"running" url:"running,omitempty"`
	Name                    string   `uri:"name" url:"name,omitempty"`
	Node                    string   `uri:"node" url:"node,omitempty"`
	PartialName             string   `uri:"partialName" url:"partialName,omitempty"`
	Project                 string   `uri:"project" url:"project,omitempty"`
	Runbook                 string   `uri:"runbook" url:"runbook,omitempty"`
	Skip                    int      `uri:"skip" url:"skip,omitempty"`
	Spaces                  []string `uri:"spaces" url:"spaces,omitempty"`
	States                  []string `uri:"states" url:"states,omitempty"`
	Take                    int      `uri:"take" url:"take,omitempty"`
	Tenant                  string   `uri:"tenant" url:"tenant,omitempty"`
}

type TeamMembershipQuery struct {
	IncludeSystem bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces" url:"spaces,omitempty"`
	UserID        string   `uri:"userId" url:"userId,omitempty"`
}

type TeamsQuery struct {
	IDs           []string `uri:"ids" url:"ids,omitempty"`
	IncludeSystem bool     `uri:"includeSystem" url:"includeSystem,omitempty"`
	PartialName   string   `uri:"partialName" url:"partialName,omitempty"`
	Skip          int      `uri:"skip" url:"skip,omitempty"`
	Spaces        []string `uri:"spaces" url:"spaces,omitempty"`
	Take          int      `uri:"take" url:"take,omitempty"`
}

type TenantsQuery struct {
	ClonedFromTenantID string   `url:"clonedFromTenantId"`
	IDs                []string `uri:"ids" url:"ids,omitempty"`
	IsClone            bool     `uri:"clone" url:"clone,omitempty"`
	Name               string   `uri:"name" url:"name,omitempty"`
	PartialName        string   `uri:"partialName" url:"partialName,omitempty"`
	ProjectID          string   `uri:"projectId" url:"projectId,omitempty"`
	Skip               int      `uri:"skip" url:"skip,omitempty"`
	Tags               []string `uri:"tags" url:"tags,omitempty"`
	Take               int      `uri:"take" url:"take,omitempty"`
}

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `uri:"environmentId" url:"environmentId,omitempty"`
	IncludeDetails bool     `uri:"includeDetails" url:"includeDetails,omitempty"`
	ProjectID      string   `uri:"projectId" url:"projectId,omitempty"`
	TenantID       string   `uri:"tenantId" url:"tenantId,omitempty"`
}

type TenantVariablesQuery struct {
	ProjectID string `uri:"projectId" url:"projectId,omitempty"`
}

type UserRolesQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type UsersQuery struct {
	Filter string   `uri:"filter" url:"filter,omitempty"`
	IDs    []string `uri:"ids" url:"ids,omitempty"`
	Skip   int      `uri:"skip" url:"skip,omitempty"`
	Take   int      `uri:"take" url:"take,omitempty"`
}

type VariableNamesQuery struct {
	Project                   string `uri:"project" url:"project,omitempty"`
	ProjectEnvironmentsFilter string `uri:"projectEnvironmentsFilter" url:"projectEnvironmentsFilter,omitempty"`
	Runbook                   string `uri:"runbook" url:"runbook,omitempty"`
}

type VariablePreviewQuery struct {
	Action      string `uri:"action" url:"action,omitempty"`
	Channel     string `uri:"channel" url:"channel,omitempty"`
	Environment string `uri:"environment" url:"environment,omitempty"`
	Machine     string `uri:"machine" url:"machine,omitempty"`
	Project     string `uri:"project" url:"project,omitempty"`
	Role        string `uri:"role" url:"role,omitempty"`
	Runbook     string `uri:"runbook" url:"runbook,omitempty"`
	Tenant      string `uri:"tenant" url:"tenant,omitempty"`
}

type VariablesQuery struct {
	IDs []string `uri:"ids" url:"ids,omitempty"`
}

type VersionRuleTestQuery struct {
	FeetType      string `uri:"feetType" url:"feetType,omitempty"`
	PreReleaseTag string `uri:"preReleaseTag" url:"preReleaseTag,omitempty"`
	Version       string `uri:"version" url:"version,omitempty"`
	VersionRange  string `uri:"versionRange" url:"versionRange,omitempty"`
}

type WorkerPoolsQuery struct {
	IDs         []string `uri:"ids" url:"ids,omitempty"`
	Name        string   `uri:"name" url:"name,omitempty"`
	PartialName string   `uri:"partialName" url:"partialName,omitempty"`
	Skip        int      `uri:"skip" url:"skip,omitempty"`
	Take        int      `uri:"take" url:"take,omitempty"`
}

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `uri:"commStyles" url:"commStyles,omitempty"`
	HealthStatuses       []string `uri:"healthStatuses" url:"healthStatuses,omitempty"`
	HideEmptyWorkerPools bool     `uri:"hideEmptyWorkerPools" url:"hideEmptyWorkerPools,omitempty"`
	IDs                  []string `uri:"ids" url:"ids,omitempty"`
	IsDisabled           bool     `uri:"isDisabled" url:"isDisabled,omitempty"`
	MachinePartialName   string   `uri:"machinePartialName" url:"machinePartialName,omitempty"`
	PartialName          string   `uri:"partialName" url:"partialName,omitempty"`
	ShellNames           []string `uri:"shellNames" url:"shellNames,omitempty"`
}

type WorkersQuery struct {
	CommunicationStyles []string `uri:"commStyles" url:"commStyles,omitempty"`
	HealthStatuses      []string `uri:"healthStatuses" url:"healthStatuses,omitempty"`
	IDs                 []string `uri:"ids" url:"ids,omitempty"`
	IsDisabled          bool     `uri:"isDisabled" url:"isDisabled,omitempty"`
	Name                string   `uri:"name" url:"name,omitempty"`
	PartialName         string   `uri:"partialName" url:"partialName,omitempty"`
	ShellNames          []string `uri:"shellNames" url:"shellNames,omitempty"`
	Skip                int      `uri:"skip" url:"skip,omitempty"`
	Take                int      `uri:"take" url:"take,omitempty"`
	Thumbprint          string   `uri:"thumbprint" url:"thumbprint,omitempty"`
	WorkerPoolIDs       []string `uri:"workerPoolIds" url:"workerPoolIds,omitempty"`
}
