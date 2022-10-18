package projects

import (
	"encoding/json"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
)

// VersionControlSettings represents version control settings associated with a project.
type VersionControlSettings struct {
	BasePath      string
	Credentials   credentials.IGitCredential
	DefaultBranch string
	Type          string
	URL           *url.URL
}

// NewVersionControlSettings creates an instance of version control settings.
func NewVersionControlSettings(
	basePath string,
	credentials credentials.IGitCredential,
	defaultBranch string,
	persistenceType string,
	url *url.URL) *VersionControlSettings {
	return &VersionControlSettings{
		BasePath:      basePath,
		Credentials:   credentials,
		DefaultBranch: defaultBranch,
		Type:          persistenceType,
		URL:           url,
	}
}

// MarshalJSON returns version control settings as its JSON encoding.
func (v *VersionControlSettings) MarshalJSON() ([]byte, error) {
	versionControlSettings := struct {
		BasePath      string                     `json:"BasePath,omitempty"`
		Credentials   credentials.IGitCredential `json:"Credentials,omitempty"`
		DefaultBranch string                     `json:"DefaultBranch,omitempty"`
		Type          string                     `json:"Type"`
		Url           string                     `json:"Url,omitempty"`
	}{
		BasePath:      v.BasePath,
		Credentials:   v.Credentials,
		DefaultBranch: v.DefaultBranch,
		Type:          string(v.Type),
		Url:           v.URL.String(),
	}

	return json.Marshal(versionControlSettings)
}

// UnmarshalJSON sets the version control settings to its representation in JSON.
func (v *VersionControlSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		BasePath      string `json:"BasePath,omitempty"`
		DefaultBranch string `json:"DefaultBranch,omitempty"`
		Type          string `json:"Type"`
		Url           string `json:"Url,omitempty"`
	}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	// return error if unable to parse URL
	url, err := url.Parse(fields.Url)
	if err != nil {
		return err
	}

	v.BasePath = fields.BasePath
	v.DefaultBranch = fields.DefaultBranch
	v.Type = fields.Type
	v.URL = url

	var versionControlSettings map[string]*json.RawMessage
	err = json.Unmarshal(b, &versionControlSettings)
	if err != nil {
		return err
	}

	if versionControlSettings["Credentials"] == nil {
		return nil
	}

	var gitCredentials *json.RawMessage
	var credentialsProperties map[string]*json.RawMessage
	var gitCredentialType string

	credentialsValue := versionControlSettings["Credentials"]

	err = json.Unmarshal(*credentialsValue, &gitCredentials)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*gitCredentials, &credentialsProperties)
	if err != nil {
		return err
	}

	if credentialsProperties["Type"] != nil {
		t := credentialsProperties["Type"]
		json.Unmarshal(*t, &gitCredentialType)
	}

	switch gitCredentialType {
	case "Anonymous":
		var anonymousGitCredential *credentials.Anonymous
		err := json.Unmarshal(*gitCredentials, &anonymousGitCredential)
		if err != nil {
			return err
		}
		v.Credentials = anonymousGitCredential
	case "UsernamePassword":
		var usernamePasswordGitCredential *credentials.UsernamePassword
		err := json.Unmarshal(*gitCredentials, &usernamePasswordGitCredential)
		if err != nil {
			return err
		}
		v.Credentials = usernamePasswordGitCredential
	}

	return nil
}
