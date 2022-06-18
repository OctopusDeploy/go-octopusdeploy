package licenses

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type LicenseService struct {
	currentLicense       string
	currentLicenseStatus string

	services.Service
}

func NewLicenseService(sling *sling.Sling, uriTemplate string, currentLicense string, currentLicenseStatus string) *LicenseService {
	return &LicenseService{
		currentLicense:       currentLicense,
		currentLicenseStatus: currentLicenseStatus,
		Service:              services.NewService(constants.ServiceLicenseService, sling, uriTemplate),
	}
}
