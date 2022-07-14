package machines

import (
	"encoding/json"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type ListeningTentacleEndpoint struct {
	ProxyID string `json:"ProxyId,omitempty"`

	tentacleEndpoint
}

func NewListeningTentacleEndpoint(uri *url.URL, thumbprint string) *ListeningTentacleEndpoint {
	return &ListeningTentacleEndpoint{
		tentacleEndpoint: *newTentacleEndpoint("TentaclePassive", thumbprint, uri),
	}
}

func (l ListeningTentacleEndpoint) MarshalJSON() ([]byte, error) {
	te := struct {
		CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
		ProxyID                       string                  `json:"ProxyId,omitempty"`
		TentacleVersionDetails        *TentacleVersionDetails `json:"TentacleVersionDetails,omitempty"`
		Thumbprint                    string                  `json:"Thumbprint" validate:"required"`
		URI                           string                  `json:"Uri" validate:"required,uri"`
		endpoint
	}{
		CertificateSignatureAlgorithm: l.CertificateSignatureAlgorithm,
		ProxyID:                       l.ProxyID,
		TentacleVersionDetails:        l.TentacleVersionDetails,
		Thumbprint:                    l.Thumbprint,
		endpoint:                      l.endpoint,
	}

	if l.URI != nil {
		te.URI = l.URI.String()
	}

	return json.Marshal(te)
}

// UnmarshalJSON sets this tentacle endpoint to its representation in JSON.
func (l *ListeningTentacleEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		CertificateSignatureAlgorithm string                  `json:"CertificateSignatureAlgorithm,omitempty"`
		ProxyID                       string                  `json:"ProxyId,omitempty"`
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

	l.CertificateSignatureAlgorithm = fields.CertificateSignatureAlgorithm
	l.ProxyID = fields.ProxyID
	l.TentacleVersionDetails = fields.TentacleVersionDetails
	l.Thumbprint = fields.Thumbprint
	l.URI = u
	l.endpoint = fields.endpoint

	return nil
}
