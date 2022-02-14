package resources

type TentacleVersionDetails struct {
	UpgradeLocked    bool    `json:"UpgradeLocked"`
	UpgradeSuggested bool    `json:"UpgradeSuggested"`
	UpgradeRequired  bool    `json:"UpgradeRequired"`
	Version          *string `json:"Version"`
}

// NewTentacleVersionDetails creates and initializes tentacle version details.
func NewTentacleVersionDetails(version *string, upgradeLocked bool, upgradeSuggested bool, upgradeRequired bool) *TentacleVersionDetails {
	return &TentacleVersionDetails{
		Version:          version,
		UpgradeLocked:    upgradeLocked,
		UpgradeSuggested: upgradeSuggested,
		UpgradeRequired:  upgradeRequired,
	}
}
