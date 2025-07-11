package config

import (
	"fmt"
	"net/url"
	"time"
)

type DatabaseDriver string
type PubsubProvider string
type CacheProvider string
type TelemetryProvider string
type SecretManagerProvider string

const (
	DatabaseDriverPostgres DatabaseDriver = "postgres"
	DatabaseDriverSqlite3  DatabaseDriver = "sqlite3"

	DefaultConfigFilePath string = ".env"

	PubsubProviderAwsSqs   PubsubProvider = "aws_sqs"
	PubsubProviderRedis    PubsubProvider = "redis"
	PubsubProviderGoogle   PubsubProvider = "google"
	PubsubProviderKafka    PubsubProvider = "kafka"
	PubsubProviderAmqp     PubsubProvider = "amqp"
	PubsubProviderInMemory PubsubProvider = "inmemory"

	CacheProviderRedis    CacheProvider = "redis"
	CacheProviderInMemory CacheProvider = "inmemory"

	TelemetryProviderSentry TelemetryProvider = "sentry"
	TelemetryProviderOtel   TelemetryProvider = "otel"
	TelemetryProviderNone   TelemetryProvider = "none"

	SecretManagerProviderAws    SecretManagerProvider = "aws"
	SecretManagerProviderGoogle SecretManagerProvider = "google"
	SecretManagerProviderLocal  SecretManagerProvider = "local"
	SecretManagerProviderNone   SecretManagerProvider = "none"
)

type Config struct {
	EncryptionKey string        `json:"encryption_key" envconfig:"ENCRYPTION_KEY"`
	Database      Database      `json:"database"`
	Redis         Redis         `json:"redis"`
	Aws           Aws           `json:"aws"`
	Server        Server        `json:"server"`
	Auth          Auth          `json:"auth"`
	Email         Email         `json:"email"`
	Job           Job           `json:"job"`
	Pubsub        Pubsub        `json:"pubsub"`
	Cors          Cors          `json:"cors"`
	Cache         Cache         `json:"cache"`
	Telemetry     Telemetry     `json:"telemetry"`
	Agent         Agent         `json:"agent"`
	Integrations  Integrations  `json:"integrations"`
	Stripe        Stripe        `json:"stripe"`
	SecretManager SecretManager `json:"secret_manager"`
}

type SecretManager struct {
	Provider SecretManagerProvider `json:"provider" envconfig:"SECRET_MANAGER_PROVIDER"`
}

type Stripe struct {
	ApiKey                     string `json:"api_key" envconfig:"STRIPE_API_KEY"`
	StarterSubscriptionPriceID string `json:"starter_subscription_price_id" envconfig:"STRIPE_STARTER_SUBSCRIPTION_PRICE_ID"`
	ProSubscriptionPriceID     string `json:"pro_subscription_price_id" envconfig:"STRIPE_PRO_SUBSCRIPTION_PRICE_ID"`
	ScaleSubscriptionPriceID   string `json:"scale_subscription_price_id" envconfig:"STRIPE_SCALE_SUBSCRIPTION_PRICE_ID"`
}

type Integrations struct {
	Discord  Discord  `json:"discord"`
	Telegram Telegram `json:"telegram"`
}

type Telegram struct {
	TelegramBotToken string `json:"telegram_bot_token" envconfig:"TELEGRAM_BOT_TOKEN"`
	ChatID           string `json:"channel_id" envconfig:"TELEGRAM_CHAT_ID"`
}

type Discord struct {
	DiscordBotToken string `json:"discord_bot_token" envconfig:"DISCORD_BOT_TOKEN"`
	ChannelID       string `json:"channel_id" envconfig:"DISCORD_CHANNEL_ID"`
}

type Telemetry struct {
	Provider TelemetryProvider `json:"provider" envconfig:"TELEMETRY_PROVIDER"`
	Sentry   Sentry            `json:"sentry"`
	OTEL     OTEL              `json:"otel"`
}

type Agent struct {
	OpenAIKey    string `json:"openai_key" envconfig:"AGENT_OPENAI_KEY"`
	DeepSeekKey  string `json:"deepseek_key" envconfig:"AGENT_DEEPSEEK_KEY"`
	AnthropicKey string `json:"anthropic_key" envconfig:"AGENT_ANTHROPIC_KEY"`
	GeminiKey    string `json:"gemini_key" envconfig:"AGENT_GEMINI_KEY"`
	XAIKey       string `json:"xai_key" envconfig:"AGENT_XAI_KEY"`
}

type Sentry struct {
	DSN string `json:"dsn" envconfig:"SENTRY_DSN"`
}

type OTEL struct {
	CollectorURL       string                `json:"collector_url" envconfig:"OTEL_COLLECTOR_URL"`
	InsecureSkipVerify bool                  `json:"insecure_skip_verify" envconfig:"OTEL_INSECURE_SKIP_VERIFY"`
	SampleRate         float64               `json:"sample_rate" envconfig:"OTEL_SAMPLE_RATE"`
	OTelAuth           OTelAuthConfiguration `json:"otel_auth"`
}

type OTelAuthConfiguration struct {
	HeaderName  string `json:"header_name" envconfig:"OTEL_HEADER_NAME"`
	HeaderValue string `json:"header_value" envconfig:"OTEL_HEADER_VALUE"`
}

type Cache struct {
	Provider CacheProvider `json:"provider" envconfig:"CACHE_PROVIDER"`
}

// Cors defines CORS configuration
type Cors struct {
	AllowedOrigins   []string `json:"allowed_origins" envconfig:"CORS_ALLOWED_ORIGINS"`
	AllowedMethods   []string `json:"allowed_methods" envconfig:"CORS_ALLOWED_METHODS"`
	AllowedHeaders   []string `json:"allowed_headers" envconfig:"CORS_ALLOWED_HEADERS"`
	ExposedHeaders   []string `json:"exposed_headers" envconfig:"CORS_EXPOSED_HEADERS"`
	AllowCredentials bool     `json:"allow_credentials" envconfig:"CORS_ALLOW_CREDENTIALS"`
	MaxAge           int      `json:"max_age" envconfig:"CORS_MAX_AGE"`
}

// Database defines database configuration
type Database struct {
	Driver   DatabaseDriver `json:"driver" envconfig:"DB_DRIVER"`
	Host     string         `json:"host" envconfig:"DB_HOST"`
	Port     int            `json:"port" envconfig:"DB_PORT"`
	User     string         `json:"user" envconfig:"DB_USER"`
	Password string         `json:"password" envconfig:"DB_PASSWORD"`
	Database string         `json:"database" envconfig:"DB_DATABASE"`
	Options  string         `json:"options" envconfig:"DB_OPTIONS"`
}

// Aws defines AWS configuration
type Aws struct {
	DefaultRegion string `json:"region" envconfig:"AWS_DEFAULT_REGION"`
	AccessKey     string `json:"access_key" envconfig:"AWS_ACCESS_KEY"`
	SecretKey     string `json:"secret_key" envconfig:"AWS_SECRET_KEY"`
}

// Server defines server configuration
type Server struct {
	Port        uint32 `json:"port" envconfig:"PORT"`
	SSL         bool   `json:"ssl" envconfig:"SSL"`
	SSLCertFile string `json:"ssl_cert_file" envconfig:"SSL_CERT_FILE"`
	SSLKeyFile  string `json:"ssl_key_file" envconfig:"SSL_KEY_FILE"`
	Timeout     uint32 `json:"timeout" envconfig:"TIMEOUT"`
}

// Redis defines redis configuration
type Redis struct {
	Host               string `json:"host" envconfig:"REDIS_HOST"`
	Port               int    `json:"port" envconfig:"REDIS_PORT"`
	Username           string `json:"username" envconfig:"REDIS_USERNAME"`
	Password           string `json:"password" envconfig:"REDIS_PASSWORD"`
	MaxRetries         int    `json:"max_retries" envconfig:"REDIS_MAX_RETRIES"`
	MinIdleConnections int    `json:"min_idle_connections" envconfig:"REDIS_MIN_IDLE_CONNECTIONS"`
	DB                 int    `json:"db" envconfig:"REDIS_DB"`
}

type Auth struct {
	RedirectUrl   string     `json:"redirect_url" envconfig:"AUTH_REDIRECT_URL"`
	UIRedirectUrl string     `json:"ui_redirect_url" envconfig:"AUTH_UI_REDIRECT_URL"`
	GoogleAuth    GoogleAuth `json:"google_auth"`
	GithubAuth    GithubAuth `json:"github_auth"`
}

type GoogleAuth struct {
	ClientID     string `json:"auth_google_client_id" envconfig:"AUTH_GOOGLE_CLIENT_ID"`
	ClientSecret string `json:"auth_google_client_secret" envconfig:"AUTH_GOOGLE_CLIENT_SECRET"`
}

type GithubAuth struct {
	ClientID     string `json:"auth_github_client_id" envconfig:"AUTH_GITHUB_CLIENT_ID"`
	ClientSecret string `json:"auth_github_client_secret" envconfig:"AUTH_GITHUB_CLIENT_SECRET"`
}

type Email struct {
	FromAddress string `json:"from_address" envconfig:"EMAIL_FROM_ADDRESS"`
}

type Job struct {
	Concurrency int `json:"concurrency" envconfig:"JOB_CONCURRENCY"`
}

type Pubsub struct {
	App            string         `json:"app" envconfig:"PUBSUB_APP"`
	Namespace      string         `json:"namespace" envconfig:"PUBSUB_NAMESPACE"`
	Provider       PubsubProvider `json:"provider" envconfig:"PUBSUB_PROVIDER"`
	SendTimeout    time.Duration  `json:"send_timeout" envconfig:"PUBSUB_SEND_TIMEOUT"`
	ChannelSize    int            `json:"channel_size" envconfig:"PUBSUB_CHANNEL_SIZE"`
	HealthInterval time.Duration  `json:"health_interval" envconfig:"PUBSUB_HEALTH_INTERVAL"`
	Google         GooglePubsub   `json:"google"`
	Amqp           Amqp           `json:"amqp"`
	Brokers        []string       `json:"brokers" envconfig:"PUBSUB_BROKERS"`
}

type GooglePubsub struct {
	ProjectID string `json:"project_id" envconfig:"PUBSUB_GOOGLE_PROJECT_ID"`
}

type Amqp struct {
	Url string `json:"url" envconfig:"PUBSUB_AMQP_URL"`
}

func (d *Database) BuildDsn() string {
	if d.Driver == "" {
		return ""
	}

	authPart := ""
	if d.User != "" || d.Password != "" {
		authPrefix := url.UserPassword(d.User, d.Password)
		authPart = fmt.Sprintf("%s@", authPrefix)
	}

	dbPart := ""
	if d.Database != "" {
		dbPart = fmt.Sprintf("/%s", d.Database)
	}

	optPart := ""
	if d.Options != "" {
		optPart = fmt.Sprintf("?%s", d.Options)
	}

	return fmt.Sprintf("%s://%s%s:%d%s%s", d.Driver, authPart, d.Host, d.Port, dbPart, optPart)
}

func (r *Redis) BuildDsn() string {
	if r.Host == "" {
		return ""
	}

	authPart := ""
	if r.Username != "" || r.Password != "" {
		authPrefix := url.UserPassword(r.Username, r.Password)
		authPart = fmt.Sprintf("%s@", authPrefix)
	}

	return fmt.Sprintf("redis://%s%s:%d", authPart, r.Host, r.Port)
}
