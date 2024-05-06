package machines

type KubernetesAgentDetails struct {
	AgentVersion        string `json:"AgentVersion"`
	TentacleVersion     string `json:"TentacleVersion"`
	UpgradeStatus       string `json:"UpgradeStatus"`
	HelmReleaseName     string `json:"HelmReleaseName"`
	KubernetesNamespace string `json:"KubernetesNamespace"`
}
