package resources

type SkipMachineBehavior string

const (
	SkipMachineBehaviorNone                    = SkipMachineBehavior("None")
	SkipMachineBehaviorSkipUnavailableMachines = SkipMachineBehavior("SkipUnavailableMachines")
)
