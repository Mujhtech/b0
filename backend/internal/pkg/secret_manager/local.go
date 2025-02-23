package secretmanager

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
)

type localClient struct {
	aesCfb encrypt.Encrypt
	cache  cache.Cache
}

func NewLocalClient(cfg *config.Config, cache cache.Cache) (SecretManager, error) {

	aesCfb, err := encrypt.NewAesCfb(cfg.EncryptionKey)

	if err != nil {
		return nil, err
	}

	return &localClient{
		aesCfb: aesCfb,
		cache:  cache,
	}, nil
}

func (l *localClient) GetSecret(ctx context.Context, secretName string) ([]byte, error) {

	var secretValue []byte

	err := l.cache.Get(ctx, secretName, &secretValue)
	if err != nil {
		return nil, err
	}

	value, err := l.aesCfb.Decrypt(string(secretValue))

	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *localClient) SetSecret(ctx context.Context, secretName string, secretValue []byte) error {

	encryptedValue, err := l.aesCfb.Encrypt(secretValue)

	if err != nil {
		return err
	}

	return l.cache.Set(ctx, secretName, []byte(encryptedValue), 0)
}
