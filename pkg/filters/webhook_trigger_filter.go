package filters

type WebhookTriggerFilter struct {
	Password  string `json:"Password"`
	WebhookId string `json:"WebhookId,omitempty"`

	triggerFilter
}

func NewWebhookTriggerFilter(password string) *WebhookTriggerFilter {

	return &WebhookTriggerFilter{
		triggerFilter: *newTriggerFilter(WebhookFilter),
		Password:      password,
	}
}

func (t *WebhookTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *WebhookTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &WebhookTriggerFilter{}
