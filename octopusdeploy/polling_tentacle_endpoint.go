package octopusdeploy

import (
	"encoding/json"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type PollingTentacleEndpoint struct {
	tentacleEndpoint `validate:"required"`
}

func NewPollingTentacleEndpoint(uri *url.URL, thumbprint string) *PollingTentacleEndpoint {
	return &PollingTentacleEndpoint{
		tentacleEndpoint: *newTentacleEndpoint("TentacleActive", thumbprint, uri),
	}
}

// Validate checks the state of the listening tentacle endpoint and returns an
// error if invalid.
func (l *PollingTentacleEndpoint) Validate() error {
	return validator.New().Struct(l)
}

func (p PollingTentacleEndpoint) MarshalJSON() ([]byte, error) {
	te := struct {
		CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
		TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
		Thumbprint                    string                  `json:"Thumbprint" validate:"required"`
		URI                           string                  `json:"Uri" validate:"required,uri"`
		endpoint
	}{
		CertificateSignatureAlgorithm: p.CertificateSignatureAlgorithm,
		TentacleVersionDetails:        p.TentacleVersionDetails,
		Thumbprint:                    p.Thumbprint,
		endpoint:                      p.endpoint,
	}

	if p.URI != nil {
		te.URI = p.URI.String()
	}

	return json.Marshal(te)
}

// UnmarshalJSON sets this tentacle endpoint to its representation in JSON.
func (p *PollingTentacleEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
		TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
		Thumbprint                    string                  `json:"Thumbprint" validate:"required"`
		URI                           string                  `json:"Uri" validate:"required,uri"`
		endpoint
	}

	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	if err := validate.Struct(fields); err != nil {
		return err
	}

	u, err := url.Parse(fields.URI)
	if err != nil {
		return err
	}

	p.CertificateSignatureAlgorithm = fields.CertificateSignatureAlgorithm
	p.TentacleVersionDetails = fields.TentacleVersionDetails
	p.Thumbprint = fields.Thumbprint
	p.URI = u
	p.endpoint = fields.endpoint

	return nil
}
