package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mujhtech/b0/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/rs/zerolog"
)

const (
	deepSeekBaseUrl = "https://api.deepseek.com/v1"
	openaiBaseUrl   = "https://api.openai.com/v1"
	geminiBaseUrl   = "https://generativelanguage.googleapis.com/v1beta/openai/"

	b0DefaultSystemMessage = `You are b0, an AI assitant for building backend service powered by %s model, created by mujhtech.xyz.`

	b0ProjectTitleAndSlugSystemMessage = b0DefaultSystemMessage + `
	You are here to help user generate project title, description and slug for a new project based on user prompt. 
	
	Take note, you will respond in valid JSON format only. 
	Example schema: {"title": "...", "description": "...", "slug": "..."} 

	And don't forget that the slug should be in lowercase and it should not be more than 3 words separated with dashes and the last word should be a unique identifier of six letter. The title should be a short and concise description of the project.
	`

	b0ProjectWorkflowSystemMessage = b0DefaultSystemMessage + `
	You are here to help user generate a workflow diagram node based on the user prompt. The workflow diagram node will be in json format and you are to generate the workflow diagram node based on the prompt.
	
	Example of workflow template are: if, for, while,
	request = {"action_id": "...",  "type": "request", "name", "...", "instruction":"...", "method": "POST" | "PUT" | "GET" | "DELETE | "PATCH", "url": "...", "body": "..."}
	if = {"action_id": "...", "type": "if", "instruction":"...", "condition": "...", "then": "...", "else": "..."}
	for = {"action_id": "...", "type": "for", "instruction":"...", "condition": "...", "body": "..."}
	while = {"action_id": "...", "type": "while", "instruction":"...", "condition": "...", "body": "..."}
	variable = {"action_id": "...", "type": "variable", "name": "...", "value": "..."}
	switch = {"action_id": "...", "type": "switch", "instruction":"...", "condition": "...", "cases": [{"value": "...", "body": "..."}]}
	response = {"action_id": "...", "type": "response", "instruction":"...", "status": "...", "body": "..."}

	Integration:
	resend = {"action_id": "...", "type": "resend", "instruction":"...", "url": "...", "method": "...", "body": "..."}
	slack = {"action_id": "...", "type": "slack", "instruction":"...", "channel": "...", "message": "..."}
	discord = {"action_id": "...", "type": "discord", "instruction":"...", "channel": "...", "message": "..."}
	telegram = {"action_id": "...", "type": "telegram", "instruction":"...", "channel": "...", "message": "..."}
	stripe = {"action_id": "...", "type": "stripe", "instruction":"...", "method": "...", "url": "...", "body": "..."}
	openai = {"action_id": "...", "type": "openai", "instruction":"...", "model": "...", "prompt": "...", "temperature": "...", "max_tokens": "...", "top_p": "...", "frequency_penalty": "...", "presence_penalty": "..."}
	supabase = {"action_id": "...", "type": "supabase", "instruction":"...", "table": "...", "method": "...", "body": "..."}
	github = {"action_id": "...", "type": "github", "instruction":"...", "method": "...", "url": "...", "body": "..."}

	## Requirements:
	- The workflow diagram will be in json format.
	- Workflow must start with a request node.
	- Workflow can be nested and can have multiple nodes that represent the workflow.
	- Make sure to follow the instructions above
	- Ignore comments in the workflow diagram.
	- action_id must be unique.
	- Use context to store and access data between nodes e.g {{context.request}}, {{context.request.body}}, {{context.response}}, {{context.response.body}},  {{context.variable_name}}
	- Make sure that http response status code is string and not int.
	- The url in the request node must be a path to the endpoint not external url.

	## Output:
	- The output should be a json string in the format of {"workflows": ["..."]}
	- For string interpolation, use {{...}} for the value.
	`

	b0WorkflowToCodeGenerationSystemMessage = b0DefaultSystemMessage + `You are here to help user generate code from a workflow diagram. The workflow diagram will be in json format and you are to generate the code based on the diagram.
	
	
	## Requirements:
	- The workflow diagram will be in json format.
	- Use %s for code generation.
	- For the code generation, you are to generate the code based on the workflow diagram.
	- Make sure to follow the workflow diagram below

	## Workflow Diagram:
	- The workflow diagram will be in json format.
	

	%s

	## Output:
	- The output should be a json string in the format of {"code": "..."}
	- Ignore comments in the workflow diagram.
	`
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
	case AgentModelGeminiFlash1Dot5:
		baseUrl = geminiBaseUrl
		apiKey = a.cfg.GeminiKey
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
		Input: fmt.Sprintf("%s%s", fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model), prompt),
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
		Input: fmt.Sprintf("%s%s", fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model), prompt),
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
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "Generated title and slug:", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "json", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "```", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, "\n", "")
	projectTitleAndSlug = strings.ReplaceAll(projectTitleAndSlug, `\`, "")

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
func (a *Agent) GenerateWorkflow(ctx context.Context, prompt string, opts ...OptionFunc) ([]*Workflow, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	client := a.client()

	agentToken := &AgentToken{
		Input: fmt.Sprintf("%s%s", fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model), prompt),
	}

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model)),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return nil, agentToken, err
	}

	workflowString := chat.Choices[0].Message.Content

	zerolog.Ctx(ctx).Info().Msgf("Generated workflows: %s", workflowString)

	agentToken.Output = workflowString

	workflowString = strings.ReplaceAll(workflowString, "json", "")
	workflowString = strings.ReplaceAll(workflowString, "```", "")
	workflowString = strings.ReplaceAll(workflowString, "\n", "")
	workflowString = strings.ReplaceAll(workflowString, `\`, "")

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
func (a *Agent) CodeGeneration(ctx context.Context, prompt string, language string, workflow string, opts ...OptionFunc) (string, *AgentToken, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	agentToken := &AgentToken{
		Input: fmt.Sprintf("%s%s", fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model), prompt),
	}

	client := a.client()

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0WorkflowToCodeGenerationSystemMessage, a.cfg.Model, language, workflow)),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return "", agentToken, err
	}

	agentToken.Output = chat.Choices[0].Message.Content

	return chat.Choices[0].Message.Content, agentToken, nil
}
