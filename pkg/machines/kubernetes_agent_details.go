package machines

type KubernetesAgentDetails struct {
	AgentVersion        string `json:"AgentVersion"`
	TentacleVersion     string `json:"TentacleVersion"`
	UpgradeStatus       string `json:"UpgradeStatus" validate:"oneof=NoUpgrades UpgradeAvailable UpgradeSuggested UpgradeRequired"`
	HelmReleaseName     string `json:"HelmReleaseName"`
	KubernetesNamespace string `json:"KubernetesNamespace"`
}
