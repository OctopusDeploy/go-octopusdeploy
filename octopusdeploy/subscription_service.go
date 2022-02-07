package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type subscriptionService struct {
	services.canDeleteService
}

func newSubscriptionService(sling *sling.Sling, uriTemplate string) *subscriptionService {
	subscriptionService := &subscriptionService{}
	subscriptionService.service = services.newService(ServiceSubscriptionService, sling, uriTemplate)

	return subscriptionService
}
