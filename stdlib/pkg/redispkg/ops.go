package redispkg

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Set(key string, val any, exp time.Duration, ctx context.Context, c *redis.Client) (string, error) {
	sts := c.Set(ctx, key, val, exp)
	return sts.Result()
}

func Get(key string, val any, exp time.Duration, ctx context.Context, c *redis.Client) (string, error) {
	sts := c.Get(ctx, key)
	return sts.Result()
}
