package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
)

type jiraIntegrationService struct {
	connectAppCredentialsTestPath string
	credentialsTestPath           string

	services.service
}

func newJiraIntegrationService(sling *sling.Sling, uriTemplate string, connectAppCredentialsTestPath string, credentialsTestPath string) *jiraIntegrationService {
	return &jiraIntegrationService{
		connectAppCredentialsTestPath: connectAppCredentialsTestPath,
		credentialsTestPath:           credentialsTestPath,
		service:                       services.newService(ServiceJiraIntegrationService, sling, uriTemplate),
	}
}
