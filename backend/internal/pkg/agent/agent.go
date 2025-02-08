package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mujhtech/b0/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
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
	request = {"type": "request", "instruction":"...", "method": "...", "url": "...", "body": "..."}
	if = {"type": "if", "instruction":"...", "condition": "...", "then": "...", "else": "..."}
	for = {"type": "for", "instruction":"...", "condition": "...", "body": "..."}
	while = {"type": "while", "instruction":"...", "condition": "...", "body": "..."}
	variable = {"type": "variable", "name": "...", "value": "..."}
	switch = {"type": "switch", "instruction":"...", "condition": "...", "cases": [{"value": "...", "body": "..."}]}
	response = {"type": "response", "instruction":"...", "value": "..."}

	Integration:
	resend = {"type": "resend", "instruction":"...", "url": "...", "method": "...", "body": "..."}
	slack = {"type": "slack", "instruction":"...", "channel": "...", "message": "..."}
	discord = {"type": "discord", "instruction":"...", "channel": "...", "message": "..."}
	telegram = {"type": "telegram", "instruction":"...", "channel": "...", "message": "..."}
	stripe = {"type": "stripe", "instruction":"...", "method": "...", "url": "...", "body": "..."}
	openai = {"type": "openai", "instruction":"...", "model": "...", "prompt": "...", "temperature": "...", "max_tokens": "...", "top_p": "...", "frequency_penalty": "...", "presence_penalty": "..."}
	supabase = {"type": "supabase", "instruction":"...", "table": "...", "method": "...", "body": "..."}
	github = {"type": "github", "instruction":"...", "method": "...", "url": "...", "body": "..."}

	## Requirements:
	- The workflow diagram will be in json format.
	- Workflow can be nested and can have multiple nodes that represent the workflow.
	- Make sure to follow the instructions above
	- Ignore comments in the workflow diagram.

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

func (a *Agent) GenerateTitleAndSlugWithSchema(ctx context.Context, prompt string, opts ...OptionFunc) (*ProjectTitleAndSlug, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
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
		return nil, err
	}

	// extract into a well-typed struct
	dst := ProjectTitleAndSlug{}
	_ = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &dst)

	return &dst, nil
}

// GenerateTitleAndSlug generates a title and slug for a new project based on the given prompt.
func (a *Agent) GenerateTitleAndSlug(ctx context.Context, prompt string, opts ...OptionFunc) (string, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
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
		return "", err
	}

	return chat.Choices[0].Message.Content, nil
}

// GenerateWorkflow generates a workflow diagram based on the given prompt.
func (a *Agent) GenerateWorkflow(ctx context.Context, prompt string, opts ...OptionFunc) (*[]Workflow, string, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
	}

	client := a.client()

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(fmt.Sprintf(b0ProjectWorkflowSystemMessage, a.cfg.Model)),
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModel(a.cfg.Model)),
	})

	if err != nil {
		return nil, "", err
	}

	workflowString := chat.Choices[0].Message.Content

	workflowString = strings.ReplaceAll(workflowString, "json", "")
	workflowString = strings.ReplaceAll(workflowString, "```", "")
	workflowString = strings.ReplaceAll(workflowString, "\n", "")
	workflowString = strings.ReplaceAll(workflowString, `\`, "")

	var rawDogData map[string]interface{}

	err = json.Unmarshal([]byte(workflowString), &rawDogData)

	if err != nil {
		return nil, workflowString, err
	}

	var workflows []Workflow
	workflowsRaw, ok := rawDogData["workflows"].([]interface{})
	if !ok {
		return nil, workflowString, fmt.Errorf("invalid workflows format")
	}

	workflowsJson, err := json.Marshal(workflowsRaw)
	if err != nil {
		return nil, workflowString, fmt.Errorf("failed to marshal workflows: %w", err)
	}

	if err := json.Unmarshal(workflowsJson, &workflows); err != nil {
		return nil, workflowString, fmt.Errorf("failed to unmarshal workflows: %w", err)
	}

	return &workflows, workflowString, nil
}

// CodeGeneration generates code from a workflow diagram.
func (a *Agent) CodeGeneration(ctx context.Context, prompt string, language string, workflow string, opts ...OptionFunc) (string, error) {
	opCfg := *a.cfg
	for _, opt := range opts {
		opt(&opCfg)
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
		return "", err
	}

	return chat.Choices[0].Message.Content, nil
}
