package v2

const (
	Template                                          = "/api/{spaceId}/environments/v2{?name,skip,ids,take,partialName,type}"
	CreateEphemeralEnvironmentTemplate                = "/api/{spaceId}/projects/{projectId}/environments/ephemeral"
	DeprovisionEphemeralEnvironmentForProjectTemplate = "/api/{spaceId}/projects/{projectId}/environments/ephemeral/{id}/deprovision"
	DeprovisionEphemeralEnvironmentTemplate           = "/api/{spaceId}/environments/ephemeral/{id}/deprovision"
	GetEphemeralEnvironmentStatusTemplate             = "/api/{spaceId}/projects/{projectId}/environments/ephemeral/{id}/status"
	DeleteEphemeralEnvironmentTemplate                = "/api/{spaceId}/environments/{id}"
)
