package db

type Config struct {
	Database string `config:"POSTGRES_DB" yaml:"database"`
	User     string `config:"POSTGRES_USER" yaml:"user"`
	Password string `config:"POSTGRES_PASSWORD" yaml:"password"`
	Host     string `config:"POSTGRES_HOST" yaml:"host"`
	Port     int    `config:"POSTGRES_PORT" yaml:"port"`
}
