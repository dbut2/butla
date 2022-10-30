package config

import (
	"context"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

type Secret string

func (g Secret) load() ([]byte, error) {
	manager, err := sm.NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	secret, err := manager.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: string(g),
	})
	if err != nil {
		return nil, err
	}

	return secret.GetPayload().GetData(), nil
}
