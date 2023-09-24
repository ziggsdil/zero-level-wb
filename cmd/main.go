package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/nats-io/stan.go"
	"github.com/ziggsdil/zero-level-wb/pkg/config"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/models"
	"github.com/ziggsdil/zero-level-wb/pkg/nats"
	"log"
)

func main() {
	ctx := context.Background()
	var cfg config.Config

	err := confita.NewLoader(
		file.NewBackend("./deploy/default.yaml"),
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		fmt.Printf("failed to parse config: %s\n", err.Error())
		return
	}

	postgres, err := db.NewDatabase(cfg.Postgres)
	if err != nil {
		fmt.Printf("failed to connect to postgresql: %s\n", err.Error())
		return
	}

	nc, err := nats.NewNatsConnection(cfg.Nats)
	if err != nil {
		fmt.Printf("failed to connect to nats-streaming: %s\n", err.Error())
		return
	}
	defer func() { _ = nc.Close() }()

	// нам пришло сообщение и необходимо проверить есть ли такое уже в бд
	sub, err := nc.Subscribe("foo", func(m *stan.Msg) {
		// todo: call a method to valid
		if nats.IsValid(m.Data) { // если данные не валидны мы должны просто игнорировать их
			// todo: call a method for save in db, если в бд уже будет существовать такой индекс то пропускаем
			// на этапе добавления в бд можно будет проверить существует ли такой индекс
			// todo: в какой момент сохранять кэш, т.е. заполнять мапу?
			var data models.Message
			err := postgres.InsertData(ctx, data.OrderUID, m.Data)
			// todo: слишком все залезает друг на друга, гадо исправить логику

		}

		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("failed to subscribe to channel: %s\n", err.Error())
	}
	defer sub.Unsubscribe()

	// for loop listen
	select {}

	// todo: handleMessage
}
