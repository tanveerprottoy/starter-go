package redispkg

import (
	"github.com/redis/go-redis/v9"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"
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
}
