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
