package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/models"
)

func RunNatsService(db *db.Database, ctx context.Context, cfg Config) error {
	nc, err := stan.Connect(cfg.ServerID, cfg.ClientID)
	if err != nil {
		fmt.Printf("Failed to connect to nats: %s\n", err.Error())
	}
	defer nc.Close()

	_, err = nc.Subscribe("foo", func(msg *stan.Msg) {
		if err = messageHandler(ctx, db, msg.Data); err != nil {
			return
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Printf("Failed to subscribe on channel: %s\n", err.Error())
		return err
	}

	select {}
}

func messageHandler(ctx context.Context, db *db.Database, data []byte) error {
	var jsonData *models.Message

	// validate is json
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Printf("Failed to unmarshall data: %s\n", err.Error())
		return err
	}

	if err := db.InsertData(ctx, jsonData.OrderUID, data); err != nil {
		fmt.Printf("Failed to insert data: %s\n", err.Error())
		return err
	}
	// после того как успешно внеслось в бд, то я должен добавить это в кэш

	return nil
}
