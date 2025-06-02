package machinepolicies

type MachinePackageCacheRetentionPolicy struct {
	Strategy                 string `json:"Strategy" validate:"required,oneof=Default Quantities"`
	QuantityOfPackagesToKeep int32  `json:"QuantityOfPackagesToKeep,omitempty"`
	PackageUnit              string `json:"PackageUnit,omitempty" validate:"omitempty,oneof=Items"`
	QuantityOfVersionsToKeep int32  `json:"QuantityOfVersionsToKeep,omitempty"`
	VersionUnit              string `json:"VersionUnit,omitempty" validate:"omitempty,oneof=Items"`
}

func NewMachinePackageCacheRetentionPolicy() *MachinePackageCacheRetentionPolicy {
	return &MachinePackageCacheRetentionPolicy{
		Strategy: "Default",
	}
}
