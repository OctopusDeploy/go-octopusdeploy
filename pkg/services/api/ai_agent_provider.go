package api

import (
	"regexp"
	"strings"
)

// AI agents are detected separately from automation (CI) environments because the two are
// independent dimensions: an agent commonly runs inside a CI system or on an Octopus worker,
// and both should be reported rather than one masking the other.

// aiAgentEnvVar is the emerging cross-vendor convention for an agent to identify itself.
// https://github.com/vercel/vercel/tree/main/packages/detect-agent
const aiAgentEnvVar = "AI_AGENT"

// aiAgentSeparators matches runs of underscores and whitespace, which are normalised to a single hyphen.
var aiAgentSeparators = regexp.MustCompile(`[_\s]+`)

type aiAgentRule struct {
	variable string
	// value, when set, must match the variable's value (case-insensitive). When empty, any
	// non-empty value matches.
	value string
}

type aiAgent struct {
	name string
	// aiAgentPrefix, when set, matches the start of a normalised AI_AGENT value. Vendors don't agree
	// on a format (e.g. "claude-code_2-1-322_agent", "github_copilot_vscode_agent", "devin@1"), so
	// values are lowercased and runs of underscores and whitespace replaced with a single hyphen
	// before matching. The prefix must be followed by the end of the value or a separator, so a short
	// prefix such as "amp" doesn't match an unrelated identifier such as "amplitude-mcp".
	aiAgentPrefix string
	// variables are vendor-specific environment variables that identify the agent.
	variables []aiAgentRule
}

// knownAIAgents lists each agent once. An agent can be recognised from AI_AGENT, from its vendor
// variables, or both. AI_AGENT is checked against the whole list before any vendor variables, so a
// self-reported identity wins (e.g. Claude Code running inside Cursor). Within each lookup the first
// match wins, so the result is deterministic when several agents' variables are present.
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
	{name: "ClaudeCode", aiAgentPrefix: "claude-code", variables: []aiAgentRule{{variable: "CLAUDECODE"}, {variable: "CLAUDE_CODE_ENTRYPOINT"}}},
	{name: "Cursor", aiAgentPrefix: "cursor", variables: []aiAgentRule{{variable: "CURSOR_AGENT"}, {variable: "CURSOR_EXTENSION_HOST_ROLE", value: "agent-exec"}}},
	{name: "Codex", aiAgentPrefix: "codex", variables: []aiAgentRule{{variable: "CODEX_SANDBOX"}, {variable: "CODEX_THREAD_ID"}, {variable: "CODEX_CI"}}},
	{name: "GeminiCLI", aiAgentPrefix: "gemini", variables: []aiAgentRule{{variable: "GEMINI_CLI"}}},
	{name: "GitHubCopilot", aiAgentPrefix: "github-copilot", variables: []aiAgentRule{{variable: "COPILOT_CLI"}, {variable: "COPILOT_AGENT"}, {variable: "GITHUB_COPILOT_CLI_MODE"}}},
	{name: "OpenCode", aiAgentPrefix: "opencode", variables: []aiAgentRule{{variable: "OPENCODE", value: "1"}, {variable: "OPENCODE_CLIENT"}, {variable: "OPENCODE_AI"}}},
	{name: "Antigravity", aiAgentPrefix: "antigravity", variables: []aiAgentRule{{variable: "ANTIGRAVITY_AGENT"}}},
	{name: "Goose", aiAgentPrefix: "goose", variables: []aiAgentRule{{variable: "GOOSE_TERMINAL"}, {variable: "AGENT", value: "goose"}}},
	{name: "Amp", aiAgentPrefix: "amp", variables: []aiAgentRule{{variable: "AGENT", value: "amp"}}},
	{name: "Cline", aiAgentPrefix: "cline", variables: []aiAgentRule{{variable: "CLINE_ACTIVE"}, {variable: "CLINE_TASK_ID"}}},
	{name: "RooCode", variables: []aiAgentRule{{variable: "ROO_CODE_TASK_ID"}}},
	{name: "Augment", aiAgentPrefix: "augment", variables: []aiAgentRule{{variable: "AUGMENT_AGENT"}}},
	{name: "QwenCode", variables: []aiAgentRule{{variable: "QWEN_CODE"}}},
	{name: "FactoryDroid", variables: []aiAgentRule{{variable: "DROID_CLI"}}},
	{name: "KimiCLI", variables: []aiAgentRule{{variable: "KIMI_CLI"}}},
	{name: "Trae", variables: []aiAgentRule{{variable: "TRAE_AI_SHELL_ID"}}},
	{name: "Aider", variables: []aiAgentRule{{variable: "OR_APP_NAME", value: "Aider"}}},
	{name: "OpenHands", variables: []aiAgentRule{{variable: "OR_APP_NAME", value: "OpenHands"}}},
	{name: "Devin", aiAgentPrefix: "devin"},
	{name: "Replit", aiAgentPrefix: "replit"},
	{name: "V0", aiAgentPrefix: "v0"},
}

// GetAIAgent returns a User-Agent comment naming the AI agent driving this process, or an empty
// string when none is detected. Only a normalised name is reported: the raw AI_AGENT value is
// free-form and user-settable, so it is used for matching and never included in the output.
// Detection is based on self-reported environment variables, so it is suitable for analytics but
// must not be relied on for security decisions.
//
// Examples:
//
//	(ai-agent=ClaudeCode)   AI_AGENT=claude-code_2-1-322_agent, or CLAUDECODE=1
//	(ai-agent=Cursor)       CURSOR_AGENT=1
//	(ai-agent=Other)        AI_AGENT set to a value that isn't recognised
func GetAIAgent(e Environment) string {
	aiAgentNameFromEnvVar := strings.TrimSpace(e.Getenv(aiAgentEnvVar))

	// Vendor variables also cover an agent that sets AI_AGENT in a format that isn't recognised.
	if agent := findAIAgentByPrefix(aiAgentNameFromEnvVar); agent != nil {
		return aiAgentComment(agent.name)
	}
	if agent := findAIAgentByVariables(e); agent != nil {
		return aiAgentComment(agent.name)
	}
	if aiAgentNameFromEnvVar != "" {
		return aiAgentComment("Other")
	}

	return ""
}

// aiAgentComment formats the agent as an RFC 9110 comment, so it describes the request rather than
// claiming to be a product in the User-Agent.
func aiAgentComment(name string) string {
	return "(ai-agent=" + name + ")"
}

// appendAIAgent adds the agent comment to the end of an existing User-Agent, so everything before it
// is unchanged and existing parsers that match on the start of the header keep working.
func appendAIAgent(userAgent string, e Environment) string {
	if aiAgent := GetAIAgent(e); aiAgent != "" {
		return userAgent + " " + aiAgent
	}
	return userAgent
}

func findAIAgentByPrefix(value string) *aiAgent {
	if value == "" {
		return nil
	}
	normalised := aiAgentSeparators.ReplaceAllString(strings.ToLower(value), "-")
	for i, agent := range knownAIAgents {
		if agent.aiAgentPrefix != "" && matchesAIAgentPrefix(normalised, agent.aiAgentPrefix) {
			return &knownAIAgents[i]
		}
	}
	return nil
}

// matchesAIAgentPrefix reports whether value starts with prefix followed by the end of the value, or by
// a separator ("-", "@", ".") or digit.
func matchesAIAgentPrefix(value, prefix string) bool {
	if !strings.HasPrefix(value, prefix) {
		return false
	}
	if len(value) == len(prefix) {
		return true
	}
	next := value[len(prefix)]
	return next == '-' || next == '@' || next == '.' || (next >= '0' && next <= '9')
}

func findAIAgentByVariables(e Environment) *aiAgent {
	for i, agent := range knownAIAgents {
		for _, rule := range agent.variables {
			value := e.Getenv(rule.variable)
			if value == "" {
				continue
			}
			if rule.value == "" || strings.EqualFold(value, rule.value) {
				return &knownAIAgents[i]
			}
		}
	}
	return nil
}
