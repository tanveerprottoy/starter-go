package redispkg

import "github.com/redis/go-redis/v9"

type Client struct {

}

func (c *Client) init() {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }
}