package service

import (
	"github.com/dghubble/sling"
)

type subscriptionService struct {
	canDeleteService
}

func newSubscriptionService(sling *sling.Sling, uriTemplate string) *subscriptionService {
	subscriptionService := &subscriptionService{}
	subscriptionService.service = newService(ServiceSubscriptionService, sling, uriTemplate)

	return subscriptionService
}
