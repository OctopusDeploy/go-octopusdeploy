package octopusdeploy

import "github.com/go-playground/validator/v10"

type tentacleEndpoint struct {
	CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
	TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
	Thumbprint                    string                  `json:"Thumbprint" validate:"required"`

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
