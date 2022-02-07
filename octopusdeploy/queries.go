package octopusdeploy

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"

type Query struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}

type IdsQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}

type PartialNameQuery struct {
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
}

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType accounts.AccountType `uri:"accountType,omitempty" url:"accountType,omitempty"`
	IdsQuery
	PartialNameQuery
	Query
}

type ActionTemplateLogoQuery struct {
	CB       string `uri:"cb,omitempty" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId,omitempty" url:"typeOrId,omitempty"`
}

// ActionTemplatesQuery represents parameters to query the ActionTemplates service.
type ActionTemplatesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ActionTemplateVersionedLogoQuery struct {
	CB       string `uri:"cb,omitempty" url:"cb,omitempty"`
	TypeOrID string `uri:"typeOrId,omitempty" url:"typeOrId,omitempty"`
	Version  string `uri:"version,omitempty" url:"version,omitempty"`
}

type APIQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}

type ArtifactsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Order       string   `uri:"order" url:"order,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Regarding   string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type BuildInformationQuery struct {
	Filter        string `uri:"filter,omitempty" url:"filter,omitempty"`
	Latest        string `uri:"latest,omitempty" url:"latest,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	PackageID     string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	Skip          int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take          int    `uri:"take,omitempty" url:"take,omitempty"`
}

type BuildInformationBulkQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}

type CertificateConfigurationQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type CertificatesQuery struct {
	Archived    string   `uri:"archived,omitempty" url:"archived,omitempty"`
	FirstResult string   `uri:"firstResult,omitempty" url:"firstResult,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	OrderBy     string   `uri:"orderBy,omitempty" url:"orderBy,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Search      string   `uri:"search,omitempty" url:"search,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
	Tenant      string   `uri:"tenant,omitempty" url:"tenant,omitempty"`
}

type ChannelsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type CloudTemplateQuery struct {
	FeedID    string `uri:"feedId,omitempty" url:"feedId,omitempty"`
	PackageID string `uri:"packageId,omitempty" url:"packageId,omitempty"`
}

type CommunityActionTemplatesQuery struct {
	IDs  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int      `uri:"take,omitempty" url:"take,omitempty"`
}

type DashboardQuery struct {
	IncludeLatest   bool     `url:"highestLatestVersionPerProjectAndEnvironment"`
	ProjectID       string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	SelectedTags    []string `uri:"selectedTags,omitempty" url:"selectedTags,omitempty"`
	SelectedTenants []string `uri:"selectedTenants,omitempty" url:"selectedTenants,omitempty"`
	ShowAll         bool     `uri:"showAll,omitempty" url:"showAll,omitempty"`
	ReleaseID       string   `uri:"releaseId,omitempty" url:"releaseId,omitempty"`
}

type DashboardDynamicQuery struct {
	Environments    []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IncludePrevious bool     `uri:"includePrevious,omitempty" url:"includePrevious,omitempty"`
	Projects        []string `uri:"projects,omitempty" url:"projects,omitempty"`
}

type DeploymentProcessesQuery struct {
	IDs  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int      `uri:"take,omitempty" url:"take,omitempty"`
}

type DeploymentQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}

type DeploymentsQuery struct {
	Channels     string   `uri:"channels,omitempty" url:"channels,omitempty"`
	Environments []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IDs          []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName  string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Projects     []string `uri:"projects,omitempty" url:"projects,omitempty"`
	Skip         int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take         int      `uri:"take,omitempty" url:"take,omitempty"`
	TaskState    string   `uri:"taskState,omitempty" url:"taskState,omitempty"`
	Tenants      []string `uri:"tenants,omitempty" url:"tenants,omitempty"`
}

type DiscoverMachineQuery struct {
	Host    string `uri:"host,omitempty" url:"host,omitempty"`
	Port    int    `uri:"port,omitempty" url:"port,omitempty"`
	ProxyID string `uri:"proxyId,omitempty" url:"proxyId,omitempty"`
	Type    string `uri:"type,omitempty" url:"type,omitempty"`
}

type DiscoverWorkerQuery struct {
	Host    string `uri:"host,omitempty" url:"host,omitempty"`
	Port    int    `uri:"port,omitempty" url:"port,omitempty"`
	ProxyID string `uri:"proxyId,omitempty" url:"proxyId,omitempty"`
	Type    string `uri:"type,omitempty" url:"type,omitempty"`
}

type EnvironmentsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Name        string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	HealthStatuses        []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	HideEmptyEnvironments bool     `uri:"hideEmptyEnvironments,omitempty" url:"hideEmptyEnvironments,omitempty"`
	IDs                   []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled            bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	MachinePartialName    string   `uri:"machinePartialName,omitempty" url:"machinePartialName,omitempty"`
	PartialName           string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Roles                 []string `uri:"roles,omitempty" url:"roles,omitempty"`
	ShellNames            []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	TenantIDs             []string `uri:"tenantIds,omitempty" url:"tenantIds,omitempty"`
	TenantTags            []string `uri:"tenantTags,omitempty" url:"tenantTags,omitempty"`
}

type EventCategoriesQuery struct {
	AppliesTo string `uri:"appliesTo" url:"appliesTo,omitempty"`
}

type EventGroupsQuery struct {
	AppliesTo string `uri:"appliesTo" url:"appliesTo,omitempty"`
}

type EventsQuery struct {
	AsCSV             string   `uri:"asCsv,omitempty" url:"asCsv,omitempty"`
	DocumentTypes     []string `uri:"documentTypes,omitempty" url:"documentTypes,omitempty"`
	Environments      []string `uri:"environments,omitempty" url:"environments,omitempty"`
	EventAgents       []string `uri:"eventAgents,omitempty" url:"eventAgents,omitempty"`
	EventCategories   []string `uri:"eventCategories,omitempty" url:"eventCategories,omitempty"`
	EventGroups       []string `uri:"eventGroups,omitempty" url:"eventGroups,omitempty"`
	ExcludeDifference bool     `uri:"excludeDifference,omitempty" url:"excludeDifference,omitempty"`
	From              string   `uri:"from,omitempty" url:"from,omitempty"`
	FromAutoID        string   `uri:"fromAutoId,omitempty" url:"fromAutoId,omitempty"`
	IDs               []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem     bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Internal          string   `uri:"interal,omitempty" url:"interal,omitempty"`
	Name              string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName       string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectGroups     []string `uri:"projectGroups,omitempty" url:"projectGroups,omitempty"`
	Projects          []string `uri:"projects,omitempty" url:"projects,omitempty"`
	Regarding         string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	RegardingAny      string   `uri:"regardingAny,omitempty" url:"regardingAny,omitempty"`
	Skip              int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces            []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take              int      `uri:"take,omitempty" url:"take,omitempty"`
	Tags              []string `uri:"tags,omitempty" url:"tags,omitempty"`
	Tenants           []string `uri:"tenants,omitempty" url:"tenants,omitempty"`
	To                string   `uri:"to,omitempty" url:"to,omitempty"`
	ToAutoID          string   `uri:"toAutoId,omitempty" url:"toAutoId,omitempty"`
	User              string   `uri:"user,omitempty" url:"user,omitempty"`
	Users             []string `uri:"users,omitempty" url:"users,omitempty"`
}

type ExternalUserSearchQuery struct {
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
}

type FeedsQuery struct {
	FeedType    string   `uri:"feedType,omitempty" url:"feedType,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type InterruptionsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PendingOnly bool     `uri:"pendingOnly,omitempty" url:"pendingOnly,omitempty"`
	Regarding   string   `uri:"regarding,omitempty" url:"regarding,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type IssueTrackersQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type LibraryVariablesQuery struct {
	ContentType string   `uri:"contentType,omitempty" url:"contentType,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type LifecyclesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type MachinePoliciesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type MachinesQuery struct {
	CommunicationStyles []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	DeploymentID        string   `uri:"deploymentId,omitempty" url:"deploymentId,omitempty"`
	EnvironmentIDs      []string `uri:"environmentIds,omitempty" url:"environmentIds,omitempty"`
	HealthStatuses      []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	IDs                 []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled          bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	Name                string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName         string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Roles               []string `uri:"roles,omitempty" url:"roles,omitempty"`
	ShellNames          []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	Skip                int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                int      `uri:"take,omitempty" url:"take,omitempty"`
	TenantIDs           []string `uri:"tenantIds,omitempty" url:"tenantIds,omitempty"`
	TenantTags          []string `uri:"tenantTags,omitempty" url:"tenantTags,omitempty"`
	Thumbprint          string   `uri:"thumbprint,omitempty" url:"thumbprint,omitempty"`
}

type OctopusServerNodesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type PackageDeltaSignatureQuery struct {
	PackageID string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	Version   string `uri:"version,omitempty" url:"version,omitempty"`
}

type PackageDeltaUploadQuery struct {
	BaseVersion   string `uri:"baseVersion,omitempty" url:"baseVersion,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	PackageID     string `uri:"packageId,omitempty" url:"packageId,omitempty"`
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
}

type PackageMetadataQuery struct {
	Filter        string `uri:"filter,omitempty" url:"filter,omitempty"`
	Latest        string `uri:"latest,omitempty" url:"latest,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
	Skip          int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take          int    `uri:"take,omitempty" url:"take,omitempty"`
}

type PackageNotesListQuery struct {
	PackageIDs []string `uri:"packageIds,omitempty" url:"packageIds,omitempty"`
}

type PackagesQuery struct {
	Filter         string `uri:"filter,omitempty" url:"filter,omitempty"`
	IncludeNotes   bool   `uri:"includeNotes,omitempty" url:"includeNotes,omitempty"`
	Latest         string `uri:"latest,omitempty" url:"latest,omitempty"`
	NuGetPackageID string `uri:"nuGetPackageId,omitempty" url:"nuGetPackageId,omitempty"`
	Skip           int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take           int    `uri:"take,omitempty" url:"take,omitempty"`
}

type PackagesBulkQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}

type PackageUploadQuery struct {
	Replace       bool   `uri:"replace,omitempty" url:"replace,omitempty"`
	OverwriteMode string `uri:"overwriteMode,omitempty" url:"overwriteMode,omitempty"`
}

type ProjectGroupsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ProjectPulseQuery struct {
	ProjectIDs []string `uri:"projectIds,omitempty" url:"projectIds,omitempty"`
}

type ProjectsQuery struct {
	ClonedFromProjectID string   `url:"clonedFromProjectId"`
	IDs                 []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone             bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	Name                string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName         string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip                int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ProjectsExperimentalSummariesQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}

type ProjectTriggersQuery struct {
	IDs      []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Runbooks []string `uri:"runbooks,omitempty" url:"runbooks,omitempty"`
	Skip     int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take     int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ProxiesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ReleaseQuery struct {
	SearchByVersion string `uri:"searchByVersion" url:"searchByVersion,omitempty"`
	Skip            int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take            int    `uri:"take,omitempty" url:"take,omitempty"`
}

type ReleasesQuery struct {
	IDs                []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IgnoreChannelRules bool     `uri:"ignoreChannelRules,omitempty" url:"ignoreChannelRules,omitempty"`
	Skip               int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take               int      `uri:"take,omitempty" url:"take,omitempty"`
}

type RunbookProcessesQuery struct {
	IDs  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int      `uri:"take,omitempty" url:"take,omitempty"`
}

type RunbookRunsQuery struct {
	Environments []string `uri:"environments,omitempty" url:"environments,omitempty"`
	IDs          []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName  string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Projects     []string `uri:"projects,omitempty" url:"projects,omitempty"`
	Runbooks     []string `uri:"runbooks,omitempty" url:"runbooks,omitempty"`
	Skip         int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take         int      `uri:"take,omitempty" url:"take,omitempty"`
	TaskState    string   `uri:"taskState,omitempty" url:"taskState,omitempty"`
	Tenants      []string `uri:"tenants,omitempty" url:"tenants,omitempty"`
}

type RunbooksQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone     bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectIDs  []string `uri:"projectIds,omitempty" url:"projectIds,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type RunbookSnapshotsQuery struct {
	IDs     []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Publish bool     `uri:"publish,omitempty" url:"publish,omitempty"`
	Skip    int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take    int      `uri:"take,omitempty" url:"take,omitempty"`
}

type ScheduledProjectTriggersQuery struct {
	IDs  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int      `uri:"take,omitempty" url:"take,omitempty"`
}

type SchedulerQuery struct {
	Verbose bool   `uri:"verbose,omitempty" url:"verbose,omitempty"`
	Tail    string `uri:"tail,omitempty" url:"tail,omitempty"`
}

type ScopedUserRolesQuery struct {
	IDs           []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	PartialName   string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip          int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take          int      `uri:"take,omitempty" url:"take,omitempty"`
}

type SearchPackagesQuery struct {
	Skip int    `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int    `uri:"take,omitempty" url:"take,omitempty"`
	Term string `uri:"term,omitempty" url:"term,omitempty"`
}

type SignInQuery struct {
	ReturnURL string `uri:"returnUrl,omitempty" url:"returnUrl,omitempty"`
}

type SkipTakeQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}

type SpaceHomeQuery struct {
	SpaceID string `uri:"spaceId,omitempty" url:"spaceId,omitempty"`
}

type SpacesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Name        string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type SubscriptionsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces      []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type TagSetsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type TasksQuery struct {
	Environment             string   `uri:"environment,omitempty" url:"environment,omitempty"`
	HasPendingInterruptions bool     `uri:"hasPendingInterruptions,omitempty" url:"hasPendingInterruptions,omitempty"`
	HasWarningsOrErrors     bool     `uri:"hasWarningsOrErrors,omitempty" url:"hasWarningsOrErrors,omitempty"`
	IDs                     []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem           bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	IsActive                bool     `uri:"active,omitempty" url:"active,omitempty"`
	IsRunning               bool     `uri:"running,omitempty" url:"running,omitempty"`
	Name                    string   `uri:"name,omitempty" url:"name,omitempty"`
	Node                    string   `uri:"node,omitempty" url:"node,omitempty"`
	PartialName             string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Project                 string   `uri:"project,omitempty" url:"project,omitempty"`
	Runbook                 string   `uri:"runbook,omitempty" url:"runbook,omitempty"`
	Skip                    int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces                  []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	States                  []string `uri:"states,omitempty" url:"states,omitempty"`
	Take                    int      `uri:"take,omitempty" url:"take,omitempty"`
	Tenant                  string   `uri:"tenant,omitempty" url:"tenant,omitempty"`
}

type TeamMembershipQuery struct {
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	UserID        string   `uri:"userId,omitempty" url:"userId,omitempty"`
}

type TeamsQuery struct {
	IDs           []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	PartialName   string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip          int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
	Take          int      `uri:"take,omitempty" url:"take,omitempty"`
}

type TenantsQuery struct {
	ClonedFromTenantID string   `uri:"clonedFromTenantId,omitempty" url:"clonedFromTenantId,omitempty"`
	IDs                []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsClone            bool     `uri:"clone,omitempty" url:"clone,omitempty"`
	Name               string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName        string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ProjectID          string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	Skip               int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Tags               []string `uri:"tags,omitempty" url:"tags,omitempty"`
	Take               int      `uri:"take,omitempty" url:"take,omitempty"`
}

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `uri:"environmentId,omitempty" url:"environmentId,omitempty"`
	IncludeDetails bool     `uri:"includeDetails,omitempty" url:"includeDetails,omitempty"`
	ProjectID      string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	TenantID       string   `uri:"tenantId,omitempty" url:"tenantId,omitempty"`
}

type TenantVariablesQuery struct {
	ProjectID string `uri:"projectId,omitempty" url:"projectId,omitempty"`
}

type UserQuery struct {
	IncludeSystem bool     `uri:"includeSystem,omitempty" url:"includeSystem,omitempty"`
	Spaces        []string `uri:"spaces,omitempty" url:"spaces,omitempty"`
}

type UserRolesQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type UsersQuery struct {
	Filter string   `uri:"filter,omitempty" url:"filter,omitempty"`
	IDs    []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Skip   int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take   int      `uri:"take,omitempty" url:"take,omitempty"`
}

type VariableNamesQuery struct {
	Project                   string `uri:"project,omitempty" url:"project,omitempty"`
	ProjectEnvironmentsFilter string `uri:"projectEnvironmentsFilter,omitempty" url:"projectEnvironmentsFilter,omitempty"`
	Runbook                   string `uri:"runbook,omitempty" url:"runbook,omitempty"`
}

type VariablePreviewQuery struct {
	Action      string `uri:"action,omitempty" url:"action,omitempty"`
	Channel     string `uri:"channel,omitempty" url:"channel,omitempty"`
	Environment string `uri:"environment,omitempty" url:"environment,omitempty"`
	Machine     string `uri:"machine,omitempty" url:"machine,omitempty"`
	Project     string `uri:"project,omitempty" url:"project,omitempty"`
	Role        string `uri:"role,omitempty" url:"role,omitempty"`
	Runbook     string `uri:"runbook,omitempty" url:"runbook,omitempty"`
	Tenant      string `uri:"tenant,omitempty" url:"tenant,omitempty"`
}

type MissingVariablesQuery struct {
	EnvironmentID  string `uri:"environmentId,omitempty" url:"environmentId,omitempty"`
	IncludeDetails bool   `uri:"includeDetails,omitempty" url:"includeDetails,omitempty"`
	ProjectID      string `uri:"projectId,omitempty" url:"projectId,omitempty"`
	TenantID       string `uri:"tenantId,omitempty" url:"tenantId,omitempty"`
}

type VariablesQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}

type VersionRuleTestQuery struct {
	FeetType      string `uri:"feetType,omitempty" url:"feetType,omitempty"`
	PreReleaseTag string `uri:"preReleaseTag,omitempty" url:"preReleaseTag,omitempty"`
	Version       string `uri:"version,omitempty" url:"version,omitempty"`
	VersionRange  string `uri:"versionRange,omitempty" url:"versionRange,omitempty"`
}

type WorkerPoolsQuery struct {
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	Name        string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
}

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	HealthStatuses       []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	HideEmptyWorkerPools bool     `uri:"hideEmptyWorkerPools,omitempty" url:"hideEmptyWorkerPools,omitempty"`
	IDs                  []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled           bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	MachinePartialName   string   `uri:"machinePartialName,omitempty" url:"machinePartialName,omitempty"`
	PartialName          string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ShellNames           []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
}

type WorkersQuery struct {
	CommunicationStyles []string `uri:"commStyles,omitempty" url:"commStyles,omitempty"`
	HealthStatuses      []string `uri:"healthStatuses,omitempty" url:"healthStatuses,omitempty"`
	IDs                 []string `uri:"ids,omitempty" url:"ids,omitempty"`
	IsDisabled          bool     `uri:"isDisabled,omitempty" url:"isDisabled,omitempty"`
	Name                string   `uri:"name,omitempty" url:"name,omitempty"`
	PartialName         string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	ShellNames          []string `uri:"shellNames,omitempty" url:"shellNames,omitempty"`
	Skip                int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take                int      `uri:"take,omitempty" url:"take,omitempty"`
	Thumbprint          string   `uri:"thumbprint,omitempty" url:"thumbprint,omitempty"`
	WorkerPoolIDs       []string `uri:"workerPoolIds" url:"workerPoolIds,omitempty"`
}
