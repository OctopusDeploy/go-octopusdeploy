package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type azureDevOpsConnectivityCheckService struct {
	services.service
}

func newAzureDevOpsConnectivityCheckService(sling *sling.Sling, uriTemplate string) *azureDevOpsConnectivityCheckService {
	return &azureDevOpsConnectivityCheckService{
		service: services.newService(ServiceAzureDevOpsConnectivityCheckService, sling, uriTemplate),
	}
}
