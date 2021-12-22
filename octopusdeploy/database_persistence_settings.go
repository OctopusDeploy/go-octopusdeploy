package octopusdeploy

// DatabasePersistenceSettings represents database persistence settings associated with a project.
type DatabasePersistenceSettings struct {
	persistenceSettings
}

// NewDatabasePersistenceSettings creates an instance of database persistence settings.
func NewDatabasePersistenceSettings() *DatabasePersistenceSettings {
	return &DatabasePersistenceSettings{
		persistenceSettings: persistenceSettings{Type: "Database"},
	}
}

// GetType returns the type for this persistence settings.
func (d *DatabasePersistenceSettings) GetType() string {
	return d.Type
}

var _ IPersistenceSettings = &DatabasePersistenceSettings{}
