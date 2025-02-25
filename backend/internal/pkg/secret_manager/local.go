package secretmanager

import (
	"context"

	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
	"github.com/mujhtech/b0/internal/redis"
	redV9 "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type localClient struct {
	aesCfb encrypt.Encrypt
	cache  cache.Cache
	redis  redV9.UniversalClient
}

func NewLocalClient(cfg *config.Config, redis *redis.Redis, cache cache.Cache) (SecretManager, error) {

	aesCfb, err := encrypt.NewAesCfb(cfg.EncryptionKey)

	if err != nil {
		return nil, err
	}

	client := redis.Client()

	return &localClient{
		aesCfb: aesCfb,
		cache:  cache,
		redis:  client,
	}, nil
}

func (l *localClient) GetSecret(ctx context.Context, secretName string) ([]byte, error) {

	//err := l.cache.Get(ctx, secretName, &secretValue)

	b, err := l.redis.Get(ctx, secretName).Bytes()

	if err != nil {

		if err == redV9.Nil {
			return nil, nil
		}

		return nil, err
	}

	zerolog.Ctx(ctx).Info().Msgf("secretValue: %s", string(b))

	value, err := l.aesCfb.Decrypt(string(b))

	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

func (l *localClient) SetSecret(ctx context.Context, secretName string, secretValue []byte) error {

	if err := l.cache.Delete(ctx, secretName); err != nil {
		zerolog.Ctx(ctx).Err(err).Msgf("failed to delete secret")
	}

	zerolog.Ctx(ctx).Info().Msgf("secretName: %s", secretName)

	encryptedValue, err := l.aesCfb.Encrypt(secretValue)

	if err != nil {
		return err
	}

	zerolog.Ctx(ctx).Info().Msgf("encryptedValue: %s", encryptedValue)

	//return l.cache.Set(ctx, secretName, encryptedValue, 0)
	return l.redis.Set(ctx, secretName, encryptedValue, 0).Err()
}
