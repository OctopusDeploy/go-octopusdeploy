package workertoolslatestimages

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type WorkerToolsLatestImageService struct {
	services.Service
}

func NewWorkerToolsLatestImageService(sling *sling.Sling, uriTemplate string) *WorkerToolsLatestImageService {
	return &WorkerToolsLatestImageService{
		Service: services.NewService(constants.ServiceWorkerToolsLatestImageService, sling, uriTemplate),
	}
}
