package jira

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
)

type JiraIntegrationService struct {
	connectAppCredentialsTestPath string
	credentialsTestPath           string

	services.Service
}

func NewJiraIntegrationService(sling *sling.Sling, uriTemplate string, connectAppCredentialsTestPath string, credentialsTestPath string) *JiraIntegrationService {
	return &JiraIntegrationService{
		connectAppCredentialsTestPath: connectAppCredentialsTestPath,
		credentialsTestPath:           credentialsTestPath,
		Service:                       services.NewService(constants.ServiceJiraIntegrationService, sling, uriTemplate),
	}
}
