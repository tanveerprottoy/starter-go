package redispkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tanveerprottoy/starter-go/pkg/config"
)

type Client struct {
	client *redis.Client
}

func (c *Client) init() {
	c.client = redis.NewClient(&redis.Options{
		Addr:     config.GetEnvValue("REDIS_URL"),
		Password: config.GetEnvValue("REDIS_PASS"), // "" empty for no password set
		DB:       0,                                // use default DB
	})
	ctx := context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
