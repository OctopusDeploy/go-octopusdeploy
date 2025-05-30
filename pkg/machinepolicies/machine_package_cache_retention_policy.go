package machinepolicies

type MachinePackageCacheRetentionUnit int

const (
	Item MachinePackageCacheRetentionUnit = iota
)

type MachinePackageCacheRetentionStrategy int

const (
	Default MachinePackageCacheRetentionStrategy = iota
	Quantities
)

type MachinePackageCacheRetentionPolicy struct {
	Strategy                 MachinePackageCacheRetentionStrategy `json:"Strategy" validate:"required"`
	QuantityOfPackagesToKeep *int32                               `json:"QuantityOfPackagesToKeep"`
	PackageUnit              *MachinePackageCacheRetentionUnit    `json:"PackageUnit"`
	QuantityOfVersionsToKeep *int32                               `json:"QuantityOfVersionsToKeep"`
	VersionUnit              *MachinePackageCacheRetentionUnit    `json:"VersionUnit"`
}

func NewMachinePackageCacheRetentionPolicy() *MachinePackageCacheRetentionPolicy {
	return &MachinePackageCacheRetentionPolicy{
		Strategy: Default,
	}
}
