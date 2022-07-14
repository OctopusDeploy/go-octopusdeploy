package machines

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/core"

// Project

// ValidProjectDefaultGuidedFailureModes provides options for "Default failure mode" - https://octopus.com/docs/deployment-process/releases/guided-failures
var ValidProjectDefaultGuidedFailureModes = []string{
	"EnvironmentDefault", "Off", "On",
}

// ValidProjectConnectivityPolicySkipMachineBehaviors provides options for "Skip Deployment Targets" - https://octopus.com/docs/deployment-patterns/elastic-and-transient-environments/deploying-to-transient-targets
var ValidProjectConnectivityPolicySkipMachineBehaviors = []core.SkipMachineBehavior{
	core.SkipMachineBehaviorNone,
	core.SkipMachineBehaviorSkipUnavailableMachines,
}

// ValidMachineStatuses provides options for valid machine status
var ValidMachineStatuses = []string{
	"Online", "Offline", "Unknown", "NeedsUpgrade", "CalamariNeedsUpgrade", "Disabled",
}
