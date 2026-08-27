package filters

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type WebhookTriggerFilter struct {
	WebhookID     string               `json:"WebhookId,omitempty"`
	Secret        *core.SensitiveValue `json:"Secret,omitempty"`
	RequireAPIKey bool                 `json:"RequireApiKey"`

	triggerFilter
}

// NewWebhookTriggerFilter creates a webhook trigger filter. Incoming requests are authenticated with
// an Octopus API key when requireAPIKey is set, in which case secret must be nil; otherwise a secret is required.
func NewWebhookTriggerFilter(requireAPIKey bool, secret *core.SensitiveValue) *WebhookTriggerFilter {
	return &WebhookTriggerFilter{
		triggerFilter: *newTriggerFilter(WebhookFilter),
		RequireAPIKey: requireAPIKey,
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
