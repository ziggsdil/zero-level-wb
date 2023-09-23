package main

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/ziggsdil/zero-level-wb/pkg/config"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
)

func main() {
	ctx := context.Background()
	var cfg config.Config

	err := confita.NewLoader(
		file.NewBackend("./default.yaml"),
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		fmt.Printf("failed to parse config: %s\n", err.Error())
		return
	}

	postgres, dbErr := db.NewDatabase(cfg.Postgres)
	if dbErr != nil {
		fmt.Printf("failed to connect to postgresql: %s\n", dbErr.Error())
		return
	}

}
