package migrations

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type MigrationService struct {
	migrationsImportPath        string
	migrationsPartialExportPath string

	services.Service
}

func NewMigrationService(sling *sling.Sling, uriTemplate string, migrationsImportPath string, migrationsPartialExportPath string) *MigrationService {
	return &MigrationService{
		migrationsImportPath:        migrationsImportPath,
		migrationsPartialExportPath: migrationsPartialExportPath,
		Service:                     services.NewService(constants.ServiceMigrationService, sling, uriTemplate),
	}
}
