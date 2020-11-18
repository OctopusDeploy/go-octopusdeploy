package octopusdeploy

type OfflineDropDestination struct {
	DestinationType string `json:"DestinationType" validate:"oneof=Artifact FileSystem"`
	DropFolderPath  string `json:"DropFolderPath,omitempty"`
}

func NewOfflineDropDestination() *OfflineDropDestination {
	return &OfflineDropDestination{}
}
