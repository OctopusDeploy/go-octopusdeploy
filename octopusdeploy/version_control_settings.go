package octopusdeploy

type VersionControlSettings struct {
	BasePath      string          `json:"BasePath,omitempty"`
	DefaultBranch string          `json:"DefaultBranch,omitempty"`
	HasValue      bool            `json:"HasValue,omitempty"`
	Password      *SensitiveValue `json:"Password,omitempty"`
	URL           string          `json:"Url,omitempty"`
	Username      string          `json:"Username,omitempty"`
}
