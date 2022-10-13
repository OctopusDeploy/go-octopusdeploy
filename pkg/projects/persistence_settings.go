package projects

// IPersistenceSettings defines the interface for persistence settings.
type IPersistenceSettings interface {
	GetType() PersistenceSettingsType
}

// persistenceSettings represents persistence settings associated with a project.
type persistenceSettings struct {
	Type PersistenceSettingsType `json:"Type"`
}
