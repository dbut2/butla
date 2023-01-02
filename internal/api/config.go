package api

import (
	"github.com/dbut2/butla/pkg/database"
	"github.com/dbut2/butla/pkg/redis"
)

type Config struct {
	Address  string          `yaml:"address"`
	Database database.Config `yaml:"database"`
	Redis    redis.Config    `yaml:"redis"`
}
