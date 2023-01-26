package main

import (
	"github.com/dbut2/butla/pkg/stores"
)

type Config struct {
	Address      string         `yaml:"address"`
	Host         host           `yaml:"host"`
	Store        *stores.Config `yaml:"store"`
	LoginEnabled bool           `yaml:"login"`
}

type host struct {
	Scheme   string `yaml:"scheme"`
	Hostname string `yaml:"hostname"`
}
