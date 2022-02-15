package service

import (
	"github.com/dghubble/sling"
)

type jiraIntegrationService struct {
	connectAppCredentialsTestPath string
	credentialsTestPath           string

	service
}

func newJiraIntegrationService(sling *sling.Sling, uriTemplate string, connectAppCredentialsTestPath string, credentialsTestPath string) *jiraIntegrationService {
	return &jiraIntegrationService{
		connectAppCredentialsTestPath: connectAppCredentialsTestPath,
		credentialsTestPath:           credentialsTestPath,
		service:                       newService(ServiceJiraIntegrationService, sling, uriTemplate),
	}
}
