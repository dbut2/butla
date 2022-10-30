package config

import (
	"context"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"gopkg.in/yaml.v3"
)

type Secret string

var _ loader = new(Secret)

func (g Secret) load(c any) error {
	manager, err := sm.NewClient(context.Background())
	if err != nil {
		return err
	}

	secret, err := manager.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: string(g),
	})
	if err != nil {
		return err
	}

	return yaml.Unmarshal(secret.GetPayload().GetData(), c)
}
