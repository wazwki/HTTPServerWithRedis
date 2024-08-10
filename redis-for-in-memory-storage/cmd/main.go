package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Set(ctx, "check", 0, time.Millisecond*2).Err(); err != nil {
		log.Fatal(err)
	}

	val, err := rdb.Get(ctx, "exampleKey").Result()
	if err != nil {
		panic(err)
	}
	log.Println("exampleKey:", val)

}
