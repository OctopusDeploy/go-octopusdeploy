package octopusdeploy

type OfflinePackageDropDestination struct {
	DestinationType string `json:"DestinationType" validate:"oneof=Artifact FileSystem"`
	DropFolderPath  string `json:"DropFolderPath,omitempty"`
}

func NewOfflinePackageDropDestination() *OfflinePackageDropDestination {
	return &OfflinePackageDropDestination{
		DestinationType: "Artifact",
	}
}
