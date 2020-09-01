package model

type VersionControlSettings struct {
	DefaultBranch string          `json:"DefaultBranch,omitempty"`
	Password      *SensitiveValue `json:"Password,omitempty"`
	URL           string          `json:"Url,omitempty"`
	Username      string          `json:"Username,omitempty"`
}
