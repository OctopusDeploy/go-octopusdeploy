package resources

type ActionType int

const (
	AutoDeploy ActionType = iota
	DeployLatestRelease
	DeployNewRelease
	RunRunbook
)
