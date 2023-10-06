package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	serviceCache "github.com/patrickmn/go-cache"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/models"
	"log"
)

func RunNatsService(db *db.Database, ctx context.Context, cfg Config, c *serviceCache.Cache) error {
	nc, err := stan.Connect(cfg.ServerID, cfg.ClientID)
	if err != nil {
		fmt.Printf("Failed to connect to nats: %s\n", err.Error())
	}
	defer nc.Close()

	_, err = nc.Subscribe("foo", func(msg *stan.Msg) {
		if err = messageHandler(ctx, db, msg.Data, c); err != nil {
			return
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		fmt.Printf("Failed to subscribe on channel: %s\n", err.Error())
		return err
	}

	select {}
}

func messageHandler(ctx context.Context, db *db.Database, data []byte, c *serviceCache.Cache) error {
	var jsonData models.Message

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Printf("error decoding json response: %v\n", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d\n", e.Offset)
		}
		log.Printf("json response: %q\n", jsonData)
		return err
	}

	if err := db.InsertData(ctx, jsonData.OrderUID, data); err != nil {
		fmt.Printf("Failed to insert data: %s\n", err.Error())
		return err
	}

	bytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Printf("Failed to marshall structure: %s\n", bytes)
		return err
	}

	c.SetDefault(jsonData.OrderUID, bytes)
	return nil
}
