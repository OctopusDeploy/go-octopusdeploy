package dashboard

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// Dashboard represents the deployment overview returned by the dashboard and
// dynamic dashboard endpoints: the deployments themselves in Items, plus the
// reference data needed to resolve the IDs those items carry.
type Dashboard struct {
	Items         []*DashboardItem         `json:"Items"`
	Projects      []*DashboardProject      `json:"Projects"`
	ProjectGroups []*DashboardProjectGroup `json:"ProjectGroups"`
	Environments  []*DashboardEnvironment  `json:"Environments"`
	Tenants       []*DashboardTenant       `json:"Tenants"`

	// ProjectLimit is set when the server caps how many projects the dashboard
	// reports on, and is nil when no cap applies. Callers showing a full list
	// should say so rather than presenting capped results as complete.
	ProjectLimit *int `json:"ProjectLimit"`

	// IsFiltered reports whether the server narrowed the response, either from
	// the query or from the caller's permissions.
	IsFiltered bool `json:"IsFiltered"`

	resources.Resource
}

// DashboardProject is the subset of a project the dashboard returns to
// accompany its items.
type DashboardProject struct {
	Name                           string   `json:"Name,omitempty"`
	Slug                           string   `json:"Slug,omitempty"`
	ProjectGroupID                 string   `json:"ProjectGroupId,omitempty"`
	EnvironmentIDs                 []string `json:"EnvironmentIds,omitempty"`
	TenantedDeploymentMode         string   `json:"TenantedDeploymentMode,omitempty"`
	CanPerformUntenantedDeployment bool     `json:"CanPerformUntenantedDeployment"`
	IsDisabled                     bool     `json:"IsDisabled"`

	resources.Resource
}

// DashboardEnvironment is the subset of an environment the dashboard returns to
// accompany its items.
type DashboardEnvironment struct {
	Name string `json:"Name,omitempty"`

	resources.Resource
}

// DashboardProjectGroup is the subset of a project group the dashboard returns
// to accompany its items.
type DashboardProjectGroup struct {
	Name           string   `json:"Name,omitempty"`
	EnvironmentIDs []string `json:"EnvironmentIds,omitempty"`

	resources.Resource
}

// DashboardTenant is the subset of a tenant the dashboard returns to accompany
// its items.
type DashboardTenant struct {
	Name                string              `json:"Name,omitempty"`
	TenantTags          []string            `json:"TenantTags,omitempty"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	IsDisabled          bool                `json:"IsDisabled"`

	resources.Resource
}
