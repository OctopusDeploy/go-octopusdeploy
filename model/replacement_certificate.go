package model

func NewReplacementCertificate(certificateData string, password string) (*ReplacementCertificate, error) {
	if isEmpty(certificateData) {
		return nil, createInvalidParameterError("NewReplacementCertificate", "certificateData")
	}

	if isEmpty(password) {
		return nil, createInvalidParameterError("NewReplacementCertificate", "password")
	}

	return &ReplacementCertificate{
		CertificateData: certificateData,
		Password:        password,
	}, nil
}

type ReplacementCertificate struct {
	CertificateData string `json:"CertificateData,omitempty"`
	Password        string `json:"Password,omitempty"`
}
