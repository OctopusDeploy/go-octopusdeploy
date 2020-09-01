package model

type MachineTentacleVersionDetails struct {
	UpgradeLocked    bool   `json:"UpgradeLocked"`
	Version          string `json:"Version"`
	UpgradeSuggested bool   `json:"UpgradeSuggested"`
	UpgradeRequired  bool   `json:"UpgradeRequired"`
}
