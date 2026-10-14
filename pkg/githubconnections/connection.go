package githubconnections

// ConnectionStatus describes the health of a GitHub App connection.
type ConnectionStatus string

const (
	ConnectionStatusConnected             = ConnectionStatus("Connected")
	ConnectionStatusConnectionNotFound    = ConnectionStatus("ConnectionNotFound")
	ConnectionStatusInstallationNotFound  = ConnectionStatus("InstallationNotFound")
	ConnectionStatusInstallationSuspended = ConnectionStatus("InstallationSuspended")
	ConnectionStatusError                 = ConnectionStatus("Error")
)

// Installation represents the GitHub App installation backing a connection.
type Installation struct {
	InstallationID   string `json:"InstallationId"`
	AccountID        string `json:"AccountId"`
	AccountLogin     string `json:"AccountLogin"`
	AccountAvatarURL string `json:"AccountAvatarUrl"`
	AccountType      string `json:"AccountType"`
	// AllRepositories is true when the installation can access every repository in the
	// account, false when it is restricted to a selected set.
	AllRepositories bool `json:"AllRepositories"`
}

// Connection represents a GitHub App connection.
type Connection struct {
	ID           string           `json:"Id"`
	Status       ConnectionStatus `json:"Status,omitempty"`
	Installation *Installation    `json:"Installation,omitempty"`
}
