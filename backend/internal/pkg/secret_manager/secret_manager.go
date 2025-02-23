package secretmanager

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
)

type SecretManager interface {
	GetSecret(ctx context.Context, secretName string) ([]byte, error)
	SetSecret(ctx context.Context, secretName string, secretValue []byte) error
}

type secretManager struct{}

func (s *secretManager) GetSecret(ctx context.Context, secretName string) ([]byte, error) {
	return nil, nil
}

func (s *secretManager) SetSecret(ctx context.Context, secretName string, secretValue []byte) error {
	return nil
}

func New(ctx context.Context, cfg *config.Config, cache cache.Cache) (SecretManager, error) {
	switch cfg.SecretManager.Provider {
	case config.SecretManagerProviderGoogle:
		return NewGoogleClient(ctx)
	case config.SecretManagerProviderAws:
		return NewAwsClient(ctx, cfg)
	case config.SecretManagerProviderLocal:
		return NewLocalClient(cfg, cache)
	default:
		return &secretManager{}, nil
	}
}
