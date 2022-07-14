package devops

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type AzureDevOpsConnectivityCheckService struct {
	services.Service
}

func NewAzureDevOpsConnectivityCheckService(sling *sling.Sling, uriTemplate string) *AzureDevOpsConnectivityCheckService {
	return &AzureDevOpsConnectivityCheckService{
		Service: services.NewService(constants.ServiceAzureDevOpsConnectivityCheckService, sling, uriTemplate),
	}
}
