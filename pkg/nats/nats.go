package nats

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/ziggsdil/zero-level-wb/pkg/models"
)

type Nats struct {
	Conn stan.Conn
}

func NewNatsConnection(cfg Config) (stan.Conn, error) {
	nc, err := stan.Connect(cfg.ServerID, cfg.ClientID)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

//	func (n *Nats) SubscribeToChannel(channelName string, handler func(m *stan.Msg)) {
//		sub, err := n.Conn.Subscribe(channelName, handler)
//		if err != nil {
//			log.Fatalf("failed to subscribe to channel: %s\n", err.Error())
//
//		}
//	}
func IsValid(data []byte) bool {
	var jsonData *models.Message

	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("Received an invalid JSON message: %s\n", err.Error())
		return false
	}
	return true
}
