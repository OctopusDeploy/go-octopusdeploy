package resources

type ReplacementCertificate struct {
	CertificateData string `json:"CertificateData,omitempty"`
	Password        string `json:"Password,omitempty"`
}

func NewReplacementCertificate(certificateData string, password string) *ReplacementCertificate {
	return &ReplacementCertificate{
		CertificateData: certificateData,
		Password:        password,
	}
}
