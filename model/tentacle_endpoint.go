package model

import "github.com/go-playground/validator/v10"

type tentacleEndpoint struct {
	CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
	Thumbprint                    string                  `json:"Thumbprint" validate:"required"`
	TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`

	endpoint
}

func newTentacleEndpoint(communicationStyle string, thumbprint string) *tentacleEndpoint {
	return &tentacleEndpoint{
		Thumbprint: thumbprint,
		endpoint:   *newEndpoint(communicationStyle),
	}
}

// Validate checks the state of the tentacle endpoint and returns an error if
// invalid.
func (t tentacleEndpoint) Validate() error {
	return validator.New().Struct(t)
}