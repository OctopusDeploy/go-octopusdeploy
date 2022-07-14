package useronboarding

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type UserOnboardingService struct {
	services.Service
}

func NewUserOnboardingService(sling *sling.Sling, uriTemplate string) *UserOnboardingService {
	return &UserOnboardingService{
		Service: services.NewService(constants.ServiceUserOnboardingService, sling, uriTemplate),
	}
}
