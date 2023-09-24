package nats

import (
	"github.com/nats-io/stan.go"
)

func NewNatsConnection(cfg Config) (stan.Conn, error) {
	nc, err := stan.Connect(cfg.ServerID, cfg.ClientID)
	if err != nil {
		return nil, err
	}
	return nc, err
}
