package secretmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/mujhtech/b0/config"
)

type awsClient struct {
	svc *secretsmanager.Client
}

func NewAwsClient(ctx context.Context, cfg *config.Config) (SecretManager, error) {
	config, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.Aws.AccessKey,
				cfg.Aws.SecretKey,
				"",
			),
		),
		awsConfig.WithRegion(cfg.Aws.DefaultRegion),
	)

	if err != nil {
		return nil, err
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	return &awsClient{
		svc: svc,
	}, nil
}

func (a *awsClient) GetSecret(ctx context.Context, secretName string) ([]byte, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := a.svc.GetSecretValue(ctx, input)
	if err != nil {
		return nil, err
	}

	return []byte(*result.SecretString), nil
}

func (a *awsClient) SetSecret(ctx context.Context, secretName string, secretValue []byte) error {
	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(secretName),
		SecretString: aws.String(string(secretValue)),
	}

	_, err := a.svc.CreateSecret(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
