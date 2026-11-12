package api_test

import (
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/stretchr/testify/require"
)

func TestGetAIAgent_None(t *testing.T) {
	e := api.NewMockEnvironment()
	require.Equal(t, "", api.GetAIAgent(e))
}

func TestGetAIAgent_VendorVariables(t *testing.T) {
	cases := []struct {
		variable string
		value    string
		expected string
	}{
		{"CLAUDECODE", "1", "AIAgent/ClaudeCode"},
		{"CLAUDE_CODE_ENTRYPOINT", "cli", "AIAgent/ClaudeCode"},
		{"CURSOR_AGENT", "1", "AIAgent/Cursor"},
		{"CURSOR_EXTENSION_HOST_ROLE", "agent-exec", "AIAgent/Cursor"},
		{"CODEX_SANDBOX", "seatbelt", "AIAgent/Codex"},
		{"CODEX_THREAD_ID", "abc", "AIAgent/Codex"},
		{"GEMINI_CLI", "1", "AIAgent/GeminiCLI"},
		{"COPILOT_CLI", "1", "AIAgent/GitHubCopilot"},
		{"OPENCODE", "1", "AIAgent/OpenCode"},
		{"OPENCODE_CLIENT", "cli", "AIAgent/OpenCode"},
		{"ANTIGRAVITY_AGENT", "1", "AIAgent/Antigravity"},
		{"GOOSE_TERMINAL", "1", "AIAgent/Goose"},
		{"AGENT", "goose", "AIAgent/Goose"},
		{"AGENT", "amp", "AIAgent/Amp"},
		{"CLINE_ACTIVE", "true", "AIAgent/Cline"},
		{"ROO_CODE_TASK_ID", "abc", "AIAgent/RooCode"},
		{"AUGMENT_AGENT", "1", "AIAgent/Augment"},
		{"QWEN_CODE", "1", "AIAgent/QwenCode"},
		{"DROID_CLI", "1", "AIAgent/FactoryDroid"},
		{"KIMI_CLI", "1", "AIAgent/KimiCLI"},
		{"TRAE_AI_SHELL_ID", "abc", "AIAgent/Trae"},
		{"OR_APP_NAME", "aider", "AIAgent/Aider"},
		{"OR_APP_NAME", "OpenHands", "AIAgent/OpenHands"},
	}

	for _, c := range cases {
		t.Run(c.variable+"="+c.value, func(t *testing.T) {
			e := api.NewMockEnvironment()
			e.Setenv(c.variable, c.value)
			require.Equal(t, c.expected, api.GetAIAgent(e))
		})
	}
}

func TestGetAIAgent_ValueRulesRequireMatch(t *testing.T) {
	cases := map[string]string{
		"CURSOR_EXTENSION_HOST_ROLE": "user",
		"OPENCODE":                   "0",
		"AGENT":                      "teamcity-agent-01",
		"OR_APP_NAME":                "SomethingElse",
	}

	for variable, value := range cases {
		t.Run(variable, func(t *testing.T) {
			e := api.NewMockEnvironment()
			e.Setenv(variable, value)
			require.Equal(t, "", api.GetAIAgent(e))
		})
	}
}

func TestGetAIAgent_ExcludesNoisyVariables(t *testing.T) {
	for _, variable := range []string{"CURSOR_TRACE_ID", "COPILOT_GITHUB_TOKEN", "GOOSE_PROVIDER", "REPL_ID", "ZED_TERM"} {
		t.Run(variable, func(t *testing.T) {
			e := api.NewMockEnvironment()
			e.Setenv(variable, "1")
			require.Equal(t, "", api.GetAIAgent(e))
		})
	}
}

func TestGetAIAgent_IsDeterministicWhenAgentsAreNested(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("CURSOR_AGENT", "1")
	e.Setenv("CLAUDECODE", "1")
	for i := 0; i < 50; i++ {
		require.Equal(t, "AIAgent/ClaudeCode", api.GetAIAgent(e))
	}
}

func TestGetAIAgent_AIAgentVariable(t *testing.T) {
	cases := []struct {
		value    string
		expected string
	}{
		{"claude-code_2-1-322_agent", "AIAgent/ClaudeCode"},
		{"claude-code", "AIAgent/ClaudeCode"},
		{"Claude-Code", "AIAgent/ClaudeCode"},
		{"Claude Code", "AIAgent/ClaudeCode"},
		{"cursor-cli", "AIAgent/Cursor"},
		{"github_copilot_vscode_agent", "AIAgent/GitHubCopilot"},
		{"github-copilot-cli", "AIAgent/GitHubCopilot"},
		{"devin@1", "AIAgent/Devin"},
		{"custom-agent@2.0", "AIAgent/Other"},
	}

	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			e := api.NewMockEnvironment()
			e.Setenv("AI_AGENT", c.value)
			require.Equal(t, c.expected, api.GetAIAgent(e))
		})
	}
}

func TestGetAIAgent_AIAgentTakesPrecedenceOverVendorVariables(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("CURSOR_AGENT", "1")
	e.Setenv("AI_AGENT", "claude-code_2-1-322_agent")
	require.Equal(t, "AIAgent/ClaudeCode", api.GetAIAgent(e))
}

func TestGetAIAgent_UnrecognisedAIAgentFallsBackToVendorName(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("GEMINI_CLI", "1")
	e.Setenv("AI_AGENT", "gcli")
	require.Equal(t, "AIAgent/GeminiCLI", api.GetAIAgent(e))
}

func TestGetAIAgent_NeverIncludesRawAIAgentValue(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("AI_AGENT", "  evil) agent\r\nX-Injected: 1 ")
	require.Equal(t, "AIAgent/Other", api.GetAIAgent(e))

	e.Setenv("AI_AGENT", strings.Repeat("a", 200))
	require.Equal(t, "AIAgent/Other", api.GetAIAgent(e))

	e.Setenv("AI_AGENT", " \t ")
	require.Equal(t, "", api.GetAIAgent(e))
}

func TestGetAutomationEnvironment_IgnoresAIAgents(t *testing.T) {
	e := api.NewMockEnvironment()
	e.Setenv("CLAUDECODE", "1")
	e.Setenv("GITHUB_ACTIONS", "true")
	require.Equal(t, "GitHubActions", api.GetAutomationEnvironment(e))
	require.Equal(t, "AIAgent/ClaudeCode", api.GetAIAgent(e))
}
