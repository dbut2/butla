package secrets

import (
	"context"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gopkg.in/yaml.v3"
)

type GsmResourceID string

func (g GsmResourceID) getGsmResourceID() string {
	return string(g)
}

func LoadSecret(g gsmResourcer) error {
	if g.getGsmResourceID() == "" {
		return nil
	}

	return loadSecret(g)
}

type gsmResourcer interface {
	getGsmResourceID() string
}

func loadSecret(config gsmResourcer) error {
	manager, err := sm.NewClient(context.Background())
	if err != nil {
		return err
	}

	secret, err := manager.AccessSecretVersion(context.Background(), &secretmanager.AccessSecretVersionRequest{
		Name: config.getGsmResourceID(),
	})
	if err != nil {
		return err
	}

	return yaml.Unmarshal(secret.GetPayload().GetData(), config)
}
