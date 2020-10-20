package octopusdeploy

import "github.com/dghubble/sling"

type workerToolsLatestImageService struct {
	service
}

func newWorkerToolsLatestImageService(sling *sling.Sling, uriTemplate string) *workerToolsLatestImageService {
	return &workerToolsLatestImageService{
		service: newService(serviceWorkerToolsLatestImageService, sling, uriTemplate, nil),
	}
}
