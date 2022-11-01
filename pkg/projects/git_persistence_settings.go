package projects

import (
	"encoding/json"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/go-playground/validator/v10"
)

type GitPersistenceSettings interface {
	BasePath() string
	SetBasePath(basePath string)

	DefaultBranch() string
	SetDefaultBranch(defaultBranch string)

	ProtectedBranchNamePatterns() []string
	SetProtectedBranchNamePatterns(protectedBranchNamePatterns []string)

	URL() *url.URL
	SetURL(url *url.URL)

	Credential() credentials.GitCredential
	SetCredential(credential credentials.GitCredential)

	PersistenceSettings
}

// GitPersistenceSettings represents persistence settings associated with a project.
type gitPersistenceSettings struct {
	basePath                    string
	conversionState             *ConversionState
	credential                  credentials.GitCredential
	defaultBranch               string
	protectedBranchNamePatterns []string
	url                         *url.URL

	persistenceSettings
}

// NewGitPersistenceSettings creates an instance of persistence settings.
func NewGitPersistenceSettings(
	basePath string,
	credentials credentials.GitCredential,
	defaultBranch string,
	protectedBranchNamePatterns []string,
	url *url.URL) GitPersistenceSettings {
	return &gitPersistenceSettings{
		basePath:                    basePath,
		credential:                  credentials,
		defaultBranch:               defaultBranch,
		protectedBranchNamePatterns: protectedBranchNamePatterns,
		url:                         url,
		persistenceSettings:         persistenceSettings{SettingsType: PersistenceSettingsTypeVersionControlled},
	}
}

// Type returns the type for this persistence settings.
func (g *gitPersistenceSettings) Type() PersistenceSettingsType {
	return g.SettingsType
}

func (g *gitPersistenceSettings) BasePath() string {
	return g.basePath
}

func (g *gitPersistenceSettings) SetBasePath(basePath string) {
	g.basePath = basePath
}

func (g *gitPersistenceSettings) DefaultBranch() string {
	return g.defaultBranch
}

func (g *gitPersistenceSettings) SetDefaultBranch(defaultBranch string) {
	g.defaultBranch = defaultBranch
}

func (g *gitPersistenceSettings) ProtectedBranchNamePatterns() []string {
	return g.protectedBranchNamePatterns
}

func (g *gitPersistenceSettings) SetProtectedBranchNamePatterns(protectedBranchNamePatterns []string) {
	g.protectedBranchNamePatterns = protectedBranchNamePatterns
}

func (g *gitPersistenceSettings) URL() *url.URL {
	return g.url
}

func (g *gitPersistenceSettings) SetURL(url *url.URL) {
	g.url = url
}

func (g *gitPersistenceSettings) Credential() credentials.GitCredential {
	return g.credential
}

func (g *gitPersistenceSettings) SetCredential(credential credentials.GitCredential) {
	g.credential = credential
}

// MarshalJSON returns persistence settings as its JSON encoding.
func (p *gitPersistenceSettings) MarshalJSON() ([]byte, error) {
	persistenceSettings := struct {
		BasePath                    string                    `json:"BasePath,omitempty"`
		ConversionState             *ConversionState          `json:"ConversionState,omitempty"`
		Credentials                 credentials.GitCredential `json:"Credentials,omitempty"`
		DefaultBranch               string                    `json:"DefaultBranch,omitempty"`
		ProtectedBranchNamePatterns []string                  `json:"ProtectedBranchNamePatterns"`
		URL                         string                    `json:"Url,omitempty"`
		Type                        PersistenceSettingsType   `json:"Type,omitempty"`
	}{
		BasePath:                    p.basePath,
		ConversionState:             p.conversionState,
		Credentials:                 p.credential,
		DefaultBranch:               p.defaultBranch,
		ProtectedBranchNamePatterns: p.protectedBranchNamePatterns,
		URL:                         p.url.String(),
		Type:                        p.persistenceSettings.SettingsType,
	}

	return json.Marshal(persistenceSettings)
}

// UnmarshalJSON sets the persistence settings to its representation in JSON.
func (p *gitPersistenceSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		BasePath                    string                  `json:"BasePath,omitempty"`
		ConversionState             *ConversionState        `json:"ConversionState,omitempty"`
		DefaultBranch               string                  `json:"DefaultBranch,omitempty"`
		ProtectedBranchNamePatterns []string                `json:"ProtectedBranchNamePatterns"`
		URL                         string                  `json:"Url,omitempty"`
		Type                        PersistenceSettingsType `json:"Type"`
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

	p.basePath = fields.BasePath
	p.conversionState = fields.ConversionState
	p.defaultBranch = fields.DefaultBranch
	p.protectedBranchNamePatterns = fields.ProtectedBranchNamePatterns
	p.SettingsType = fields.Type
	p.url = url

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
		p.credential = anonymousGitCredential
	case credentials.GitCredentialTypeReference:
		var referenceProjectGitCredential *credentials.Reference
		err := json.Unmarshal(*gitCredentials, &referenceProjectGitCredential)
		if err != nil {
			return err
		}
		p.credential = referenceProjectGitCredential
	case credentials.GitCredentialTypeUsernamePassword:
		var usernamePasswordGitCredential *credentials.UsernamePassword
		err := json.Unmarshal(*gitCredentials, &usernamePasswordGitCredential)
		if err != nil {
			return err
		}
		p.credential = usernamePasswordGitCredential
	}

	return nil
}

var _ GitPersistenceSettings = &gitPersistenceSettings{}
