package api_test

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAutomationEnvironment_None(t *testing.T) {
	e := api.NewMockEnvironment()
	result := api.GetAutomationEnvironment(e)
	assert.Equal(t, "NoneOrUnknown", result)
}

func TestGetAutomationEnvironment(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("GITHUB_ACTIONS", "GITHUB_ACTIONS")
	result := api.GetAutomationEnvironment(e)
	assert.Equal(t, "GitHubActions", result)
}

func TestGetAutomationEnvironment_TeamCity(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("TEAMCITY_VERSION", "2018.1.3 (Build 12345)")
	result := api.GetAutomationEnvironment(e)
	assert.Equal(t, "TeamCity/2018.1.3", result)
}
