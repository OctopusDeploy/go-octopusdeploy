package projects

import (
	"encoding/json"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/go-playground/validator/v10"
)

type GitPersistenceSettings interface {
	GetBasePath() string
	GetDefaultBranch() string
	GetProtectedBranchNamePatterns() []string
	GetURL() *url.URL
	GetCredential() credentials.IGitCredential
	PersistenceSettings
}

// GitPersistenceSettings represents persistence settings associated with a project.
type gitPersistenceSettings struct {
	BasePath                    string
	ConversionState             *ConversionState
	Credentials                 credentials.IGitCredential
	DefaultBranch               string
	ProtectedBranchNamePatterns []string
	URL                         *url.URL

	persistenceSettings
}

// NewGitPersistenceSettings creates an instance of persistence settings.
func NewGitPersistenceSettings(
	basePath string,
	credentials credentials.IGitCredential,
	defaultBranch string,
	protectedBranchNamePatterns []string,
	url *url.URL) GitPersistenceSettings {
	return &gitPersistenceSettings{
		BasePath:                    basePath,
		Credentials:                 credentials,
		DefaultBranch:               defaultBranch,
		ProtectedBranchNamePatterns: protectedBranchNamePatterns,
		URL:                         url,
		persistenceSettings:         persistenceSettings{Type: PersistenceSettingsTypeVersionControlled},
	}
}

// GetType returns the type for this persistence settings.
func (g gitPersistenceSettings) GetType() PersistenceSettingsType {
	return g.Type
}

func (g gitPersistenceSettings) GetBasePath() string {
	return g.BasePath
}

func (g gitPersistenceSettings) GetDefaultBranch() string {
	return g.DefaultBranch
}

func (g gitPersistenceSettings) GetProtectedBranchNamePatterns() []string {
	return g.ProtectedBranchNamePatterns
}

func (g gitPersistenceSettings) GetURL() *url.URL {
	return g.URL
}

func (g gitPersistenceSettings) GetCredential() credentials.IGitCredential {
	return g.Credentials
}

// MarshalJSON returns persistence settings as its JSON encoding.
func (p gitPersistenceSettings) MarshalJSON() ([]byte, error) {
	persistenceSettings := struct {
		BasePath                    string                     `json:"BasePath,omitempty"`
		ConversionState             *ConversionState           `json:"ConversionState,omitempty"`
		Credentials                 credentials.IGitCredential `json:"Credentials,omitempty"`
		DefaultBranch               string                     `json:"DefaultBranch,omitempty"`
		ProtectedBranchNamePatterns []string                   `json:"ProtectedBranchNamePatterns"`
		URL                         string                     `json:"Url,omitempty"`
		persistenceSettings
	}{
		BasePath:                    p.BasePath,
		ConversionState:             p.ConversionState,
		Credentials:                 p.Credentials,
		DefaultBranch:               p.DefaultBranch,
		ProtectedBranchNamePatterns: p.ProtectedBranchNamePatterns,
		URL:                         p.URL.String(),
		persistenceSettings:         p.persistenceSettings,
	}

	return json.Marshal(persistenceSettings)
}

// UnmarshalJSON sets the persistence settings to its representation in JSON.
func (p gitPersistenceSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		BasePath                    string           `json:"BasePath,omitempty"`
		ConversionState             *ConversionState `json:"ConversionState,omitempty"`
		DefaultBranch               string           `json:"DefaultBranch,omitempty"`
		ProtectedBranchNamePatterns []string         `json:"ProtectedBranchNamePatterns"`
		URL                         string           `json:"Url,omitempty"`
		persistenceSettings
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
	p.ConversionState = fields.ConversionState
	p.DefaultBranch = fields.DefaultBranch
	p.ProtectedBranchNamePatterns = fields.ProtectedBranchNamePatterns
	p.Type = fields.Type
	p.URL = url
	p.persistenceSettings = fields.persistenceSettings

	var persistenceSettings map[string]*json.RawMessage
	err = json.Unmarshal(b, &persistenceSettings)
	if err != nil {
		return err
	}

	var gitCredentials *json.RawMessage
	var credentialsProperties map[string]*json.RawMessage
	var gitCredentialType credentials.Type

	if persistenceSettings["Credentials"] != nil {
		credentialsValue := persistenceSettings["Credentials"]

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
	}

	switch gitCredentialType {
	case credentials.GitCredentialTypeAnonymous:
		var anonymousGitCredential *credentials.Anonymous
		err := json.Unmarshal(*gitCredentials, &anonymousGitCredential)
		if err != nil {
			return err
		}
		p.Credentials = anonymousGitCredential
	case credentials.GitCredentialTypeReference:
		var referenceProjectGitCredential *credentials.Reference
		err := json.Unmarshal(*gitCredentials, &referenceProjectGitCredential)
		if err != nil {
			return err
		}
		p.Credentials = referenceProjectGitCredential
	case credentials.GitCredentialTypeUsernamePassword:
		var usernamePasswordGitCredential *credentials.UsernamePassword
		err := json.Unmarshal(*gitCredentials, &usernamePasswordGitCredential)
		if err != nil {
			return err
		}
		p.Credentials = usernamePasswordGitCredential
	}

	return nil
}

var _ GitPersistenceSettings = &gitPersistenceSettings{}
