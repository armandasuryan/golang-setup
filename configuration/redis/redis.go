package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func RedisConnect(host string, password string, db int) *redis.Client {
	fmt.Println(host)
	client := redis.NewClient(&redis.Options{
		// Addr:     "redstag.telkomsigma.co.id?directConnection=true",
		// Password: "p@ssworD123",
		Addr:     host,
		Password: "",
		DB:       0,
	})
	ctx := context.TODO()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("error redis:", err)

	} else {
		fmt.Println("Redis Connect")
	}
	return client

}
