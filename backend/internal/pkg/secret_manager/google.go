package secretmanager

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type googleClient struct {
	client *secretmanager.Client
}

func NewGoogleClient(ctx context.Context) (SecretManager, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	return &googleClient{
		client: client,
	}, nil
}

func (g *googleClient) GetSecret(ctx context.Context, secretName string) ([]byte, error) {
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	// Call the API.
	result, err := g.client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return nil, err
	}

	return result.Payload.Data, nil
}

func (g *googleClient) SetSecret(ctx context.Context, secretName string, secretValue []byte) error {

	secret, err := g.client.CreateSecret(ctx, &secretmanagerpb.CreateSecretRequest{
		Parent:   "projects/b0",
		SecretId: fmt.Sprintf("projects/b0/%s", secretName),
		Secret: &secretmanagerpb.Secret{
			Name: secretName,
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	})
	if err != nil {
		return nil
	}

	_, err = g.client.AddSecretVersion(ctx, &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: secretValue,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
