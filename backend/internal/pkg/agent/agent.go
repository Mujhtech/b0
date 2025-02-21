package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/internal/util"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/rs/zerolog"
)

const (
	deepSeekBaseUrl = "https://api.deepseek.com/v1"
	openaiBaseUrl   = "https://api.openai.com/v1"
	geminiBaseUrl   = "https://generativelanguage.googleapis.com/v1beta/openai/"
	xAIbaseUrl      = "https://api.x.ai/v1"
)

type Agent struct {
	cfg *Config
}

func New(cfg *config.Config) *Agent {

	return &Agent{
		cfg: &Config{
			Model:        AgentModelGeminiFlash1Dot5,
			OpenAIKey:    cfg.Agent.OpenAIKey,
			DeepSeekKey:  cfg.Agent.DeepSeekKey,
			AnthropicKey: cfg.Agent.AnthropicKey,
			GeminiKey:    cfg.Agent.GeminiKey,
			XAIKey:       cfg.Agent.XAIKey,
		},
	}
}

func (a *Agent) client(opts ...option.RequestOption) *openai.Client {

	baseUrl := ""
	apiKey := ""
	switch a.cfg.Model {
	case AgentModelGPT4, AgentModelGPT3Dot5:
		baseUrl = openaiBaseUrl
		apiKey = a.cfg.OpenAIKey
	case AgentModelDeepSeekR1:
		baseUrl = deepSeekBaseUrl
		apiKey = a.cfg.DeepSeekKey
	case AgentModelClaudeSonnet3Dot5:
		baseUrl = openaiBaseUrl
		apiKey = a.cfg.AnthropicKey
	case AgentModelGeminiFlash1Dot5, AgentModelGeminiFlash2Dot0:
		baseUrl = geminiBaseUrl
		apiKey = a.cfg.GeminiKey
	case AgentModelGrok2Dot0:
		baseUrl = xAIbaseUrl
		apiKey = a.cfg.XAIKey
	}

	opts = append(opts, option.WithBaseURL(baseUrl), option.WithAPIKey(apiKey))

	client := openai.NewClient(opts...)

	return client
}

func (a *Agent) GenerateTitleAndSlugWithSchema(ctx context.Context, prompt string, opts ...OptionFunc) (*ProjectTitleAndSlug, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	agentToken := &AgentToken{
		Input: fmt.Sprintf(`
		%s
		
		%s
		`, fmt.Sprintf(b0ProjectTitleAndSlugSystemMessage, a.cfg.Model), prompt),
	}

	client := a.client()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("project"),
		Description: openai.F("Project title, description and slug"),
		Schema:      openai.F(ProjectTitleAndSlugResponseSchema),
		Strict:      openai.Bool(true),
	}

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0ProjectTitleAndSlugSystemMessage, a.cfg.Model)),
			openai.UserMessage(prompt),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		Model: openai.F(openai.ChatModelGPT4o2024_08_06),
	})

	if err != nil {
		return nil, agentToken, err
	}

	// extract into a well-typed struct
	dst := ProjectTitleAndSlug{}
	_ = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &dst)

	return &dst, agentToken, nil
}

// GenerateTitleAndSlug generates a title and slug for a new project based on the given prompt.
func (a *Agent) GenerateTitleAndSlug(ctx context.Context, prompt string, opts ...OptionFunc) (*ProjectTitleAndSlug, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	agentToken := &AgentToken{
		Input: fmt.Sprintf(`
		%s
		
		%s
		`, fmt.Sprintf(b0ProjectTitleAndSlugSystemMessage, a.cfg.Model), prompt),
	}

	client := a.client()

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0ProjectTitleAndSlugSystemMessage, a.cfg.Model)),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return nil, agentToken, err
	}

	projectTitleAndSlug := chat.Choices[0].Message.Content

	agentToken.Output = chat.Choices[0].Message.Content

	zerolog.Ctx(ctx).Info().Msgf("Generated title and slug: %s", projectTitleAndSlug)

	if strings.HasPrefix(projectTitleAndSlug, "I'm b0, an AI assistant") {
		return nil, agentToken, fmt.Errorf("invalid prompt")
	}

	// cleanup the message
	projectTitleAndSlug = removeJSONMarkdown(projectTitleAndSlug)

	zerolog.Ctx(ctx).Info().Msgf("cleanup response: %s", strings.TrimSpace(projectTitleAndSlug))

	// unmarshal the projectTitleAndSlug
	var agentProjectTitleAndSlug *ProjectTitleAndSlug

	if err := json.Unmarshal([]byte(projectTitleAndSlug), &agentProjectTitleAndSlug); err != nil {
		return nil, agentToken, err
	}

	zerolog.Ctx(ctx).Info().Msgf("%s", agentProjectTitleAndSlug)

	return agentProjectTitleAndSlug, agentToken, nil
}

// GenerateWorkflow generates a workflow diagram based on the given prompt.
func (a *Agent) GenerateWorkflow(ctx context.Context, options WorkflowGenerationOption, opts ...OptionFunc) ([]*Workflow, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	client := a.client()

	otherInstructions := ""

	if len(options.Workflows) > 0 {

		workflowToString, _ := util.MarshalJSONToString(options.Workflows)

		otherInstructions = workflowToString
	}

	systemPrompt := fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model, otherInstructions)

	if len(options.Workflows) > 0 {
		systemPrompt = fmt.Sprintf(b0UpdateProjectWorkflowSystemMessage, a.cfg.Model, otherInstructions)
	}

	agentToken := &AgentToken{
		Input: fmt.Sprintf(`
		%s
		
		%s
		`, systemPrompt, options.Prompt),
	}

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPrompt),
			openai.UserMessage(options.Prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return nil, agentToken, err
	}

	workflowString := chat.Choices[0].Message.Content

	zerolog.Ctx(ctx).Info().Msgf("Generated workflows: %s", workflowString)

	agentToken.Output = workflowString

	workflowString = removeJSONMarkdown(workflowString)

	var rawDogData map[string]interface{}

	err = json.Unmarshal([]byte(workflowString), &rawDogData)

	if err != nil {
		return nil, agentToken, err
	}

	var workflows []*Workflow
	workflowsRaw, ok := rawDogData["workflows"].([]interface{})
	if !ok {
		return nil, agentToken, fmt.Errorf("invalid workflows format")
	}

	workflowsJson, err := json.Marshal(workflowsRaw)
	if err != nil {
		return nil, agentToken, fmt.Errorf("failed to marshal workflows: %w", err)
	}

	if err := json.Unmarshal(workflowsJson, &workflows); err != nil {
		return nil, agentToken, fmt.Errorf("failed to unmarshal workflows: %w", err)
	}

	return workflows, agentToken, nil
}

// CodeGeneration generates code from a workflow diagram.
func (a *Agent) CodeGeneration(ctx context.Context, prompt string, option CodeGenerationOption, opts ...OptionFunc) (*CodeGeneration, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	workflowToString, err := util.MarshalJSONToString(option.Workflows)

	if err != nil {
		return nil, nil, err
	}

	agentToken := &AgentToken{
		Input: fmt.Sprintf(`
		%s
		
		%s
		`, fmt.Sprintf(b0WorkflowToCodeGenerationSystemMessage, a.cfg.Model, option.Language, option.FrameworkInsructions, workflowToString), prompt),
	}

	client := a.client()

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0WorkflowToCodeGenerationSystemMessage, a.cfg.Model, option.Language, option.FrameworkInsructions, workflowToString)),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return nil, agentToken, err
	}

	agentToken.Output = chat.Choices[0].Message.Content

	zerolog.Ctx(ctx).Info().Msgf("Generated code: %s", chat.Choices[0].Message.Content)

	var codeGeneration *CodeGeneration

	if err := json.Unmarshal([]byte(removeJSONMarkdown(chat.Choices[0].Message.Content)), &codeGeneration); err != nil {
		return nil, agentToken, err
	}

	return codeGeneration, agentToken, nil
}
