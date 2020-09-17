package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type Certificates struct {
	Items []Certificate `json:"Items"`
	PagedResults
}

type Certificate struct {
	Name                            string                      `json:"Name,omitempty" validate:"required"`
	Notes                           string                      `json:"Notes,omitempty"`
	CertificateData                 *SensitiveValue             `json:"CertificateData,omitempty" validate:"required"`
	Password                        *SensitiveValue             `json:"Password,omitempty"`
	EnvironmentIDs                  []string                    `json:"EnvironmentIds,omitempty"`
	TenantedDeploymentParticipation enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIds                       []string                    `json:"TenantIds,omitempty"`
	TenantTags                      []string                    `json:"TenantTags,omitempty"`
	CertificateDataFormat           string                      `json:"CertificateDataFormat,omitempty"`
	Archived                        string                      `json:"Archived,omitempty"`
	ReplacedBy                      string                      `json:"ReplacedBy,omitempty"`
	SubjectDistinguishedName        string                      `json:"SubjectDistinguishedName,omitempty"`
	SubjectCommonName               string                      `json:"SubjectCommonName,omitempty"`
	SubjectOrganization             string                      `json:"SubjectOrganization,omitempty"`
	IssuerDistinguishedName         string                      `json:"IssuerDistinguishedName,omitempty"`
	IssuerCommonName                string                      `json:"IssuerCommonName,omitempty"`
	IssuerOrganization              string                      `json:"IssuerOrganization,omitempty"`
	SelfSigned                      bool                        `json:"SelfSigned,omitempty"`
	Thumbprint                      string                      `json:"Thumbprint,omitempty"`
	NotAfter                        string                      `json:"NotAfter,omitempty"`
	NotBefore                       string                      `json:"NotBefore,omitempty"`
	IsExpired                       bool                        `json:"IsExpired,omitempty"`
	HasPrivateKey                   bool                        `json:"HasPrivateKey,omitempty"`
	Version                         int                         `json:"Version,omitempty"`
	SerialNumber                    string                      `json:"SerialNumber,omitempty"`
	SignatureAlgorithmName          string                      `json:"SignatureAlgorithmName,omitempty"`
	SubjectAlternativeNames         []string                    `json:"SubjectAlternativeNames,omitempty"`

	Resource
}

func (c *Certificate) GetID() string {
	return c.ID
}

func (c *Certificate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func NewCertificate(name string, certificateData SensitiveValue, password SensitiveValue) (*Certificate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewCertificate", "name")
	}

	return &Certificate{
		Name:            name,
		CertificateData: &certificateData,
		Password:        &password,
	}, nil
}

var _ ResourceInterface = &Certificate{}
