package main

import (
	"github.com/dbut2/butla/pkg/stores"
)

type Config struct {
	Address string         `yaml:"address"`
	Store   *stores.Config `yaml:"store"`
}
