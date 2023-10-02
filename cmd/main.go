package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	serviceCache "github.com/patrickmn/go-cache"
	"github.com/ziggsdil/zero-level-wb/pkg/cache"
	"github.com/ziggsdil/zero-level-wb/pkg/config"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"github.com/ziggsdil/zero-level-wb/pkg/handler"
	"github.com/ziggsdil/zero-level-wb/pkg/nats"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
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
	c, err := cache.New(ctx, postgres, wg)
	if err != nil {
		fmt.Printf("Failed to run cache: %s\n", err.Error())
		return
	}

	err = InitNatsStreaming(cfg, ctx, postgres, wg, c)
	if err != nil {
		fmt.Printf("Failed to init nats-streaming: %s\n", err.Error())
		return
	}

	handlers := handler.NewHandler(postgres, c) // если мы выдаем из кэша данные, то не нужна база данных
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: handlers.Router(),
	}

	go func() {
		fmt.Println("server started...")
		_ = srv.ListenAndServe()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}

func InitNatsStreaming(cfg config.Config, ctx context.Context, db *db.Database, wg *sync.WaitGroup, c *serviceCache.Cache) error {
	wg.Add(1)
	var err error
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err = nats.RunNatsService(db, ctx, cfg.Nats, c)
		if err != nil {
			log.Fatalf("failed to run nats service: %s\n", err.Error())
			return
		}
	}(wg)
	return err
}
