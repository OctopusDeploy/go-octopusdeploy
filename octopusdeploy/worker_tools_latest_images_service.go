package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type workerToolsLatestImageService struct {
	services.service
}

func newWorkerToolsLatestImageService(sling *sling.Sling, uriTemplate string) *workerToolsLatestImageService {
	return &workerToolsLatestImageService{
		service: services.newService(ServiceWorkerToolsLatestImageService, sling, uriTemplate),
	}
}
