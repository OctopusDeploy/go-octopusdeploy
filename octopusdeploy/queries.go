package octopusdeploy

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType AccountType `url:"accountType,omitempty"`
	IDs         []string    `url:"ids,omitempty"`
	PartialName string      `url:"partialName,omitempty"`
	Skip        int         `url:"skip,omitempty"`
	Take        int         `url:"take,omitempty"`
}

type ActionTemplateLogoQuery struct {
	CB       string `uri:"cb"`
	TypeOrID string `uri:"typeOrId"`
}

// ActionTemplatesQuery represents parameters to query the ActionTemplates service.
type ActionTemplatesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type ActionTemplateVersionedLogoQuery struct {
	CB       string `uri:"cb"`
	TypeOrID string `uri:"typeOrId"`
	Version  string `uri:"version"`
}

type APIQuery struct {
	Skip int `uri:"skip"`
	Take int `uri:"take"`
}

type ArtifactsQuery struct {
	IDs         []string `uri:"ids"`
	Order       string   `uri:"order"`
	PartialName string   `uri:"partialName"`
	Regarding   string   `uri:"regarding"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type BuildInformationQuery struct {
	Filter        string `uri:"filter"`
	Latest        string `uri:"latest"`
	OverwriteMode string `uri:"overwriteMode"`
	PackageID     string `uri:"packageId"`
	Skip          int    `uri:"skip"`
	Take          int    `uri:"take"`
}

type BuildInformationBulkQuery struct {
	IDs []string `uri:"ids"`
}

type CertificateConfigurationQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type CertificatesQuery struct {
	Archived    string   `uri:"archived"`
	FirstResult string   `uri:"firstResult"`
	IDs         []string `uri:"ids"`
	OrderBy     string   `uri:"orderBy"`
	PartialName string   `uri:"partialName"`
	Search      string   `uri:"search"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
	Tenant      string   `uri:"tenant"`
}

type ChannelsQuery struct {
	IDs         []string `url:"ids,omitempty"`
	PartialName string   `url:"partialName,omitempty"`
	Skip        int      `url:"skip,omitempty"`
	Take        int      `url:"take,omitempty"`
}

type CloudTemplateQuery struct {
	FeedID    string `uri:"feedId"`
	PackageID string `uri:"packageId"`
}

type CommunityActionTemplatesQuery struct {
	IDs  []string `uri:"ids"`
	Skip int      `uri:"skip"`
	Take int      `uri:"take"`
}

type DashboardQuery struct {
	IncludeLatest   bool     `uri:"highestLatestVersionPerProjectAndEnvironment"`
	ProjectID       string   `uri:"projectId"`
	SelectedTags    []string `uri:"selectedTags"`
	SelectedTenants []string `uri:"selectedTenants"`
	ShowAll         bool     `uri:"showAll"`
	ReleaseID       string   `uri:"releaseId"`
}

type DashboardDynamicQuery struct {
	Environments    []string `uri:"environments"`
	IncludePrevious bool     `uri:"includePrevious"`
	Projects        []string `uri:"projects"`
}

type DeploymentProcessesQuery struct {
	IDs  []string `uri:"ids"`
	Skip int      `uri:"skip"`
	Take int      `uri:"take"`
}

type DeploymentQuery struct {
	Skip int `uri:"skip"`
	Take int `uri:"take"`
}

type DeploymentsQuery struct {
	Channels     string   `uri:"channels"`
	Environments []string `uri:"environments"`
	IDs          []string `uri:"ids"`
	PartialName  string   `uri:"partialName"`
	Projects     []string `uri:"projects"`
	Skip         int      `uri:"skip"`
	Take         int      `uri:"take"`
	TaskState    string   `uri:"taskState"`
	Tenants      []string `uri:"tenants"`
}

type DiscoverMachineQuery struct {
	Host    string `uri:"host"`
	Port    int    `uri:"port"`
	ProxyID string `uri:"proxyId"`
	Type    string `uri:"type"`
}

type DiscoverWorkerQuery struct {
	Host    string `uri:"host"`
	Port    int    `uri:"port"`
	ProxyID string `uri:"proxyId"`
	Type    string `uri:"type"`
}

type EnvironmentsQuery struct {
	IDs         []string `uri:"ids"`
	Name        string   `uri:"name"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type EnvironmentsSummaryQuery struct {
	CommunicationStyles   []string `uri:"commStyles"`
	HealthStatuses        []string `uri:"healthStatuses"`
	HideEmptyEnvironments bool     `uri:"hideEmptyEnvironments"`
	IDs                   []string `uri:"ids"`
	IsDisabled            bool     `uri:"isDisabled"`
	MachinePartialName    string   `uri:"machinePartialName"`
	PartialName           string   `uri:"partialName"`
	Roles                 []string `uri:"roles"`
	ShellNames            []string `uri:"shellNames"`
	TenantIDs             []string `uri:"tenantIds"`
	TenantTags            []string `uri:"tenantTags"`
}

type EventCategoriesQuery struct {
	AppliesTo string `uri:"appliesTo"`
}

type EventGroupsQuery struct {
	AppliesTo string `uri:"appliesTo"`
}

type EventsQuery struct {
	AsCSV             string   `uri:"asCsv"`
	DocumentTypes     []string `uri:"documentTypes"`
	Environments      []string `uri:"environments"`
	EventAgents       []string `uri:"eventAgents"`
	EventCategories   []string `uri:"eventCategories"`
	EventGroups       []string `uri:"eventGroups"`
	ExcludeDifference bool     `uri:"excludeDifference"`
	From              string   `uri:"from"`
	FromAutoID        string   `uri:"fromAutoId"`
	IDs               []string `uri:"ids"`
	IncludeSystem     bool     `uri:"includeSystem"`
	Internal          string   `uri:"interal"`
	Name              string   `uri:"name"`
	PartialName       string   `uri:"partialName"`
	ProjectGroups     []string `uri:"projectGroups"`
	Projects          []string `uri:"projects"`
	Regarding         string   `uri:"regarding"`
	RegardingAny      string   `uri:"regardingAny"`
	Skip              int      `uri:"skip"`
	Spaces            []string `uri:"spaces"`
	Take              int      `uri:"take"`
	Tags              []string `uri:"tags"`
	Tenants           []string `uri:"tenants"`
	To                string   `uri:"to"`
	ToAutoID          string   `uri:"toAutoId"`
	User              string   `uri:"user"`
	Users             []string `uri:"users"`
}

type ExternalUserSearchQuery struct {
	PartialName string `uri:"partialName"`
}

type FeedsQuery struct {
	FeedType    string   `uri:"feedType"`
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type InterruptionsQuery struct {
	IDs         []string `uri:"ids"`
	PendingOnly bool     `uri:"pendingOnly"`
	Regarding   string   `uri:"regarding"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type IssueTrackersQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type LibraryVariablesQuery struct {
	ContentType string   `uri:"contentType"`
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type LifecyclesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type MachinePoliciesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type MachinesQuery struct {
	CommunicationStyles []string `uri:"commStyles"`
	DeploymentID        string   `uri:"deploymentId"`
	EnvironmentIDs      []string `uri:"environmentIds"`
	HealthStatuses      []string `uri:"healthStatuses"`
	IDs                 []string `uri:"ids"`
	IsDisabled          bool     `uri:"isDisabled"`
	Name                string   `uri:"name"`
	PartialName         string   `uri:"partialName"`
	Roles               []string `uri:"roles"`
	ShellNames          []string `uri:"shellNames"`
	Skip                int      `uri:"skip"`
	Take                int      `uri:"take"`
	TenantIDs           []string `uri:"tenantIds"`
	TenantTags          []string `uri:"tenantTags"`
	Thumbprint          string   `uri:"thumbprint"`
}

type OctopusServerNodesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type PackageDeltaSignatureQuery struct {
	PackageID string `uri:"packageId"`
	Version   string `uri:"version"`
}

type PackageDeltaUploadQuery struct {
	BaseVersion   string `uri:"baseVersion"`
	OverwriteMode string `uri:"overwriteMode"`
	PackageID     string `uri:"packageId"`
	Replace       bool   `uri:"replace"`
}

type PackageMetadataQuery struct {
	Filter        string `uri:"filter"`
	Latest        string `uri:"latest"`
	OverwriteMode string `uri:"overwriteMode"`
	Replace       bool   `uri:"replace"`
	Skip          int    `uri:"skip"`
	Take          int    `uri:"take"`
}

type PackageNotesListQuery struct {
	PackageIDs []string `uri:"packageIds"`
}

type PackagesQuery struct {
	Filter         string `uri:"filter"`
	IncludeNotes   bool   `uri:"includeNotes"`
	Latest         string `uri:"latest"`
	NuGetPackageID string `uri:"nuGetPackageId"`
	Skip           int    `uri:"skip"`
	Take           int    `uri:"take"`
}

type PackagesBulkQuery struct {
	IDs []string `uri:"ids"`
}

type PackageUploadQuery struct {
	Replace       bool   `uri:"replace"`
	OverwriteMode string `uri:"overwriteMode"`
}

type UserQuery struct {
	IncludeSystem bool     `uri:"includeSystem"`
	Spaces        []string `uri:"spaces"`
}

type ProjectGroupsQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type ProjectPulseQuery struct {
	ProjectIDs []string `uri:"projectIds"`
}

type ProjectsQuery struct {
	ClonedFromProjectID string   `uri:"clonedFromProjectId"`
	IDs                 []string `uri:"ids"`
	IsClone             bool     `uri:"clone"`
	Name                string   `uri:"name"`
	PartialName         string   `uri:"partialName"`
	Skip                int      `uri:"skip"`
	Take                int      `uri:"take"`
}

type ProjectsExperimentalSummariesQuery struct {
	IDs []string `uri:"ids"`
}

type ProjectTriggersQuery struct {
	IDs      []string `uri:"ids"`
	Runbooks []string `uri:"runbooks"`
	Skip     int      `uri:"skip"`
	Take     int      `uri:"take"`
}

type ProxiesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type ReleaseQuery struct {
	SearchByVersion string `uri:"searchByVersion"`
	Skip            int    `uri:"skip"`
	Take            int    `uri:"take"`
}

type ReleasesQuery struct {
	IDs                []string `uri:"ids"`
	IgnoreChannelRules bool     `uri:"ignoreChannelRules"`
	Skip               int      `uri:"skip"`
	Take               int      `uri:"take"`
}

type RunbookProcessesQuery struct {
	IDs  []string `uri:"ids"`
	Skip int      `uri:"skip"`
	Take int      `uri:"take"`
}

type RunbookRunsQuery struct {
	Environments []string `uri:"environments"`
	IDs          []string `uri:"ids"`
	PartialName  string   `uri:"partialName"`
	Projects     []string `uri:"projects"`
	Runbooks     []string `uri:"runbooks"`
	Skip         int      `uri:"skip"`
	Take         int      `uri:"take"`
	TaskState    string   `uri:"taskState"`
	Tenants      []string `uri:"tenants"`
}

type RunbooksQuery struct {
	IDs         []string `uri:"ids"`
	IsClone     bool     `uri:"clone"`
	PartialName string   `uri:"partialName"`
	ProjectIDs  []string `uri:"projectIds"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type RunbookSnapshotsQuery struct {
	IDs     []string `uri:"ids"`
	Publish bool     `uri:"publish"`
	Skip    int      `uri:"skip"`
	Take    int      `uri:"take"`
}

type ScheduledProjectTriggersQuery struct {
	IDs  []string `uri:"ids"`
	Skip int      `uri:"skip"`
	Take int      `uri:"take"`
}

type SchedulerQuery struct {
	Verbose bool   `uri:"verbose"`
	Tail    string `uri:"tail"`
}

type ScopedUserRolesQuery struct {
	IDs           []string `uri:"ids"`
	IncludeSystem bool     `uri:"includeSystem"`
	PartialName   string   `uri:"partialName"`
	Skip          int      `uri:"skip"`
	Spaces        []string `uri:"spaces"`
	Take          int      `uri:"take"`
}

type SignInQuery struct {
	ReturnURL string `uri:"returnUrl"`
}

type SpaceHomeQuery struct {
	SpaceID string `uri:"spaceId"`
}

type SpacesQuery struct {
	IDs         []string `uri:"ids"`
	Name        string   `uri:"name"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type SubscriptionsQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Spaces      []string `uri:"spaces"`
	Take        int      `uri:"take"`
}

type TagSetsQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type TasksQuery struct {
	Environment             string   `uri:"environment"`
	HasPendingInterruptions bool     `uri:"hasPendingInterruptions"`
	HasWarningsOrErrors     bool     `uri:"hasWarningsOrErrors"`
	IDs                     []string `uri:"ids"`
	IncludeSystem           bool     `uri:"includeSystem"`
	IsActive                bool     `uri:"active"`
	IsRunning               bool     `uri:"running"`
	Name                    string   `uri:"name"`
	Node                    string   `uri:"node"`
	PartialName             string   `uri:"partialName"`
	Project                 string   `uri:"project"`
	Runbook                 string   `uri:"runbook"`
	Skip                    int      `uri:"skip"`
	Spaces                  []string `uri:"spaces"`
	States                  []string `uri:"states"`
	Take                    int      `uri:"take"`
	Tenant                  string   `uri:"tenant"`
}

type TeamMembershipQuery struct {
	IncludeSystem bool     `uri:"includeSystem"`
	Spaces        []string `uri:"spaces"`
	UserID        string   `uri:"userId"`
}

type TeamsQuery struct {
	IDs           []string `uri:"ids"`
	IncludeSystem bool     `uri:"includeSystem"`
	PartialName   string   `uri:"partialName"`
	Skip          int      `uri:"skip"`
	Spaces        []string `uri:"spaces"`
	Take          int      `uri:"take"`
}

type TenantsQuery struct {
	ClonedFromTenantID string   `uri:"clonedFromTenantId"`
	IDs                []string `uri:"ids"`
	IsClone            bool     `uri:"clone"`
	Name               string   `uri:"name"`
	PartialName        string   `uri:"partialName"`
	ProjectID          string   `uri:"projectId"`
	Skip               int      `uri:"skip"`
	Tags               []string `uri:"tags"`
	Take               int      `uri:"take"`
}

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `uri:"environmentId"`
	IncludeDetails bool     `uri:"includeDetails"`
	ProjectID      string   `uri:"projectId"`
	TenantID       string   `uri:"tenantId"`
}

type TenantVariablesQuery struct {
	ProjectID string `uri:"projectId"`
}

type UserRolesQuery struct {
	IDs         []string `uri:"ids"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type UsersQuery struct {
	Filter string   `uri:"filter"`
	IDs    []string `uri:"ids"`
	Skip   int      `uri:"skip"`
	Take   int      `uri:"take"`
}

type VariableNamesQuery struct {
	Project                   string `uri:"project"`
	ProjectEnvironmentsFilter string `struct:"projectEnvironmentsFilter"`
	Runbook                   string `uri:"runbook"`
}

type VariablePreviewQuery struct {
	Action      string `uri:"action"`
	Channel     string `uri:"channel"`
	Environment string `uri:"environment"`
	Machine     string `uri:"machine"`
	Project     string `uri:"project"`
	Role        string `uri:"role"`
	Runbook     string `uri:"runbook"`
	Tenant      string `uri:"tenant"`
}

type VariablesQuery struct {
	IDs []string `uri:"ids"`
}

type VersionRuleTestQuery struct {
	FeetType      string `uri:"feetType"`
	PreReleaseTag string `uri:"preReleaseTag"`
	Version       string `uri:"version"`
	VersionRange  string `uri:"versionRange"`
}

type WorkerPoolsQuery struct {
	IDs         []string `uri:"ids"`
	Name        string   `uri:"name"`
	PartialName string   `uri:"partialName"`
	Skip        int      `uri:"skip"`
	Take        int      `uri:"take"`
}

type WorkerPoolsSummaryQuery struct {
	CommunicationStyles  []string `uri:"commStyles"`
	HealthStatuses       []string `uri:"healthStatuses"`
	HideEmptyWorkerPools bool     `uri:"hideEmptyWorkerPools"`
	IDs                  []string `uri:"ids"`
	IsDisabled           bool     `uri:"isDisabled"`
	MachinePartialName   string   `uri:"machinePartialName"`
	PartialName          string   `uri:"partialName"`
	ShellNames           []string `uri:"shellNames"`
}

type WorkersQuery struct {
	CommunicationStyles []string `uri:"commStyles"`
	HealthStatuses      []string `uri:"healthStatuses"`
	IDs                 []string `uri:"ids"`
	IsDisabled          bool     `uri:"isDisabled"`
	Name                string   `uri:"name"`
	PartialName         string   `uri:"partialName"`
	ShellNames          []string `uri:"shellNames"`
	Skip                int      `uri:"skip"`
	Take                int      `uri:"take"`
	Thumbprint          string   `uri:"thumbprint"`
	WorkerPoolIDs       []string `uri:"workerPoolIds"`
}
