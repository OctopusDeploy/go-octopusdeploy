package filters

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type WebhookTriggerFilter struct {
	Secret    core.SensitiveValue `json:"Secret"`
	WebhookId string              `json:"WebhookId,omitempty"`

	triggerFilter
}

func NewWebhookTriggerFilter(secret core.SensitiveValue) *WebhookTriggerFilter {

	return &WebhookTriggerFilter{
		triggerFilter: *newTriggerFilter(WebhookFilter),
		Secret:        secret,
	}
}

func (t *WebhookTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *WebhookTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &WebhookTriggerFilter{}
