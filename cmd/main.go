package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/ziggsdil/zero-level-wb/pkg/config"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/handler"
	"github.com/ziggsdil/zero-level-wb/pkg/nats"
	"log"
	"sync"
)

func main() {
	ctx := context.Background()

	var cfg config.Config

	var wg *sync.WaitGroup
	defer wg.Wait()

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

	InitNatsStreaming(cfg, ctx, postgres, wg)

	handlers := handler.NewHandler(postgres) // если мы выдаем из кэша данные, то не нужна база данных

	// todo: handleMessage
}

func InitNatsStreaming(cfg config.Config, ctx context.Context, db *db.Database, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := nats.RunNatsService(db, ctx, cfg.Nats)
		if err != nil {
			log.Fatalf("failed to run nats service: %s\n", err.Error())
			return
		}
	}(wg)
}
