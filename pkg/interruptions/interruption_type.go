package interruptions

type InterruptionType string

const (
	InterruptionTypeManualIntervention             = InterruptionType("ManualIntervention")
	InterruptionTypeGuidedFailure                  = InterruptionType("GuidedFailure")
	InterruptionTypePullRequestCompletion          = InterruptionType("PullRequestCompletion")
	InterruptionTypeArgoCDApplicationSync          = InterruptionType("ArgoCDApplicationSync")
	InterruptionTypeKubernetesResourceVerification = InterruptionType("KubernetesResourceVerification")
)
