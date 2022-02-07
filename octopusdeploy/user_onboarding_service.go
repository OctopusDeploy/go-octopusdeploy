package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type userOnboardingService struct {
	services.service
}

func newUserOnboardingService(sling *sling.Sling, uriTemplate string) *userOnboardingService {
	return &userOnboardingService{
		service: services.newService(ServiceUserOnboardingService, sling, uriTemplate),
	}
}
