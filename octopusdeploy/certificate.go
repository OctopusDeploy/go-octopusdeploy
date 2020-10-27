package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type Certificates struct {
	Items []*Certificate `json:"Items"`
	PagedResults
}

type Certificate struct {
	Archived                 string                 `json:"Archived,omitempty"`
	CertificateData          *SensitiveValue        `json:"CertificateData,omitempty" validate:"required"`
	CertificateDataFormat    string                 `json:"CertificateDataFormat,omitempty"`
	EnvironmentIDs           []string               `json:"EnvironmentIds,omitempty"`
	HasPrivateKey            bool                   `json:"HasPrivateKey,omitempty"`
	IsExpired                bool                   `json:"IsExpired,omitempty"`
	IssuerCommonName         string                 `json:"IssuerCommonName,omitempty"`
	IssuerDistinguishedName  string                 `json:"IssuerDistinguishedName,omitempty"`
	IssuerOrganization       string                 `json:"IssuerOrganization,omitempty"`
	Name                     string                 `json:"Name,omitempty" validate:"required"`
	NotAfter                 string                 `json:"NotAfter,omitempty"`
	NotBefore                string                 `json:"NotBefore,omitempty"`
	Notes                    string                 `json:"Notes,omitempty"`
	Password                 *SensitiveValue        `json:"Password,omitempty"`
	ReplacedBy               string                 `json:"ReplacedBy,omitempty"`
	SerialNumber             string                 `json:"SerialNumber,omitempty"`
	SignatureAlgorithmName   string                 `json:"SignatureAlgorithmName,omitempty"`
	SubjectAlternativeNames  []string               `json:"SubjectAlternativeNames,omitempty"`
	SubjectDistinguishedName string                 `json:"SubjectDistinguishedName,omitempty"`
	SubjectCommonName        string                 `json:"SubjectCommonName,omitempty"`
	SubjectOrganization      string                 `json:"SubjectOrganization,omitempty"`
	SelfSigned               bool                   `json:"SelfSigned,omitempty"`
	TenantedDeploymentMode   TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantIDs                []string               `json:"TenantIds,omitempty"`
	TenantTags               []string               `json:"TenantTags,omitempty"`
	Thumbprint               string                 `json:"Thumbprint,omitempty"`
	Version                  int                    `json:"Version,omitempty"`

	resource
}

// NewCertificate initializes a Certificate with a name and credentials.
func NewCertificate(name string, certificateData *SensitiveValue, password *SensitiveValue) *Certificate {
	return &Certificate{
		Name:                   name,
		CertificateData:        certificateData,
		Password:               password,
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
		resource:               *newResource(),
	}
}

// Validate checks the state of the certificate and returns an error if invalid.
func (c Certificate) Validate() error {
	return validator.New().Struct(c)
}
