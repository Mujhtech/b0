package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

const (
	ProjectName = "b0"
)

var DefaultConfig = &Config{
	Cache: Cache{
		Provider: CacheProviderRedis,
	},
	Cors: Cors{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Authorization", "Accept-Encoding", "Content-Length", "Content-Type", "X-CSRF-Token", "X-Requested-With", "X-Requested-Id", "x-app-id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	},
	Database: Database{
		Driver:   DatabaseDriverPostgres,
		Host:     "localhost",
		Port:     5432,
		User:     ProjectName,
		Password: ProjectName,
		Database: ProjectName,
		Options:  "sslmode=disable&connect_timeout=30",
	},
	Redis: Redis{
		Host:               "localhost",
		Port:               6379,
		Username:           "",
		Password:           "",
		MinIdleConnections: 0,
		MaxRetries:         3,
		DB:                 1,
	},
	Aws: Aws{
		DefaultRegion: "eu-west-2",
	},
	Server: Server{
		Port: 5555,
		SSL:  false,
	},
	Auth: Auth{
		RedirectUrl:   "http://localhost:5555",
		UIRedirectUrl: "http://localhost:3000/auth/callback",
	},
	Email: Email{
		FromAddress: "b0 <no-reply@b0.dev>",
	},
	Job: Job{
		Concurrency: 10,
	},
	Pubsub: Pubsub{
		Provider:       PubsubProviderInMemory,
		App:            ProjectName,
		Namespace:      ProjectName,
		HealthInterval: 2,
		SendTimeout:    60,
		ChannelSize:    500,
	},
}

func LoadConfig() (*Config, error) {
	config := DefaultConfig

	// Override config from environment variables
	err := envconfig.Process("", config)
	if err != nil {
		return nil, err
	}

	if err = config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) validate() error {
	// Validate database configuration
	if c.Database.Host == "" {
		return fmt.Errorf("database host cannot be empty")
	}
	if c.Database.Port == 0 {
		return fmt.Errorf("database port cannot be zero")
	}

	dbDsn := c.Database.BuildDsn()
	if dbDsn == "" {
		return fmt.Errorf("database dsn is empty")
	}

	return nil
}
