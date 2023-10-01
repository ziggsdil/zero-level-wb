package cache

import (
	"context"
	serviceCache "github.com/patrickmn/go-cache"
	"github.com/ziggsdil/zero-level-wb/pkg/db"
	"sync"
	"time"
)

func New(ctx context.Context, db *db.Database, wg *sync.WaitGroup) (*serviceCache.Cache, error) {
	var err error
	c := serviceCache.New(time.Minute*30, time.Hour*24)

	wg.Add(1)
	go func(c *serviceCache.Cache, err error) {
		defer wg.Done()
		data, err := db.GetAllData(ctx)
		if err != nil {
			return
		}

		for key, val := range data {
			c.SetDefault(key, val)
		}
	}(c, err)

	return c, nil
}
