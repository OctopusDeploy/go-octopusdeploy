package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type Certificates struct {
	Items []Certificate `json:"Items"`
	PagedResults
}

type Certificate struct {
	Archived                        string                      `json:"Archived,omitempty"`
	CertificateData                 *SensitiveValue             `json:"CertificateData,omitempty" validate:"required"`
	CertificateDataFormat           string                      `json:"CertificateDataFormat,omitempty"`
	EnvironmentIDs                  []string                    `json:"EnvironmentIds,omitempty"`
	HasPrivateKey                   bool                        `json:"HasPrivateKey,omitempty"`
	IsExpired                       bool                        `json:"IsExpired,omitempty"`
	IssuerCommonName                string                      `json:"IssuerCommonName,omitempty"`
	IssuerDistinguishedName         string                      `json:"IssuerDistinguishedName,omitempty"`
	IssuerOrganization              string                      `json:"IssuerOrganization,omitempty"`
	Name                            string                      `json:"Name,omitempty" validate:"required"`
	NotAfter                        string                      `json:"NotAfter,omitempty"`
	NotBefore                       string                      `json:"NotBefore,omitempty"`
	Notes                           string                      `json:"Notes,omitempty"`
	Password                        *SensitiveValue             `json:"Password,omitempty"`
	ReplacedBy                      string                      `json:"ReplacedBy,omitempty"`
	SerialNumber                    string                      `json:"SerialNumber,omitempty"`
	SignatureAlgorithmName          string                      `json:"SignatureAlgorithmName,omitempty"`
	SubjectAlternativeNames         []string                    `json:"SubjectAlternativeNames,omitempty"`
	SubjectDistinguishedName        string                      `json:"SubjectDistinguishedName,omitempty"`
	SubjectCommonName               string                      `json:"SubjectCommonName,omitempty"`
	SubjectOrganization             string                      `json:"SubjectOrganization,omitempty"`
	SelfSigned                      bool                        `json:"SelfSigned,omitempty"`
	TenantedDeploymentParticipation enum.TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIDs                       []string                    `json:"TenantIds,omitempty"`
	TenantTags                      []string                    `json:"TenantTags,omitempty"`
	Thumbprint                      string                      `json:"Thumbprint,omitempty"`
	Version                         int                         `json:"Version,omitempty"`

	Resource
}

// NewCertificate initializes a Certificate with a name and credentials.
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

// GetID returns the ID value of the Certificate.
func (resource Certificate) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Certificate.
func (resource Certificate) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Certificate was changed.
func (resource Certificate) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Certificate.
func (resource Certificate) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Certificate and returns an error if invalid.
func (resource Certificate) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &Certificate{}
