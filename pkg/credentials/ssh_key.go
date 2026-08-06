package credentials

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

type SshKey struct {
	Username              string               `json:"Username,omitempty"`
	PrivateKey            *core.SensitiveValue `json:"PrivateKey,omitempty"`
	Passphrase            *core.SensitiveValue `json:"Passphrase,omitempty"`
	KeyName               string               `json:"KeyName,omitempty"`
	PrivateKeyFingerprint string               `json:"PrivateKeyFingerprint,omitempty"`

	gitCredential
}

func NewSshKey(privateKey *core.SensitiveValue) *SshKey {
	return &SshKey{
		PrivateKey: privateKey,
		gitCredential: gitCredential{
			CredentialType: GitCredentialTypeSshKey,
		},
	}
}

func (s *SshKey) Type() Type {
	return s.CredentialType
}

var _ GitCredential = &SshKey{}
