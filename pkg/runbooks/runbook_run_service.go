package runbooks

import (
	"github.com/dghubble/sling"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
)

type RunbookRunService struct {
	services.CanDeleteService
}

func NewRunbookRunService(sling *sling.Sling, uriTemplate string) *RunbookRunService {
	return &RunbookRunService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceRunbookRunService, sling, uriTemplate),
		},
	}
}
