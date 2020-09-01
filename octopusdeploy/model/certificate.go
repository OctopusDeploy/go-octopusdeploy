package model

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type Certificates struct {
	Items []Certificate `json:"Items"`
	PagedResults
}

type CertificateReplace struct {
	CertificateData string `json:"CertificateData,omitempty"`
	Password        string `json:"Password,omitempty"`
}

type Certificate struct {
	ID                              string                      `json:"Id,omitempty"`
	Name                            string                      `json:"Name,omitempty"`
	Notes                           string                      `json:"Notes,omitempty"`
	CertificateData                 SensitiveValue              `json:"CertificateData,omitempty"`
	Password                        SensitiveValue              `json:"Password,omitempty"`
	EnvironmentIds                  []string                    `json:"EnvironmentIds,omitempty"`
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
}

func (t *Certificate) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewCertificate(name string, certificateData SensitiveValue, password SensitiveValue) *Certificate {
	return &Certificate{
		Name:            name,
		CertificateData: certificateData,
		Password:        password,
	}
}

func NewCertificateReplace(certificateData string, password string) *CertificateReplace {
	return &CertificateReplace{
		CertificateData: certificateData,
		Password:        password,
	}
}
