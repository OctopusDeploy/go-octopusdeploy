package octopusdeploy

import "github.com/dghubble/sling"

type runbookRunService struct {
	service
}

func newRunbookRunService(sling *sling.Sling, uriTemplate string) *runbookRunService {
	runbookRunService := &runbookRunService{}
	runbookRunService.service = newService(serviceRunbookRunService, sling, uriTemplate, nil)

	return runbookRunService
}
