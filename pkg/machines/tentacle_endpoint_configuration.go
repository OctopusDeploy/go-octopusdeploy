package machines

import (
	"encoding/json"
	"net/url"
)

type TentacleEndpointConfiguration struct {
	CertificateSignatureAlgorithm string   `json:"CertificateSignatureAlgorithm,omitempty"`
	Thumbprint                    string   `json:"Thumbprint" validate:"required"`
	URI                           *url.URL `json:"Uri" validate:"required"`
	CommunicationMode             string   `json:"CommunicationMode" validate:"required,oneof=Polling Listening"`
}

func (l TentacleEndpointConfiguration) MarshalJSON() ([]byte, error) {
	te := struct {
		CertificateSignatureAlgorithm string `json:"CertificateSignatureAlgorithm,omitempty"`
		Thumbprint                    string `json:"Thumbprint" validate:"required"`
		URI                           string `json:"Uri" validate:"required,uri"`
		CommunicationMode             string `json:"CommunicationMode" validate:"required,oneof=Polling Listening"`
	}{
		CertificateSignatureAlgorithm: l.CertificateSignatureAlgorithm,
		Thumbprint:                    l.Thumbprint,
		CommunicationMode:             l.CommunicationMode,
	}

	if l.URI != nil {
		te.URI = l.URI.String()
	}

	return json.Marshal(te)
}

func (l *TentacleEndpointConfiguration) UnmarshalJSON(b []byte) error {
	var fields struct {
		CertificateSignatureAlgorithm string `json:"CertificateSignatureAlgorithm,omitempty"`
		Thumbprint                    string `json:"Thumbprint" validate:"required"`
		URI                           string `json:"Uri" validate:"required,uri"`
		CommunicationMode             string `json:"CommunicationMode" validate:"required,oneof=Polling Listening"`
	}

	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	uri, err := url.Parse(fields.URI)
	if err != nil {
		return err
	}

	l.CertificateSignatureAlgorithm = fields.CertificateSignatureAlgorithm
	l.Thumbprint = fields.Thumbprint
	l.CommunicationMode = fields.CommunicationMode
	l.URI = uri

	return nil
}
