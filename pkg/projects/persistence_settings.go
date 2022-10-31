package projects

// PersistenceSettings defines the interface for persistence settings.
type PersistenceSettings interface {
	GetType() PersistenceSettingsType
}

// persistenceSettings represents persistence settings associated with a project.
type persistenceSettings struct {
	Type PersistenceSettingsType `json:"Type"`
}
