package agent

type AgentModel string

const (
	AgentModelGPT3Dot5          AgentModel = "gpt-3.5"
	AgentModelGPT4              AgentModel = "gpt-4"
	AgentModelClaudeSonnet3Dot5 AgentModel = "claude-sonnet-3.5"
	AgentModelDeepSeekR1        AgentModel = "deepseek-reasoner"
	AgentModelGeminiFlash1Dot5  AgentModel = "gemini-1.5-flash"
)

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
		cfg.Model = model
	}
}
