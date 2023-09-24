package nats

type Config struct {
	ServerID string `config:"SERVER_NAME" yaml:"serverID"`
	ClientID string `config:"CLIENT_NAME" yaml:"clientID"`
}
