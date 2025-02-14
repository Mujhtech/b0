package agent

import "fmt"

type AgentModel string
type WorkflowType string

const (
	AgentModelGPT3Dot5          AgentModel = "gpt-3.5"
	AgentModelGPT4              AgentModel = "gpt-4"
	AgentModelClaudeSonnet3Dot5 AgentModel = "claude-sonnet-3.5"
	AgentModelDeepSeekR1        AgentModel = "deepseek-reasoner"
	AgentModelGeminiFlash1Dot5  AgentModel = "gemini-1.5-flash"
	AgentModelGeminiFlash2Dot0  AgentModel = "gemini-2.0-flash"
	AgentModelNone              AgentModel = "none"

	WorkflowTypeRequest  WorkflowType = "request"
	WorkflowTypeResponse WorkflowType = "response"
	WorkflowTypeError    WorkflowType = "error"
	WorkflowTypeIf       WorkflowType = "if"
	WorkflowTypeFor      WorkflowType = "for"
	WorkflowTypeWhile    WorkflowType = "while"
	WorkflowTypeSwitch   WorkflowType = "switch"
	WorkflowTypeVariable WorkflowType = "variable"

	WorkflowTypeResend   WorkflowType = "resend"
	WorkflowTypeOpenAI   WorkflowType = "openai"
	WorkflowTypeSlack    WorkflowType = "slack"
	WorkflowTypeDiscord  WorkflowType = "discord"
	WorkflowTypeTelegram WorkflowType = "telegram"
	WorkflowTypeGithub   WorkflowType = "github"
	WorkflowTypeSupabase WorkflowType = "supabase"
	WorkflowTypeStripe   WorkflowType = "stripe"
)

type ModeCatalog struct {
	Name           string     `json:"name"`
	Model          AgentModel `json:"model"`
	IsEnabled      bool       `json:"is_enabled"`
	IsExperimental bool       `json:"is_experimental"`
	IsDefault      bool       `json:"is_default"`
}

type WorkflowCase struct {
	Body  interface{} `json:"body"`
	Value string      `json:"value"`
}

type Workflow struct {
	Type        WorkflowType   `json:"type"`
	Instruction string         `json:"instruction"`
	Then        []Workflow     `json:"then,omitempty"`
	Else        []Workflow     `json:"else,omitempty"`
	Condition   string         `json:"condition,omitempty"`
	Cases       []WorkflowCase `json:"cases,omitempty"`
	Value       interface{}    `json:"value,omitempty"`
	Url         string         `json:"url,omitempty"`
	Method      string         `json:"method,omitempty"`
	Name        string         `json:"name,omitempty"`
	Headers     []string       `json:"headers,omitempty"`
	Body        interface{}    `json:"body,omitempty"`
	Variables   []string       `json:"variables,omitempty"`
	ActionID    string         `json:"action_id,omitempty"`
	Status      string         `json:"status,omitempty"`
}

type CodeGenerationOption struct {
	Language             string      `json:"language"`
	Framework            string      `json:"framework"`
	FrameworkInsructions string      `json:"-"`
	Workflows            interface{} `json:"-"`
	Image                string      `json:"-"`
}

type CodeGeneration struct {
	FileContents    []FileContent `json:"fileContents"`
	InstallCommands []string      `json:"installCommands"`
	RunCommands     string        `json:"runCommands"`
}

type FileContent struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

var AvailableCodeGenerationOptions = []CodeGenerationOption{
	{
		Language:  "Go",
		Framework: "Chi",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- For router, use go-chi/chi/v5
		`,
	},
	{
		Language:  "Go",
		Framework: "Echo",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Echo framework
		`,
	},
	{
		Language:  "Go",
		Framework: "Gin",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Gin framework
		`,
	},
	{
		Language:  "Node.js (TypeScript)",
		Framework: "Express",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Express framework
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		- Use port %s for the server
		`,
	},
	{
		Language:  "Node.js (TypeScript)",
		Framework: "Fastify",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Fastify framework
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		`,
	},
	{
		Language:  "Node.js (TypeScript)",
		Framework: "Hono",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Hono framework
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		- Below are the basic dependencies you need to add to your package.json file
		1. "hono": "^4.7.1"
		- Make sure all code is written in TypeScript
		- Make sure all files are in the src directory except the configuration files like package.json, tsconfig.json, and tsconfig.node.json, etc
		`,
	},
}

func GetLanguageCodeGeneration(language string, framework string) (CodeGenerationOption, error) {
	for _, option := range AvailableCodeGenerationOptions {
		if option.Language == language && option.Framework == framework {
			return option, nil
		}
	}
	return CodeGenerationOption{}, fmt.Errorf("language not found")
}

type AgentToken struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

var AvailableCatalogs = []ModeCatalog{
	{
		Name:           "GPT 3.5",
		Model:          AgentModelGPT3Dot5,
		IsEnabled:      false,
		IsExperimental: false,
		IsDefault:      false,
	},
	{
		Name:           "GPT 4",
		Model:          AgentModelGPT4,
		IsEnabled:      false,
		IsExperimental: false,
		IsDefault:      false,
	},
	{
		Name:           "Claude Sonnet 3.5",
		Model:          AgentModelClaudeSonnet3Dot5,
		IsEnabled:      false,
		IsExperimental: false,
		IsDefault:      false,
	},
	{
		Name:           "DeepSeek R1",
		Model:          AgentModelDeepSeekR1,
		IsEnabled:      false,
		IsExperimental: true,
		IsDefault:      false,
	},
	{
		Name:           "Gemini 1.5 Flash",
		Model:          AgentModelGeminiFlash1Dot5,
		IsEnabled:      true,
		IsExperimental: true,
		IsDefault:      true,
	},
	{
		Name:           "Gemini 2.0 Flash",
		Model:          AgentModelGeminiFlash2Dot0,
		IsEnabled:      true,
		IsExperimental: true,
		IsDefault:      false,
	},
}

func GetModel(model string) (AgentModel, error) {
	for _, catalog := range AvailableCatalogs {
		if catalog.Model == AgentModel(model) {
			return catalog.Model, nil
		}
	}
	return AgentModelNone, fmt.Errorf("model not found")
}

type Config struct {
	Model        AgentModel
	OpenAIKey    string
	DeepSeekKey  string
	AnthropicKey string
	GeminiKey    string
}

type OptionFunc func(*Config)

// WithModel sets the model to be used by the agent
func WithModel(model AgentModel) OptionFunc {
	return func(cfg *Config) {
		if model != AgentModelNone {
			cfg.Model = model
		}
	}
}

func ToModel(model string) AgentModel {
	switch model {
	case string(AgentModelGPT3Dot5):
		return AgentModelGPT3Dot5
	case string(AgentModelGPT4):
		return AgentModelGPT4
	case string(AgentModelClaudeSonnet3Dot5):
		return AgentModelClaudeSonnet3Dot5
	case string(AgentModelDeepSeekR1):
		return AgentModelDeepSeekR1
	case string(AgentModelGeminiFlash1Dot5):
		return AgentModelGeminiFlash1Dot5
	default:
		return AgentModelNone
	}
}
