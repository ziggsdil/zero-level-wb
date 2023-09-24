package config

import (
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/nats"
)

type Config struct {
	Host string `config:"APP_HOST" yaml:"host"`
	Port string `config:"APP_PORT" yaml:"port"`

	Postgres db.Config   `config:"postgres"`
	Nats     nats.Config `config:"nats"`
}
