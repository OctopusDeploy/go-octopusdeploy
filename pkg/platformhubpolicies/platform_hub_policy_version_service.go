package platformhubpolicies

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

// PlatformHubPolicyVersion represents an immutable published version of a Platform Hub policy.
type PlatformHubPolicyVersion struct {
	ID              string    `json:"Id"`
	Slug            string    `json:"Slug"`
	Version         string    `json:"Version"`
	PublishedDate   time.Time `json:"PublishedDate"`
	GitRef          string    `json:"GitRef"`
	GitCommit       string    `json:"GitCommit"`
	Name            string    `json:"Name"`
	Description     string    `json:"Description,omitempty"`
	ViolationReason string    `json:"ViolationReason,omitempty"`
	ViolationAction string    `json:"ViolationAction"`
	RegoScope       string    `json:"RegoScope"`
	RegoConditions  string    `json:"RegoConditions"`
	IsActive        bool      `json:"IsActive"`
}

// Publish publishes a Platform Hub policy version.
func Publish(client newclient.Client, gitRef string, slug string, version string) (PlatformHubPolicyVersion, error) {
	path, pathError := client.URITemplateCache().Expand("/api/platformhub/{gitRef}/policies/{slug}/publish", map[string]any{"gitRef": gitRef, "slug": slug})
	if pathError != nil {
		return PlatformHubPolicyVersion{}, pathError
	}

	body := struct {
		Version string `json:"Version"`
	}{Version: version}

	publishedVersion, err := newclient.Post[PlatformHubPolicyVersion](client.HttpSession(), path, body)
	if err != nil {
		return PlatformHubPolicyVersion{}, err
	}

	return *publishedVersion, nil
}

// VersionsQuery represents query parameters for listing policy versions.
type VersionsQuery struct {
	Slug string `uri:"slug"`
	Skip int    `uri:"skip,omitempty"`
	Take int    `uri:"take,omitempty"`
}

// GetVersions returns published versions of a Platform Hub policy.
func GetVersions(client newclient.Client, query VersionsQuery) ([]PlatformHubPolicyVersion, error) {
	path, pathError := client.URITemplateCache().Expand("/api/platformhub/policies/{slug}/versions{?skip,take}", query)
	if pathError != nil {
		return nil, pathError
	}

	versions, err := newclient.Get[[]PlatformHubPolicyVersion](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return *versions, nil
}

// ActivateVersion activates a published Platform Hub policy version.
func ActivateVersion(client newclient.Client, version PlatformHubPolicyVersion) (PlatformHubPolicyVersion, error) {
	return modifyVersionStatus(client, version, true)
}

// DeactivateVersion deactivates a published Platform Hub policy version.
func DeactivateVersion(client newclient.Client, version PlatformHubPolicyVersion) (PlatformHubPolicyVersion, error) {
	return modifyVersionStatus(client, version, false)
}

func modifyVersionStatus(client newclient.Client, version PlatformHubPolicyVersion, isActive bool) (PlatformHubPolicyVersion, error) {
	path, pathError := client.URITemplateCache().Expand("/api/platformhub/policies/{slug}/versions/{version}/modify-status", map[string]any{
		"slug":    version.Slug,
		"version": version.Version,
	})
	if pathError != nil {
		return version, pathError
	}

	body := struct {
		IsActive bool `json:"IsActive"`
	}{IsActive: isActive}

	modifiedVersion, err := newclient.Post[PlatformHubPolicyVersion](client.HttpSession(), path, body)
	if err != nil {
		return version, err
	}

	return *modifiedVersion, nil
}
