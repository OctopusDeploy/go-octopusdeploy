package certificates

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type CertificateResource struct {
	Archived                 string                      `json:"Archived,omitempty"`
	CertificateData          *core.SensitiveValue        `json:"CertificateData,omitempty" validate:"required"`
	CertificateDataFormat    string                      `json:"CertificateDataFormat,omitempty"`
	EnvironmentIDs           []string                    `json:"EnvironmentIds,omitempty"`
	HasPrivateKey            bool                        `json:"HasPrivateKey"`
	IsExpired                bool                        `json:"IsExpired"`
	IssuerCommonName         string                      `json:"IssuerCommonName,omitempty"`
	IssuerDistinguishedName  string                      `json:"IssuerDistinguishedName,omitempty"`
	IssuerOrganization       string                      `json:"IssuerOrganization,omitempty"`
	Name                     string                      `json:"Name,omitempty" validate:"required"`
	NotAfter                 string                      `json:"NotAfter,omitempty"`
	NotBefore                string                      `json:"NotBefore,omitempty"`
	Notes                    string                      `json:"Notes,omitempty"`
	Password                 *core.SensitiveValue        `json:"Password,omitempty"`
	ReplacedBy               string                      `json:"ReplacedBy,omitempty"`
	SelfSigned               bool                        `json:"SelfSigned"`
	SerialNumber             string                      `json:"SerialNumber,omitempty"`
	SignatureAlgorithmName   string                      `json:"SignatureAlgorithmName,omitempty"`
	SubjectAlternativeNames  []string                    `json:"SubjectAlternativeNames,omitempty"`
	SubjectCommonName        string                      `json:"SubjectCommonName,omitempty"`
	SubjectDistinguishedName string                      `json:"SubjectDistinguishedName,omitempty"`
	SubjectOrganization      string                      `json:"SubjectOrganization,omitempty"`
	TenantedDeploymentMode   core.TenantedDeploymentMode `json:"TenantedDeploymentParticipation"`
	TenantIDs                []string                    `json:"TenantIds,omitempty"`
	TenantTags               []string                    `json:"TenantTags,omitempty"`
	Thumbprint               string                      `json:"Thumbprint,omitempty"`
	Version                  int                         `json:"Version,omitempty"`

	resources.Resource
}

// NewCertificateResource initializes a certificate resource with a name and
// credentials.
func NewCertificateResource(name string, certificateData *core.SensitiveValue, password *core.SensitiveValue) *CertificateResource {
	return &CertificateResource{
		Name:                   name,
		CertificateData:        certificateData,
		Password:               password,
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
		Resource:               *resources.NewResource(),
	}
}

// Validate checks the state of the certificate resource and returns an error
// if invalid.
func (c CertificateResource) Validate() error {
	return validator.New().Struct(c)
}
