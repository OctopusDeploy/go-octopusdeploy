package v2

const Template = "/api/{spaceId}/environments/v2{?name,skip,ids,take,partialName,type}"
const CreateEphemeralEnvironmentTemplate = "/api/{spaceId}/projects/{projectId}/environments/ephemeral"
const DeprovisionEphemeralEnvironmentForProjectTemplate = "/api/{spaceId}/projects/{projectId}/environments/ephemeral/{id}/deprovision"
const DeprovisionEphemeralEnvironmentTemplate = "/api/{spaceId}/environments/ephemeral/{id}/deprovision"
