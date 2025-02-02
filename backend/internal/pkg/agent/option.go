package agent

type AgentModel string

type ModeCatalog struct {
	Name           string     `json:"name"`
	Model          AgentModel `json:"model"`
	IsEnabled      bool       `json:"is_enabled"`
	IsExperimental bool       `json:"is_experimental"`
	IsDefault      bool       `json:"is_default"`
}

const (
	AgentModelGPT3Dot5          AgentModel = "gpt-3.5"
	AgentModelGPT4              AgentModel = "gpt-4"
	AgentModelClaudeSonnet3Dot5 AgentModel = "claude-sonnet-3.5"
	AgentModelDeepSeekR1        AgentModel = "deepseek-reasoner"
	AgentModelGeminiFlash1Dot5  AgentModel = "gemini-1.5-flash"
	AgentModelNone              AgentModel = "none"
)

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
