package service

import (
	"github.com/dghubble/sling"
)

type azureDevOpsConnectivityCheckService struct {
	service
}

func newAzureDevOpsConnectivityCheckService(sling *sling.Sling, uriTemplate string) *azureDevOpsConnectivityCheckService {
	return &azureDevOpsConnectivityCheckService{
		service: newService(ServiceAzureDevOpsConnectivityCheckService, sling, uriTemplate),
	}
}
