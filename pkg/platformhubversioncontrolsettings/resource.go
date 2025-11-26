package platformhubversioncontrolsettings

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
)

// Resource represents Platform Hub version control settings
type Resource struct {
	URL           string                    `json:"Url"`
	Credentials   credentials.GitCredential `json:"Credentials"`
	DefaultBranch string                    `json:"DefaultBranch"`
	BasePath      string                    `json:"BasePath"`
}

// NewResource creates a new Platform Hub version control settings resource
func NewResource(url string, creds credentials.GitCredential, defaultBranch string, basePath string) *Resource {
	return &Resource{
		URL:           url,
		Credentials:   creds,
		DefaultBranch: defaultBranch,
		BasePath:      basePath,
	}
}

// UnmarshalJSON sets the resource to its representation in JSON
func (r *Resource) UnmarshalJSON(b []byte) error {
	var fields struct {
		URL           string `json:"Url"`
		DefaultBranch string `json:"DefaultBranch"`
		BasePath      string `json:"BasePath"`
	}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	r.URL = fields.URL
	r.DefaultBranch = fields.DefaultBranch
	r.BasePath = fields.BasePath

	var rawResource map[string]*json.RawMessage
	if err := json.Unmarshal(b, &rawResource); err != nil {
		return err
	}

	var credentialsRaw *json.RawMessage
	var credentialsProperties map[string]*json.RawMessage
	var credentialType credentials.Type

	if rawResource["Credentials"] != nil {
		credentialsValue := rawResource["Credentials"]

		if err := json.Unmarshal(*credentialsValue, &credentialsRaw); err != nil {
			return err
		}

		if err := json.Unmarshal(*credentialsRaw, &credentialsProperties); err != nil {
			return err
		}

		if credentialsProperties["Type"] != nil {
			t := credentialsProperties["Type"]
			json.Unmarshal(*t, &credentialType)
		}
	}

	switch credentialType {
	case credentials.GitCredentialTypeAnonymous:
		var anonymousCredential *credentials.Anonymous
		if err := json.Unmarshal(*credentialsRaw, &anonymousCredential); err != nil {
			return err
		}
		r.Credentials = anonymousCredential
	case credentials.GitCredentialTypeUsernamePassword:
		var usernamePasswordCredential *credentials.UsernamePassword
		if err := json.Unmarshal(*credentialsRaw, &usernamePasswordCredential); err != nil {
			return err
		}
		r.Credentials = usernamePasswordCredential
	}

	return nil
}
