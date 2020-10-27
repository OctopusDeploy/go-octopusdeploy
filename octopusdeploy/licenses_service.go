package octopusdeploy

import "github.com/dghubble/sling"

type licenseService struct {
	currentLicense       string
	currentLicenseStatus string

	service
}

func newLicenseService(sling *sling.Sling, uriTemplate string, currentLicense string, currentLicenseStatus string) *licenseService {
	return &licenseService{
		currentLicense:       currentLicense,
		currentLicenseStatus: currentLicenseStatus,
		service:              newService(ServiceLicenseService, sling, uriTemplate),
	}
}
