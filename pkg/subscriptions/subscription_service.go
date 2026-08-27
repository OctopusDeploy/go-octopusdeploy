package subscriptions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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

const template = "/api/{spaceId}/subscriptions{/id}{?skip,take,ids,partialName,spaces}"

// Add creates a new subscription.
func Add(client newclient.Client, spaceID string, subscription *Subscription) (*Subscription, error) {
	if subscription == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterSubscription)
	}
	return newclient.Add[Subscription](client, template, spaceID, subscription)
}

// GetByID returns the subscription that matches the input ID.
func GetByID(client newclient.Client, spaceID string, id string) (*Subscription, error) {
	return newclient.GetByID[Subscription](client, template, spaceID, id)
}

// Update modifies a subscription based on the one provided as input.
func Update(client newclient.Client, spaceID string, subscription *Subscription) (*Subscription, error) {
	if subscription == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterSubscription)
	}
	return newclient.Update[Subscription](client, template, spaceID, subscription.ID, subscription)
}

// DeleteByID deletes the subscription that matches the input ID.
func DeleteByID(client newclient.Client, spaceID string, id string) error {
	return newclient.DeleteByID(client, template, spaceID, id)
}

// GetAll returns all subscriptions for the given space.
func GetAll(client newclient.Client, spaceID string) ([]*Subscription, error) {
	return newclient.GetAll[Subscription](client, template, spaceID)
}

// Get returns a collection of subscriptions based on the criteria defined by its input query parameter.
func Get(client newclient.Client, spaceID string, query SubscriptionsQuery) (*resources.Resources[*Subscription], error) {
	return newclient.GetByQuery[Subscription](client, template, spaceID, query)
}
