package ephemeralenvironments

type CreateEnvironmentResponse struct {
	Id string `json:"Id"`
}

type CreateEnvironmentCommand struct {
	EnvironmentName string `json:"EnvironmentName"`
	SpaceID         string `uri:"spaceId"`
	ProjectID       string `uri:"projectId"`
}

type DeprovisionEphemeralEnvironmentProjectCommand struct{}

type DeprovisionEphemeralEnvironmentProjectResponse struct {
	DeprovisioningRun DeprovisioningRunbookRun `json:"DeprovisioningRun"`
}

type DeprovisionEphemeralEnvironmentCommand struct{}

type DeprovisionEphemeralEnvironmentResponse struct {
	DeprovisioningRuns []DeprovisioningRunbookRun `json:"DeprovisioningRuns"`
}

type DeprovisioningRunbookRun struct {
	RunbookRunID string `json:"RunbookRunId"`
	TaskId       string `json:"TaskId"`
}
