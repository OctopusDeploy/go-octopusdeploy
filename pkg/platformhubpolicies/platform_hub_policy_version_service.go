package platformhubpolicies

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const publishTemplate = "/api/platformhub/{gitRef}/policies/{slug}/publish"

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
	path, pathError := client.URITemplateCache().Expand(publishTemplate, map[string]any{"gitRef": gitRef, "slug": slug})
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
