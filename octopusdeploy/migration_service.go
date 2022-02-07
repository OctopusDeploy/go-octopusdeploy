package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type migrationService struct {
	migrationsImportPath        string
	migrationsPartialExportPath string

	services.service
}

func newMigrationService(sling *sling.Sling, uriTemplate string, migrationsImportPath string, migrationsPartialExportPath string) *migrationService {
	return &migrationService{
		migrationsImportPath:        migrationsImportPath,
		migrationsPartialExportPath: migrationsPartialExportPath,
		service:                     services.newService(ServiceMigrationService, sling, uriTemplate),
	}
}
