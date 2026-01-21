package projects

import (
	"encoding/json"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
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

	VariablesAreInGit() bool
	RunbooksAreInGit() bool

	// Deprecated: This is not settable against a real Octopus project it is only used for testing purposes.
	SetRunbooksAreInGit()

	PersistenceSettings
}

// GitPersistenceSettings represents persistence settings associated with a project.
type gitPersistenceSettings struct {
	basePath                    string
	credential                  credentials.GitCredential
	defaultBranch               string
	protectedBranchNamePatterns []string
	url                         *url.URL
	conversionState             gitPersistenceSettingsConversionState

	persistenceSettings
}

type gitPersistenceSettingsConversionState struct {
	VariablesAreInGit bool `json:"VariablesAreInGit,omitempty"`
	RunbooksAreInGit  bool `json:"RunbooksAreInGit,omitempty"`
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
		conversionState:             gitPersistenceSettingsConversionState{VariablesAreInGit: false, RunbooksAreInGit: false},
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

func (g *gitPersistenceSettings) VariablesAreInGit() bool {
	return g.conversionState.VariablesAreInGit
}

func (g *gitPersistenceSettings) RunbooksAreInGit() bool {
	return g.conversionState.RunbooksAreInGit
}

func (g *gitPersistenceSettings) SetRunbooksAreInGit() {
	g.conversionState.RunbooksAreInGit = true
}

// MarshalJSON returns persistence settings as its JSON encoding.
func (p *gitPersistenceSettings) MarshalJSON() ([]byte, error) {
	defaultBranch := p.DefaultBranch()
	protectedBranches := p.ProtectedBranchNamePatterns()
	isDefaultBranchProtected := slices.Contains(protectedBranches, defaultBranch)

	if isDefaultBranchProtected {
		i := slices.Index(protectedBranches, defaultBranch)
		protectedBranches = slices.Delete(protectedBranches, i, i+1)
	}

	persistenceSettings := struct {
		BasePath                    string                    `json:"BasePath,omitempty"`
		Credentials                 credentials.GitCredential `json:"Credentials,omitempty"`
		DefaultBranch               string                    `json:"DefaultBranch,omitempty"`
		IsDefaultBranchProtected    bool                      `json:"ProtectedDefaultBranch"`
		ProtectedBranchNamePatterns []string                  `json:"ProtectedBranchNamePatterns"`
		URL                         string                    `json:"Url,omitempty"`
		Type                        PersistenceSettingsType   `json:"Type,omitempty"`
	}{
		BasePath:                    p.BasePath(),
		Credentials:                 p.Credential(),
		DefaultBranch:               p.DefaultBranch(),
		IsDefaultBranchProtected:    isDefaultBranchProtected,
		ProtectedBranchNamePatterns: protectedBranches,
		URL:                         p.URL().String(),
		Type:                        p.Type(),
	}

	return json.Marshal(persistenceSettings)
}

// UnmarshalJSON sets the persistence settings to its representation in JSON.
func (p *gitPersistenceSettings) UnmarshalJSON(b []byte) error {
	var fields struct {
		BasePath                    string                  `json:"BasePath,omitempty"`
		DefaultBranch               string                  `json:"DefaultBranch,omitempty"`
		IsDefaultBranchProtected    bool                    `json:"ProtectedDefaultBranch"`
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

	isDefaultBranchProtected := fields.IsDefaultBranchProtected

	p.basePath = fields.BasePath
	p.defaultBranch = fields.DefaultBranch
	p.protectedBranchNamePatterns = fields.ProtectedBranchNamePatterns

	if isDefaultBranchProtected {
		p.protectedBranchNamePatterns = append(p.protectedBranchNamePatterns, p.defaultBranch)
	}

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
	case credentials.GitCredentialTypeGitHub:
		var githubGitCredential *credentials.GitHub
		err := json.Unmarshal(*gitCredentials, &githubGitCredential)
		if err != nil {
			return err
		}
		p.credential = githubGitCredential
	}

	var conversionState *json.RawMessage
	var conversionStateFields gitPersistenceSettingsConversionState

	if persistenceSettings["ConversionState"] != nil {
		conversionStateValue := persistenceSettings["ConversionState"]

		err = json.Unmarshal(*conversionStateValue, &conversionState)
		if err != nil {
			return err
		}

		err = json.Unmarshal(*conversionState, &conversionStateFields)
		if err != nil {
			return err
		}

		p.conversionState = conversionStateFields
	}

	return nil
}

var _ GitPersistenceSettings = &gitPersistenceSettings{}
