package subscriptions

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

// TeamsChannelSubscriptionTarget is one Microsoft Teams destination a subscription notifies.
// Type is "Webhook" (incoming webhook URL) or "AppChannel" (Octopus Teams app channel).
// WebhookUrl is sensitive and write-only: the API accepts a NewValue on write but never returns it; reads show HasValue only.
type TeamsChannelSubscriptionTarget struct {
	Id         string               `json:"Id"`
	Type       string               `json:"Type,omitempty"`
	Name       string               `json:"Name"`
	WebhookUrl *core.SensitiveValue `json:"WebhookUrl,omitempty"`
	ChannelId  *string              `json:"ChannelId,omitempty"`
	TeamId     *string              `json:"TeamId,omitempty"`
	TeamName   *string              `json:"TeamName,omitempty"`
}

type EventNotificationSubscriptionFilter struct {
	DocumentTypes   []string `json:"DocumentTypes"`
	EventAgents     []string `json:"EventAgents"`
	EventCategories []string `json:"EventCategories"`
	EventGroups     []string `json:"EventGroups"`
	Environments    []string `json:"Environments"`
	ProjectGroups   []string `json:"ProjectGroups"`
	Projects        []string `json:"Projects"`
	Tags            []string `json:"Tags"`
	Tenants         []string `json:"Tenants"`
	Users           []string `json:"Users"`
}

type EventNotificationSubscription struct {
	EmailDigestLastProcessed            *time.Time                           `json:"EmailDigestLastProcessed,omitempty"`
	EmailDigestLastProcessedEventAutoId *int64                               `json:"EmailDigestLastProcessedEventAutoId,omitempty"`
	EmailFrequencyPeriod                string                               `json:"EmailFrequencyPeriod"`
	EmailPriority                       string                               `json:"EmailPriority"`
	EmailShowDatesInTimeZoneId          string                               `json:"EmailShowDatesInTimeZoneId"`
	EmailTeams                          []string                             `json:"EmailTeams"`
	Filter                              *EventNotificationSubscriptionFilter `json:"Filter"`
	SlackChannelIds                     []string                             `json:"SlackChannelIds"`
	SlackChannelNames                   []string                             `json:"SlackChannelNames"`
	// Deprecated: SlackDigestFormat is no longer used by Octopus Server; Slack digests always send a summary. It will be removed in a future release.
	SlackDigestFormat                   string                           `json:"SlackDigestFormat"`
	SlackFrequencyPeriod                string                           `json:"SlackFrequencyPeriod"`
	TeamsDigestLastProcessed            *time.Time                       `json:"TeamsDigestLastProcessed,omitempty"`
	TeamsDigestLastProcessedEventAutoId *int64                           `json:"TeamsDigestLastProcessedEventAutoId,omitempty"`
	TeamsFrequencyPeriod                string                           `json:"TeamsFrequencyPeriod"`
	TeamsChannels                       []TeamsChannelSubscriptionTarget `json:"TeamsChannels"`
	WebhookHeaderKey                    string                           `json:"WebhookHeaderKey"`
	WebhookHeaderValue                  string                           `json:"WebhookHeaderValue"`
	WebhookLastProcessed                *time.Time                       `json:"WebhookLastProcessed,omitempty"`
	WebhookLastProcessedEventAutoId     *int64                           `json:"WebhookLastProcessedEventAutoId,omitempty"`
	WebhookTeams                        []string                         `json:"WebhookTeams"`
	WebhookTimeout                      string                           `json:"WebhookTimeout"`
	WebhookURI                          string                           `json:"WebhookURI"`
}

type Subscription struct {
	EventNotificationSubscription *EventNotificationSubscription `json:"EventNotificationSubscription"`
	IsDisabled                    bool                           `json:"IsDisabled"`
	Name                          string                         `json:"Name" validate:"required"`
	SpaceID                       string                         `json:"SpaceId,omitempty"`
	Type                          string                         `json:"Type,omitempty"`

	resources.Resource
}

func NewSubscription(name string) *Subscription {
	return &Subscription{
		Name:     name,
		Resource: *resources.NewResource(),
		EventNotificationSubscription: &EventNotificationSubscription{
			EmailFrequencyPeriod:       "01:00:00",
			EmailPriority:              "Normal",
			EmailShowDatesInTimeZoneId: "UTC",
			EmailTeams:                 []string{},
			SlackChannelIds:            []string{},
			SlackChannelNames:          []string{},
			SlackDigestFormat:          "Summary",
			SlackFrequencyPeriod:       "01:00:00",
			TeamsChannels:              []TeamsChannelSubscriptionTarget{},
			TeamsFrequencyPeriod:       "01:00:00",
			WebhookTeams:               []string{},
			WebhookTimeout:             "00:00:10",
			Filter: &EventNotificationSubscriptionFilter{
				DocumentTypes:   []string{},
				EventAgents:     []string{},
				EventCategories: []string{},
				EventGroups:     []string{},
				Environments:    []string{},
				ProjectGroups:   []string{},
				Projects:        []string{},
				Tags:            []string{},
				Tenants:         []string{},
				Users:           []string{},
			},
		},
	}
}

func (s *Subscription) GetName() string     { return s.Name }
func (s *Subscription) SetName(name string) { s.Name = name }

func (s *Subscription) Validate() error {
	return validator.New().Struct(s)
}

var _ resources.IHasName = &Subscription{}
