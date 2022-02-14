package resources

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"net/url"

	"github.com/go-playground/validator/v10"
)

// GitPersistenceSettings represents persistence settings associated with a project.
type GitPersistenceSettings struct {
	BasePath      string
	Credentials   IGitCredential
	DefaultBranch string
	URL           *url.URL

	octopusdeploy.persistenceSettings
}

// NewGitPersistenceSettings creates an instance of persistence settings.
func NewGitPersistenceSettings(
	basePath string,
	credentials IGitCredential,
	defaultBranch string,
	url *url.URL) *GitPersistenceSettings {
	return &GitPersistenceSettings{
		BasePath:            basePath,
		Credentials:         credentials,
		DefaultBranch:       defaultBranch,
		URL:                 url,
		persistenceSettings: octopusdeploy.persistenceSettings{Type: "VersionControlled"},
	}
}

// GetType returns the type for this persistence settings.
func (g *GitPersistenceSettings) GetType() string {
	return g.Type
}

// MarshalJSON returns persistence settings as its JSON encoding.
func (p *GitPersistenceSettings) MarshalJSON() ([]byte, error) {
	persistenceSettings := struct {
		BasePath      string         `json:"BasePath,omitempty"`
		Credentials   IGitCredential `json:"Credentials,omitempty"`
		DefaultBranch string         `json:"DefaultBranch,omitempty"`
		URL           string         `json:"Url,omitempty"`
		octopusdeploy.persistenceSettings
	}{
		BasePath:            p.BasePath,
		Credentials:         p.Credentials,
		DefaultBranch:       p.DefaultBranch,
		URL:                 p.URL.String(),
		persistenceSettings: p.persistenceSettings,
	}

	return json.Marshal(persistenceSettings)
}

// UnmarshalJSON sets the persistence settings to its representation in JSON.
func (p *GitPersistenceSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		BasePath      string `json:"BasePath,omitempty"`
		DefaultBranch string `json:"DefaultBranch,omitempty"`
		URL           string `json:"Url,omitempty"`
		octopusdeploy.persistenceSettings
	}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	if err := validate.Struct(fields); err != nil {
		return err
	}

	var url *url.URL
	if len(fields.URL) > 0 {
		url, err = url.Parse(fields.URL)
		if err != nil {
			return err
		}
	}

	p.BasePath = fields.BasePath
	p.DefaultBranch = fields.DefaultBranch
	p.Type = fields.Type
	p.URL = url
	p.persistenceSettings = fields.persistenceSettings

	var persistenceSettings map[string]*json.RawMessage
	err = json.Unmarshal(b, &persistenceSettings)
	if err != nil {
		return err
	}

	var credentials *json.RawMessage
	var credentialsProperties map[string]*json.RawMessage
	var gitCredentialType string

	if persistenceSettings["Credentials"] != nil {
		credentialsValue := persistenceSettings["Credentials"]

		err = json.Unmarshal(*credentialsValue, &credentials)
		if err != nil {
			return err
		}

		err = json.Unmarshal(*credentials, &credentialsProperties)
		if err != nil {
			return err
		}

		if credentialsProperties["Type"] != nil {
			t := credentialsProperties["Type"]
			json.Unmarshal(*t, &gitCredentialType)
		}
	}

	switch gitCredentialType {
	case "Anonymous":
		var anonymousGitCredential *AnonymousGitCredential
		err := json.Unmarshal(*credentials, &anonymousGitCredential)
		if err != nil {
			return err
		}
		p.Credentials = anonymousGitCredential
	case "UsernamePassword":
		var usernamePasswordGitCredential *UsernamePasswordGitCredential
		err := json.Unmarshal(*credentials, &usernamePasswordGitCredential)
		if err != nil {
			return err
		}
		p.Credentials = usernamePasswordGitCredential
	}

	return nil
}

var _ IPersistenceSettings = &GitPersistenceSettings{}
