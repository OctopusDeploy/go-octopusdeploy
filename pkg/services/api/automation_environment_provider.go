package api

import (
	"fmt"
	"strings"
)

var knownEnvironments = map[string][]string{
	"Octopus": {"AgentProgramDirectoryPath"},

	// https://confluence.jetbrains.com/display/TCD9/Predefined+Build+Parameters
	"TeamCity": {"TEAMCITY_VERSION"},

	// https://docs.github.com/en/actions/reference/environment-variables#default-environment-variables
	"GitHubActions": {"GITHUB_ACTIONS"},

	// https://docs.microsoft.com/en-us/azure/devops/pipelines/build/variables
	"AzureDevOps": {"TF_BUILD", "BUILD_BUILDID", "AGENT_WORKFOLDER"},

	// https://confluence.atlassian.com/bamboo/bamboo-variables-289277087.html
	"Bamboo":       {"bamboo_agentId"},
	"AppVeyor":     {"APPVEYOR"},
	"BitBucket":    {"BITBUCKET_BUILD_NUMBER"},
	"Jenkins":      {"JENKINS_URL"},
	"CircleCI":     {"CIRCLECI"},
	"GitLabCI":     {"GITLAB_CI"},
	"Travis":       {"TRAVIS"},
	"GoCD":         {"GO_PIPELINE_LABEL"},
	"BitRise":      {"BITRISE_IO"},
	"Buddy":        {"BUDDY_WORKSPACE_ID"},
	"BuildKite":    {"BUILDKITE"},
	"CirrusCI":     {"CIRRUS_CI"},
	"AWSCodeBuild": {"CODEBUILD_BUILD_ARN"},
	"Codeship":     {"CI_NAME"},
	"Drone":        {"DRONE"},
	"Dsari":        {"DSARI"},
	"Hudson":       {"HUDSON_URL"},
	"MagnumCI":     {"MAGNUM"},
	"SailCI":       {"SAILCI"},
	"Semaphore":    {"SEMAPHORE"},
	"Shippable":    {"SHIPPABLE"},
	"SolanoCI":     {"TDDIUM"},
	"StriderCD":    {"STRIDER"},
}

func GetAutomationEnvironment(e Environment) string {
	var environment = "NoneOrUnknown"
	for key, variables := range knownEnvironments {
		for _, v := range variables {
			if e.Getenv(v) != "" {
				environment = key
			}
		}
	}

	if environment == "TeamCity" {
		value := e.Getenv(knownEnvironments[environment][0])
		parts := strings.Split(value, " ")
		environment = fmt.Sprintf("%s/%s", environment, parts[0])
	}

	return environment
}
