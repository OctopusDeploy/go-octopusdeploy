package api

import (
	"fmt"
	"strings"
)

// AI agents are detected separately from automation (CI) environments because the two are
// independent dimensions: an agent commonly runs inside a CI system or on an Octopus worker,
// and both should be reported rather than one masking the other.

// aiAgentEnvVar is the emerging cross-vendor convention for an agent to identify itself.
// https://github.com/vercel/vercel/tree/main/packages/detect-agent
const aiAgentEnvVar = "AI_AGENT"

const maxAIAgentValueLength = 64

type aiAgentRule struct {
	variable string
	// value, when set, must match the variable's value (case-insensitive). When empty, any
	// non-empty value matches.
	value string
}

type aiAgent struct {
	name  string
	rules []aiAgentRule
}

// knownAIAgents is ordered: the first agent with a matching rule wins, so the result is
// deterministic when several agents' variables are present (e.g. one agent running another).
//
// Variables that are commonly present without an agent driving the process are deliberately
// excluded: CURSOR_TRACE_ID (any Cursor terminal), COPILOT_GITHUB_TOKEN and GOOSE_PROVIDER
// (user configuration), REPL_ID (any Replit workspace) and ZED_TERM (any Zed terminal).
//
// References:
// https://code.claude.com/docs/en/env-vars
// https://source.dot.net/aspire/Telemetry/CodingAgentDetector.cs.html
// https://docs.rs/is-ai-agent
var knownAIAgents = []aiAgent{
	{"ClaudeCode", []aiAgentRule{{variable: "CLAUDECODE"}, {variable: "CLAUDE_CODE_ENTRYPOINT"}}},
	{"Cursor", []aiAgentRule{{variable: "CURSOR_AGENT"}, {variable: "CURSOR_EXTENSION_HOST_ROLE", value: "agent-exec"}}},
	{"Codex", []aiAgentRule{{variable: "CODEX_SANDBOX"}, {variable: "CODEX_THREAD_ID"}, {variable: "CODEX_CI"}}},
	{"GeminiCLI", []aiAgentRule{{variable: "GEMINI_CLI"}}},
	{"GitHubCopilot", []aiAgentRule{{variable: "COPILOT_CLI"}, {variable: "COPILOT_AGENT"}, {variable: "GITHUB_COPILOT_CLI_MODE"}}},
	{"OpenCode", []aiAgentRule{{variable: "OPENCODE", value: "1"}, {variable: "OPENCODE_CLIENT"}, {variable: "OPENCODE_AI"}}},
	{"Antigravity", []aiAgentRule{{variable: "ANTIGRAVITY_AGENT"}}},
	{"Goose", []aiAgentRule{{variable: "GOOSE_TERMINAL"}, {variable: "AGENT", value: "goose"}}},
	{"Amp", []aiAgentRule{{variable: "AGENT", value: "amp"}}},
	{"Cline", []aiAgentRule{{variable: "CLINE_ACTIVE"}, {variable: "CLINE_TASK_ID"}}},
	{"RooCode", []aiAgentRule{{variable: "ROO_CODE_TASK_ID"}}},
	{"Augment", []aiAgentRule{{variable: "AUGMENT_AGENT"}}},
	{"QwenCode", []aiAgentRule{{variable: "QWEN_CODE"}}},
	{"FactoryDroid", []aiAgentRule{{variable: "DROID_CLI"}}},
	{"KimiCLI", []aiAgentRule{{variable: "KIMI_CLI"}}},
	{"Trae", []aiAgentRule{{variable: "TRAE_AI_SHELL_ID"}}},
	{"Aider", []aiAgentRule{{variable: "OR_APP_NAME", value: "Aider"}}},
	{"OpenHands", []aiAgentRule{{variable: "OR_APP_NAME", value: "OpenHands"}}},
}

// aiAgentValuePrefixes maps the start of a normalised AI_AGENT value to a known agent name.
// Vendors don't agree on a format (e.g. "claude-code_2-1-322_agent", "github_copilot_vscode_agent",
// "devin@1"), so values are lowercased and underscores replaced with hyphens before matching.
var aiAgentValuePrefixes = []struct {
	prefix string
	name   string
}{
	{"claude-code", "ClaudeCode"},
	{"cursor", "Cursor"},
	{"codex", "Codex"},
	{"gemini", "GeminiCLI"},
	{"github-copilot", "GitHubCopilot"},
	{"opencode", "OpenCode"},
	{"antigravity", "Antigravity"},
	{"goose", "Goose"},
	{"amp", "Amp"},
	{"cline", "Cline"},
	{"augment", "Augment"},
	{"devin", "Devin"},
	{"replit", "Replit"},
	{"v0", "V0"},
}

// GetAIAgent returns a user agent token describing the AI agent driving this process, or an
// empty string when none is detected. Detection is based on self-reported environment
// variables, so it is suitable for analytics but must not be relied on for security decisions.
//
// Examples:
//
//	AIAgent/ClaudeCode (claude-code_2-1-322_agent)   AI_AGENT set by a known agent
//	AIAgent/Other (devin@1)                          AI_AGENT set by an unrecognised agent
//	AIAgent/Cursor                                   detected from vendor-specific variables
func GetAIAgent(e Environment) string {
	if raw := sanitiseAIAgentValue(e.Getenv(aiAgentEnvVar)); raw != "" {
		name := nameFromAIAgentValue(raw)
		if name == "" {
			// Vendor variables may still identify an agent that sets AI_AGENT in an unfamiliar format.
			name = nameFromVendorVariables(e)
		}
		if name == "" {
			name = "Other"
		}
		return fmt.Sprintf("AIAgent/%s (%s)", name, raw)
	}

	if name := nameFromVendorVariables(e); name != "" {
		return "AIAgent/" + name
	}

	return ""
}

func nameFromVendorVariables(e Environment) string {
	for _, agent := range knownAIAgents {
		for _, rule := range agent.rules {
			value := e.Getenv(rule.variable)
			if value == "" {
				continue
			}
			if rule.value == "" || strings.EqualFold(value, rule.value) {
				return agent.name
			}
		}
	}
	return ""
}

func nameFromAIAgentValue(raw string) string {
	normalised := strings.ReplaceAll(strings.ToLower(raw), "_", "-")
	for _, p := range aiAgentValuePrefixes {
		if strings.HasPrefix(normalised, p.prefix) {
			return p.name
		}
	}
	return ""
}

// sanitiseAIAgentValue keeps AI_AGENT safe to embed in a User-Agent comment: it is free-form and
// user-settable, so restrict it to a conservative character set and bound its length.
func sanitiseAIAgentValue(value string) string {
	var b strings.Builder
	for _, r := range strings.TrimSpace(value) {
		if b.Len() >= maxAIAgentValueLength {
			break
		}
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '.', r == '_', r == '-', r == '@', r == '+':
			b.WriteRune(r)
		}
	}
	return b.String()
}
