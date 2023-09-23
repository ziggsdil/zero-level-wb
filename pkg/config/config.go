package config

import "github.com/ziggsdil/zero-level-wb/pkg/db"

type Config struct {
	Host string `config:"APP_HOST" yaml:"host"`
	Port string `config:"APP_PORT" yaml:"port"`

	Postgres db.Config `config:"postgres"`
}
