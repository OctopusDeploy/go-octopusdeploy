package credentials

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Resource struct {
	Description            string                  `json:"Description,omitempty"`
	Details                GitCredential           `json:"Details"`
	Name                   string                  `json:"Name"`
	SpaceID                string                  `json:"SpaceId,omitempty"`
	RepositoryRestrictions *RepositoryRestrictions `json:"RepositoryRestrictions"`
	resources.Resource
}

type RepositoryRestrictions struct {
	Enabled             bool     `json:"Enabled"`
	AllowedRepositories []string `json:"AllowedRepositories"`
}

func NewResource(name string, credential GitCredential) *Resource {
	return &Resource{
		Details:  credential,
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

// GetName returns the name of the resource.
func (r *Resource) GetName() string {
	return r.Name
}

// SetName sets the name of the resource.
func (r *Resource) SetName(name string) {
	r.Name = name
}

// UnmarshalJSON sets the resource to its representation in JSON.
func (r *Resource) UnmarshalJSON(b []byte) error {
	var fields struct {
		Description string `json:"Description,omitempty"`
		Name        string `json:"Name"`
		SpaceID     string `json:"SpaceId,omitempty"`
		resources.Resource
	}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	r.Description = fields.Description
	r.Name = fields.Name
	r.SpaceID = fields.SpaceID
	r.Resource = fields.Resource

	var rawResource map[string]*json.RawMessage
	if err := json.Unmarshal(b, &rawResource); err != nil {
		return err
	}

	var repositoryRestrictions RepositoryRestrictions
	if rawResource["RepositoryRestrictions"] != nil {
		restrictionsValue := rawResource["RepositoryRestrictions"]

		if err := json.Unmarshal(*restrictionsValue, &repositoryRestrictions); err != nil {
			return err
		}
		r.RepositoryRestrictions = &repositoryRestrictions
	}

	var gitCredentials *json.RawMessage
	var credentialsProperties map[string]*json.RawMessage
	var gitCredentialType Type

	if rawResource["Details"] != nil {
		detailsValue := rawResource["Details"]

		if err := json.Unmarshal(*detailsValue, &gitCredentials); err != nil {
			return err
		}

		if err := json.Unmarshal(*gitCredentials, &credentialsProperties); err != nil {
			return err
		}

		if credentialsProperties["Type"] != nil {
			t := credentialsProperties["Type"]
			json.Unmarshal(*t, &gitCredentialType)
		}
	}

	switch gitCredentialType {
	case GitCredentialTypeAnonymous:
		var anonymousGitCredential *Anonymous
		if err := json.Unmarshal(*gitCredentials, &anonymousGitCredential); err != nil {
			return err
		}
		r.Details = anonymousGitCredential
	case GitCredentialTypeReference:
		var referenceProjectGitCredential *Reference
		if err := json.Unmarshal(*gitCredentials, &referenceProjectGitCredential); err != nil {
			return err
		}
		r.Details = referenceProjectGitCredential
	case GitCredentialTypeUsernamePassword:
		var usernamePasswordGitCredential *UsernamePassword
		if err := json.Unmarshal(*gitCredentials, &usernamePasswordGitCredential); err != nil {
			return err
		}
		r.Details = usernamePasswordGitCredential
	}

	return nil
}

var _ resources.IHasName = &Resource{}
