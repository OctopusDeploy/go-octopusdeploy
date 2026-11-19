package api_test

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/stretchr/testify/require"
)

func getAIAgent(variables map[string]string) string {
	e := api.NewMockEnvironment()
	for k, v := range variables {
		e.Setenv(k, v)
	}
	return api.GetAIAgent(e)
}

func TestGetAIAgent_None(t *testing.T) {
	require.Equal(t, "", getAIAgent(nil))
}

func TestGetAIAgent_DetectsAgentFromVendorVariable(t *testing.T) {
	require.Equal(t, "AIAgent/ClaudeCode", getAIAgent(map[string]string{"CLAUDECODE": "1"}))
}

func TestGetAIAgent_ValueRulesRequireAMatchingValue(t *testing.T) {
	cases := []struct {
		variable string
		value    string
		expected string
	}{
		{"CURSOR_EXTENSION_HOST_ROLE", "agent-exec", "AIAgent/Cursor"},
		{"CURSOR_EXTENSION_HOST_ROLE", "user", ""},
		{"AGENT", "amp", "AIAgent/Amp"},
		{"AGENT", "teamcity-agent-01", ""},
	}

	for _, c := range cases {
		t.Run(c.variable+"="+c.value, func(t *testing.T) {
			require.Equal(t, c.expected, getAIAgent(map[string]string{c.variable: c.value}))
		})
	}
}

func TestGetAIAgent_IgnoresVariablesSetWithoutAnAgent(t *testing.T) {
	for _, variable := range []string{"CURSOR_TRACE_ID", "COPILOT_GITHUB_TOKEN"} {
		t.Run(variable, func(t *testing.T) {
			require.Equal(t, "", getAIAgent(map[string]string{variable: "1"}))
		})
	}
}

func TestGetAIAgent_IsDeterministicWhenAgentsAreNested(t *testing.T) {
	for i := 0; i < 50; i++ {
		require.Equal(t, "AIAgent/ClaudeCode", getAIAgent(map[string]string{"CURSOR_AGENT": "1", "CLAUDECODE": "1"}))
	}
}

func TestGetAIAgent_MapsAIAgentToAKnownName(t *testing.T) {
	cases := []struct {
		value    string
		expected string
	}{
		{"claude-code_2-1-322_agent", "AIAgent/ClaudeCode"},
		{"Claude Code", "AIAgent/ClaudeCode"},
		{"github_copilot_vscode_agent", "AIAgent/GitHubCopilot"},
		{"devin@1", "AIAgent/Devin"},
		{"custom-agent@2.0", "AIAgent/Other"},
	}

	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			require.Equal(t, c.expected, getAIAgent(map[string]string{"AI_AGENT": c.value}))
		})
	}
}

func TestGetAIAgent_AIAgentTakesPrecedenceOverVendorVariables(t *testing.T) {
	require.Equal(t, "AIAgent/ClaudeCode", getAIAgent(map[string]string{"CURSOR_AGENT": "1", "AI_AGENT": "claude-code_2-1-322_agent"}))
}

func TestGetAIAgent_UnrecognisedAIAgentFallsBackToVendorName(t *testing.T) {
	require.Equal(t, "AIAgent/GeminiCLI", getAIAgent(map[string]string{"GEMINI_CLI": "1", "AI_AGENT": "gcli"}))
}

func TestGetAIAgent_NeverIncludesRawAIAgentValue(t *testing.T) {
	require.Equal(t, "AIAgent/Other", getAIAgent(map[string]string{"AI_AGENT": "  evil) agent\r\nX-Injected: 1 "}))
	require.Equal(t, "", getAIAgent(map[string]string{"AI_AGENT": " \t "}))
}

func TestGetAutomationEnvironment_IgnoresAIAgents(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("CLAUDECODE", "1")
	e.Setenv("GITHUB_ACTIONS", "true")
	require.Equal(t, "GitHubActions", api.GetAutomationEnvironment(e))
	require.Equal(t, "AIAgent/ClaudeCode", api.GetAIAgent(e))
}
