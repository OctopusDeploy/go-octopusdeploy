package machines

import (
	"net/url"

	"github.com/go-playground/validator/v10"
)

type tentacleEndpoint struct {
	CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
	TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
	Thumbprint                    string                  `json:"Thumbprint" validate:"required"`
	URI                           *url.URL                `json:"Uri" validate:"required,uri"`

	endpoint
}

func newTentacleEndpoint(communicationStyle string, thumbprint string, uri *url.URL) *tentacleEndpoint {
	return &tentacleEndpoint{
		Thumbprint: thumbprint,
		URI:        uri,
		endpoint:   *newEndpoint(communicationStyle),
	}
}

// Validate checks the state of the tentacle endpoint and returns an error if
// invalid.
func (t tentacleEndpoint) Validate() error {
	return validator.New().Struct(t)
}
