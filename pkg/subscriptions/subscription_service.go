package subscriptions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type SubscriptionService struct {
	services.CanDeleteService
}

func NewSubscriptionService(sling *sling.Sling, uriTemplate string) *SubscriptionService {
	return &SubscriptionService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceSubscriptionService, sling, uriTemplate),
		},
	}
}
