package octopusdeploy

type ActionType int

const (
	AutoDeploy ActionType = iota
	DeployLatestRelease
	DeployNewRelease
	RunRunbook
)
