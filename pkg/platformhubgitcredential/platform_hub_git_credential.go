package platformhubgitcredential

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// PlatformHubGitCredential represents a Platform Hub git credential resource.
type PlatformHubGitCredential struct {
	Description            string                              `json:"Description,omitempty"`
	Details                credentials.GitCredential           `json:"Details"`
	Name                   string                              `json:"Name"`
	RepositoryRestrictions *credentials.RepositoryRestrictions `json:"RepositoryRestrictions"`
	resources.Resource
}

// NewPlatformHubGitCredential creates and initializes a Platform Hub git credential resource.
func NewPlatformHubGitCredential(name string, credential credentials.GitCredential) *PlatformHubGitCredential {
	return &PlatformHubGitCredential{
		Details:  credential,
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

// GetName returns the name of the resource.
func (r *PlatformHubGitCredential) GetName() string {
	return r.Name
}

// SetName sets the name of the resource.
func (r *PlatformHubGitCredential) SetName(name string) {
	r.Name = name
}

// UnmarshalJSON sets the resource to its representation in JSON.
func (r *PlatformHubGitCredential) UnmarshalJSON(b []byte) error {
	var fields struct {
		Description string `json:"Description,omitempty"`
		Name        string `json:"Name"`
		resources.Resource
	}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	r.Description = fields.Description
	r.Name = fields.Name
	r.Resource = fields.Resource

	var rawResource map[string]*json.RawMessage
	if err := json.Unmarshal(b, &rawResource); err != nil {
		return err
	}

	var repositoryRestrictions credentials.RepositoryRestrictions
	if rawResource["RepositoryRestrictions"] != nil {
		restrictionsValue := rawResource["RepositoryRestrictions"]

		if err := json.Unmarshal(*restrictionsValue, &repositoryRestrictions); err != nil {
			return err
		}
		r.RepositoryRestrictions = &repositoryRestrictions
	}

	var gitCredentials *json.RawMessage

	if rawResource["Details"] != nil {
		detailsValue := rawResource["Details"]

		if err := json.Unmarshal(*detailsValue, &gitCredentials); err != nil {
			return err
		}

		var usernamePasswordGitCredential *credentials.UsernamePassword
		if err := json.Unmarshal(*gitCredentials, &usernamePasswordGitCredential); err != nil {
			return err
		}
		r.Details = usernamePasswordGitCredential
	}

	return nil
}

var _ resources.IHasName = &PlatformHubGitCredential{}
