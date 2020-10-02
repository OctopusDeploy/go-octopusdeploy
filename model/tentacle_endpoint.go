package model

type tentacleEndpoint struct {
	CertificateSignatureAlgorithm *string                `json:"CertificateSignatureAlgorithm"`
	Thumbprint                    string                 `json:"Thumbprint" validate:"required"`
	TentacleVersionDetails        TentacleVersionDetails `json:"TentacleVersionDetails"`

	endpoint
}
