package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type licenseService struct {
	currentLicense       string
	currentLicenseStatus string

	services.service
}

func newLicenseService(sling *sling.Sling, uriTemplate string, currentLicense string, currentLicenseStatus string) *licenseService {
	return &licenseService{
		currentLicense:       currentLicense,
		currentLicenseStatus: currentLicenseStatus,
		service:              services.newService(ServiceLicenseService, sling, uriTemplate),
	}
}
