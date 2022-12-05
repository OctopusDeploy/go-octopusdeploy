package extensions

type ExtensionID string

const (
	JiraExtensionID                  = ExtensionID("jira-integration")
	JiraServiceManagementExtensionID = ExtensionID("jiraservicemanagement-integration")
	ServiceNowExtensionID            = ExtensionID("servicenow-integration")
)
