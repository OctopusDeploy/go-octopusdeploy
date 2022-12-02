package extensions

type ExtensionID string

const (
	ExtensionIDJira                  = ExtensionID("jira-integration")
	ExtensionIDJiraServiceManagement = ExtensionID("jiraservicemanagement-integration")
	ExtensionIDServiceNow            = ExtensionID("servicenow-integration")
)
