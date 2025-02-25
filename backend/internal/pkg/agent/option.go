package agent

import "fmt"

type AgentModel string
type WorkflowType string

const (
	AgentModelGPT3Dot5          AgentModel = "gpt-3.5"
	AgentModelGPT4              AgentModel = "gpt-4"
	AgentModelClaudeSonnet3Dot5 AgentModel = "claude-sonnet-3.5"
	AgentModelClaudeSonnet3Dot7 AgentModel = "claude-sonnet-3.7"
	AgentModelDeepSeekR1        AgentModel = "deepseek-reasoner"
	AgentModelGeminiFlash1Dot5  AgentModel = "gemini-1.5-flash"
	AgentModelGeminiFlash2Dot0  AgentModel = "gemini-2.0-flash"
	AgentModelGrok2Dot0         AgentModel = "grok-2-latest"
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

	nodeJSInstructions = `
	## Integration Instructions
	- Github
	Use the following package to interact with the Github API:

	dependencies:
	octokit
	@octokit/request-error

	devDependencies:
	@octokit/types

	e.g To create a new issue:

	import { Octokit, App } from "octokit";
	import { RequestError } from "@octokit/request-error";
	import { RequestRequestOptions } from "@octokit/types";

	const client = new Octokit({
		auth: process.env.B0_GITHUB_TOKEN,
		retry: {
        	enabled: false,
        },
	});


	const createIssue = async () => {
		try {
			const res = await client.request("POST /repos/{owner}/{repo}/issues", {
				owner: process.env.B0_GITHUB_OWNER,
				repo: process.env.B0_GITHUB_REPO,
				title: "Found a bug",
				body: "I'm having a problem with this.",
			});
			console.log(res);
		} catch (error) {
			if (error instanceof RequestError) {
				console.error(error.message);
			}
		}
	};

	- Slack
	Use the following package to interact with the Slack API:

	dependencies:
	@slack/web-api

	e.g To send a message to a channel:

	import { WebClient } from "@slack/web-api";

	const web = new WebClient(process.env.B0_SLACK_KEY);

	(async () => {
		const res = await web.chat.postMessage({
			channel: process.env.B0_SLACK_CHANNEL_ID,
			text: "Hello, world!",
		});
		console.log(res);
	})();

	- Discord
	Use the following package to interact with the Discord API:
	dependencies:
	@discordjs/rest: "^2.4.3"
	discord-api-types: "^0.37.119"

	e.g To send a message to a channel:

	import { REST } from "@discordjs/rest";

	const rest = new REST({ version: "10" }).setToken(process.env.B0_DISCORD_KEY!);

	await rest.post(Routes.channelMessages(process.env.B0_DISCORD_CHANNEL_ID!), { {
		body: {
			content: "Hello, world!",
		},
	});

	- Telegram
	Use fetch to interact with the Telegram BOT API, e.g to send message

	fetch("https://api.telegram.org/bot<process.env.B0_TELEGRAM_BOT_TOKEN>/sendMessage", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			chat_id: "<chat_id>",
			text: "Hello, world!",
		}),
	});

	- Resend
	Use the following package to interact with the Resend API:
	dependencies:
	resend: "^4.1.2"

	e.g To send an email:
	import { Resend } from "resend";

	const resend = new Resend(process.env.B0_RESEND_KEY);

	await resend.emails.send({
      from: process.env.B0_RESEND_FROM,
      to: process.env.B0_RESEND_TO,
      subject: 'Hello World',
      html: '<strong>It works!</strong>'
    });

	- Stripe
	Use the following package to interact with the Stripe API:
	dependencies:
	stripe: "^17.6.0"

	e.g To create a new customer:
	import Stripe from "stripe";

	const stripe = new Stripe(process.env.B0_STRIPE_KEY, {
		apiVersion: "2023-08-16",
	});

	const customer = await stripe.customers.create({
		email: "jenny.rosen@example.com",
	});

	- Supabase
	Use the following package to interact with the Supabase API:
	dependencies:
	@supabase/supabase-js: "^2.48.1"

	e.g To interact with the Supabase API:
	import { createClient } from '@supabase/supabase-js';

	const supabase = createClient(process.env.B0_SUPABASE_URL, process.env.B0_SUPABASE_KEY);

	const { data, error } = await supabase.from('users').select('*');

	- OpenAI, Gemini, Anthropic, DeepSeek
	Use the following package to interact with the OpenAI API:
	dependencies:
	ai: "^4.1.41"

	// For openai
	@ai-sdk/openai: "^1.1.12"

	// For anthropic
	@ai-sdk/anthropic: "^1.1.8"

	// For Gemini/Google
	@ai-sdk/google: "^1.1.14"

	// DeepSeek
	@ai-sdk/deepseek: "^0.1.10"

	E.g To interact with the OpenAI API:
	import { generateText } from 'ai';
	import { createOpenAI } from '@ai-sdk/openai';

	const openai = createOpenAI({
		apiKey: process.env.B0_OPENAI_KEY,
		compatibility: 'strict',
	});

	const { text } = await generateText({
		model: openai('gpt-4o'),
		system: 'You are a friendly assistant!',
		prompt: 'Why is the sky blue?',
	});

	E.g To interact with the Anthropic API:
	import { createAnthropic } from '@ai-sdk/anthropic';
	import { generateText } from 'ai';

	const anthropic = createAnthropic({
		apiKey: process.env.B0_ANTHROPIC_KEY,
	});

	const { text } = await generateText({
		model: anthropic('claude-3-haiku-20240307'),
		prompt: 'Write a vegetarian lasagna recipe for 4 people.',
	});

	E.g To interact with the Gemini/Google API:
	import { createGoogleGenerativeAI } from '@ai-sdk/google';
	import { generateText } from 'ai';

	const google = createGoogleGenerativeAI({
		apiKey: process.env.B0_GEMINI_KEY,
	});

	const { text } = await generateText({
		model: google('gemini-1.5-flash'),
		prompt: 'Write a vegetarian lasagna recipe for 4 people.',
	})

	E.g To interact with the DeepSeek API:
	import { createDeepSeek } from '@ai-sdk/deepseek';
	import { generateText } from 'ai';

	const deepseek = createDeepSeek({
		apiKey: process.env.B0_DEEPSEEK_KEY,
	});

	const { text } = await generateText({
		model: deepseek('deepseek-chat'),
		prompt: 'Write a vegetarian lasagna recipe for 4 people.',
	});

	Using any of the provider above, beyond prompt, you can also use the following parameters e.g

	const { text } = await generateText({
		model: openai('gpt-4o'),
		messages: [
			{
				role: 'system',
				text: 'You are a friendly assistant!',
			},
			{
				role: 'user',
				text: 'Why is the sky blue?',
			},
		]
	});

	To use any of the mentioned sdk above make sure you install the ai package mentioned above.

	While working with Node cron, you can use the following package:
	dependencies:
	node-cron: "^3.0.3"

	devDependencies:
	@types/node-cron: "^3.0.11"

	For type checking, when using process environment variables in your code, make sure to add ! to the variable to avoid type checking error e.g process.env.B0_DISCORD_KEY! instead of process.env.B0_DISCORD_KEY excluding the B0_PORT variable.
	`

	goInstructions = `
	## Integration Instructions
	- Resend
	Use github.com/resend/resend-go/v2 to interact with the Resend API:

	E.g To send an email:
	import (
		"context"
		"fmt"
		"os"
		"github.com/resend/resend-go/v2"
	)

	ctx := context.TODO()
	apiKey := os.Getenv("B0_RESEND_KEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		To:      []string{"delivered@resend.dev"},
		From:    "onboarding@resend.dev",
		Text:    "hello world",
		Subject: "Hello from Golang",
		Cc:      []string{"cc@example.com"},
		Bcc:     []string{"ccc@example.com"},
		ReplyTo: "to@example.com",
	}

	sent, err := client.Emails.SendWithContext(ctx, params)
	if err != nil {
		panic(err)
	}
	fmt.Println(sent.Id)

	- Slack
	Use github.com/slack-go/slack to interact with the Slack API:

	E.g To send a message to a channel:

	import (
		"context"
		"fmt"
		"os"
		"github.com/slack-go/slack"
	)

	ctx := context.TODO()
	token  := os.Getenv("B0_SLACK_KEY")
	client := slack.New(token)

	_, _, err := client.PostMessageContext(ctx, "CHANNEL_ID", slack.MsgOptionText("Hello, world!", false))


	- Stripe
	Use github.com/stripe/stripe-go/v81 to interact with the Stripe API:
	E.g To create a new customer:
	import (
		"context"
		"fmt"
		"os"
		"github.com/stripe/stripe-go/v81"
		"github.com/stripe/stripe-go/v81/client"
	)

	apiKey = os.Getenv("B0_STRIPE_KEY")
	
	stripe := &client.API{}
	stripe.Init(apiKey, nil)

	params := &stripe.CustomerParams{
		Email: stripe.String("jenny.rosen@example.com"),
	}

	customer, err := stripe.Customers.New(params)

	if err != nil {
		panic(err)
	}
	fmt.Println(customer.ID)

	- Supabase
	- Telegram
	- Discord
	- Github
	
	- OpenAI, Gemini, Anthropic, DeepSeek
	Use github.com/openai/openai-go to interact with the OpenAI API:

	E.g To interact with the OpenAI API:
	import (
		"context"
		"fmt"
		"os"
		"github.com/openai/openai-go"
		"github.com/openai/openai-go/option"
	)

	provider := "openai"

	deepSeekBaseUrl = "https://api.deepseek.com/v1"
	openaiBaseUrl   = "https://api.openai.com/v1"
	geminiBaseUrl   = "https://generativelanguage.googleapis.com/v1beta/openai/"

	baseUrl := ""

	if provider == "openai" {
		baseUrl = openaiBaseUrl
	} else if provider == "deepseek" {
		baseUrl = deepSeekBaseUrl
	} else if provider == "gemini" {
		baseUrl = geminiBaseUrl
	}

	apiKey := os.Getenv("B0_OPENAI_KEY")

	client := openai.NewClient(option.WithBaseURL(baseUrl), option.WithAPIKey(apiKey))

	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage("You are a friendly assistant!"),
			openai.UserMessage("Why is the sky blue?"),
		}),
		Model: openai.F("gpt-4"),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(chat.Choices[0].Message.Content)
	`
)

type ModeCatalog struct {
	Name           string     `json:"name"`
	Model          AgentModel `json:"model"`
	IsEnabled      bool       `json:"is_enabled"`
	IsExperimental bool       `json:"is_experimental"`
	IsDefault      bool       `json:"is_default"`
	IsPremium      bool       `json:"is_premium"`
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
	Model       string         `json:"model,omitempty"`
	Provider    string         `json:"provider,omitempty"`
	Prompt      string         `json:"prompt,omitempty"`
	ActionID    string         `json:"action_id,omitempty"`
	Status      string         `json:"status,omitempty"`
}

type WorkflowGenerationOption struct {
	Workflows []*Workflow `json:"workflows"`
	Prompt    string      `json:"prompt"`
}

type CodeGenerationOption struct {
	ID                   string      `json:"id"`
	Language             string      `json:"language"`
	Framework            string      `json:"framework"`
	FrameworkInsructions string      `json:"-"`
	Workflows            interface{} `json:"-"`
	Image                string      `json:"-"`
}

type CodeGenEnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CodeGeneration struct {
	FileContents    []FileContent   `json:"fileContents"`
	InstallCommands []string        `json:"installCommands"`
	BuildCommands   string          `json:"buildCommands"`
	RunCommands     string          `json:"runCommands"`
	EnvVars         []CodeGenEnvVar `json:"envVars"`
}

type FileContent struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

var AvailableCodeGenerationOptions = []CodeGenerationOption{
	{
		ID:        "1",
		Language:  "Go",
		Framework: "Chi",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- For router, use go-chi/chi/v5
		- Access the server port from OS environment variable B0_PORT


		` + goInstructions,
	},
	{
		ID:        "2",
		Language:  "Go",
		Framework: "Echo",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Echo framework
		- Access the server port from OS environment variable B0_PORT

		` + goInstructions,
	},
	{
		ID:        "3",
		Language:  "Go",
		Framework: "Gin",
		Image:     "golang:1.23-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Gin framework
		- Access the server port from OS environment variable B0_PORT

		` + goInstructions,
	},
	{
		ID:        "4",
		Language:  "Node.js (TypeScript)",
		Framework: "Express",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Express framework
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		- Use process.env.B0_PORT for the server port
		- Use npm run build for buildCommands
		- Use npm run start for runCommands
		- Make sure to add the necessary scripts in the package.json file e.g "build": "npx -y tsc", "start": "node dist/index.js"
		- For all api key, use this format: process.env.B0_API_KEY e.g process.env.B0_OPENAI_KEY, process.env.B0_SLACK_KEY, etc


		` + nodeJSInstructions,
	},
	{
		ID:        "5",
		Language:  "Node.js (TypeScript)",
		Framework: "Fastify",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Fastify framework
		- Use process.env.B0_PORT for the server port
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		- For all api key, use this format: process.env.B0_API_KEY e.g process.env.B0_OPENAI_KEY, process.env.B0_SLACK_KEY, etc


		` + nodeJSInstructions,
	},
	{
		ID:        "6",
		Language:  "Node.js (TypeScript)",
		Framework: "Hono",
		Image:     "node:20-alpine3.20",
		FrameworkInsructions: `
		## Framework instructions
		- Use Hono framework
		- Use process.env.B0_PORT for the server port
		- Make sure to add package.json and don't forget to include all neccessary dependencies
		- Below are the basic dependencies you need to add to your package.json file
		1. "hono": "^4.7.1"
		- Make sure all code is written in TypeScript
		- Make sure all files are in the src directory except the configuration files like package.json, tsconfig.json, and tsconfig.node.json, etc
		- For all api key, use this format: process.env.B0_API_KEY e.g process.env.B0_OPENAI_KEY, process.env.B0_SLACK_KEY, etc

		Example of simple Hono web api:

		import { Hono } from 'hono';

		const app = new Hono();

		app.get('/', (c) => c.text('Pretty Blog API'))

		export default app;

		To bind request body, here is example of how to bind request body:

		type Bindings = {
			USERNAME: string
			PASSWORD: string
		}

		const api = new Hono<{ Bindings: Bindings }>()

		api.post(
			'/posts',
			async (c, next) => {
				const { USERNAME, PASSWORD } = c.env
				// do something with USERNAME and PASSWORD
			},
			async (c) => {
				const post = await c.req.json<Post>()
				const ok = createPost({ post })
				return c.json({ ok })
			}
		)

		app.route('/api', api)

		To return json response, here is example of how to return json response:

		import { prettyJSON } from 'hono/pretty-json'

		app.use(prettyJSON())

		app.get('/', (c) => c.json({ message: 'Hello', ok: true }, 200))

		To handle 404 error, here is example of how to handle 404 error:

		app.notFound((c) => c.json({ message: 'Not Found', ok: false }, 404))

		To handle cors, here is example of how to handle cors:

		import { cors } from 'hono/cors'

		app.use('*', cors())

		For basic auth implementation, here is example of how to implement basic auth:

		import { basicAuth } from 'hono/basic-auth'

		For request body validation, here is example of how to validate request body:
		depedencies:
		@hono/zod-validator: "^0.4.3"
		zod: "^3.24.2"


		import { z } from 'zod'
		import { zValidator } from '@hono/zod-validator'

		const schema = z.object({
			name: z.string(),
			age: z.number(),
		})

		app.post('/author', zValidator('json', schema), (c) => {
			const data = c.req.valid('json')
			return c.json({
				success: true,
				message: data.name+' is '+data.age+' years old',
			})
		})

		For cache control, here is example of how to implement cache control:
		import { cache } from 'hono/cache'

		app.get(
			'*',
			cache({
				cacheName: 'my-app',
				cacheControl: 'max-age=3600',
			})
		)

		For compression, here is example of how to implement compression:
		import { compress } from 'hono/compress'

		app.use(compress())

		For timeout, here is example of how to implement timeout:
		import { timeout } from 'hono/timeout'

		app.use('/api', timeout(5000))

		For request id, here is example of how to implement request id:
		import { requestId } from 'hono/request-id'

		app.use('*', requestId())

		app.get('/', (c) => {
			return c.text('Your request id is ' + c.get('requestId'))
		})

		For logging, here is example of how to implement logging:
		import { logger } from 'hono/logger'

		app.use(logger())

		For jwt,, here is example of how to implement jwt:

		import { jwt } from 'hono/jwt'
		import type { JwtVariables } from 'hono/jwt'

		type Variables = JwtVariables

		const app = new Hono<{ Variables: Variables }>()

		app.use(
			'/auth/*',
			jwt({
				secret: 'it-is-very-secret',
			})
		)

		app.get('/auth/page', (c) => {
			return c.text('You are authorized')
		})

		app.get('/auth/page', (c) => {
			const payload = c.get('jwtPayload')
			return c.json(payload) // eg: { "sub": "1234567890", "name": "John Doe", "iat": 1516239022 }
		})
		` + nodeJSInstructions,
	},
}

func GetLanguageCodeGenerationByID(id string) (CodeGenerationOption, error) {
	for _, option := range AvailableCodeGenerationOptions {
		if option.ID == id {
			return option, nil
		}
	}
	return CodeGenerationOption{}, fmt.Errorf("language not found")
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
		IsEnabled:      true,
		IsExperimental: false,
		IsDefault:      false,
		IsPremium:      false,
	},
	{
		Name:           "GPT 4",
		Model:          AgentModelGPT4,
		IsEnabled:      true,
		IsExperimental: false,
		IsDefault:      false,
		IsPremium:      true,
	},
	{
		Name:           "Claude Sonnet 3.5",
		Model:          AgentModelClaudeSonnet3Dot5,
		IsEnabled:      true,
		IsExperimental: false,
		IsDefault:      false,
		IsPremium:      true,
	},
	{
		Name:           "Claude Sonnet 3.7",
		Model:          AgentModelClaudeSonnet3Dot7,
		IsEnabled:      true,
		IsExperimental: false,
		IsDefault:      false,
		IsPremium:      true,
	},
	{
		Name:           "DeepSeek R1",
		Model:          AgentModelDeepSeekR1,
		IsEnabled:      true,
		IsExperimental: true,
		IsDefault:      false,
		IsPremium:      true,
	},
	{
		Name:           "Gemini 1.5 Flash",
		Model:          AgentModelGeminiFlash1Dot5,
		IsEnabled:      false,
		IsExperimental: true,
		IsDefault:      false,
		IsPremium:      false,
	},
	{
		Name:           "Gemini 2.0 Flash",
		Model:          AgentModelGeminiFlash2Dot0,
		IsEnabled:      true,
		IsExperimental: true,
		IsDefault:      true,
		IsPremium:      false,
	},
	{
		Name:           "Grok 2.0",
		Model:          AgentModelGrok2Dot0,
		IsEnabled:      true,
		IsExperimental: true,
		IsDefault:      false,
		IsPremium:      false,
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

func GetModelCatalog(model string) (ModeCatalog, error) {
	for _, catalog := range AvailableCatalogs {
		if catalog.Model == AgentModel(model) {
			return catalog, nil
		}
	}
	return ModeCatalog{}, fmt.Errorf("model not found")
}

type Config struct {
	Model        AgentModel
	OpenAIKey    string
	DeepSeekKey  string
	AnthropicKey string
	GeminiKey    string
	XAIKey       string
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
