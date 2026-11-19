package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendAIAgent_KeepsTheExistingUserAgentAsAPrefix(t *testing.T) {
	const userAgent = "go-octopusdeploy/2.0.0 (linux; amd64) go/go1.24.7 GitHubActions octopus/2.18.0 (release;deploy)"

	e := NewMockEnvironment()
	require.Equal(t, userAgent, appendAIAgent(userAgent, e))

	e.Setenv("CLAUDECODE", "1")
	require.Equal(t, userAgent+" (ai-agent=ClaudeCode)", appendAIAgent(userAgent, e))
}
