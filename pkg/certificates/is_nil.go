package certificates

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *CertificateResource:
		return v == nil
	default:
		return v == nil
	}
}
