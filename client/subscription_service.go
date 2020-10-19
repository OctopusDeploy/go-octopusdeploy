package client

import "github.com/dghubble/sling"

type subscriptionService struct {
	service
}

func newSubscriptionService(sling *sling.Sling, uriTemplate string) *subscriptionService {
	subscriptionService := &subscriptionService{}
	subscriptionService.service = newService(serviceSubscriptionService, sling, uriTemplate, nil)

	return subscriptionService
}
