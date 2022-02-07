package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type runbookRunService struct {
	services.canDeleteService
}

func newRunbookRunService(sling *sling.Sling, uriTemplate string) *runbookRunService {
	runbookRunService := &runbookRunService{}
	runbookRunService.service = services.newService(ServiceRunbookRunService, sling, uriTemplate)

	return runbookRunService
}
