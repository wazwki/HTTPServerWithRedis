package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"rediscache/internal/types"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx context.Context
)

func Conn() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func CheckKey(key int) bool {
	fmt.Print(key)
	exists, err := rdb.Exists(ctx, string(key)).Result()
	if err != nil {
		panic(err)
	}
	if exists == 1 {
		return true
	}
	return false
}

func AddKey(user types.User) {
	data, _ := json.Marshal(&user)
	err := rdb.Set(ctx, string(user.ID), string(data), 0).Err()
	if err != nil {
		panic(err)
	}
}

func DeleteKey(key int) {
	err := rdb.Del(ctx, string(key)).Err()
	if err != nil {
		panic(err)
	}
}

func GetKey(key int) types.User {
	var user types.User
	data, err := rdb.Get(ctx, string(key)).Result()
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(data), &user)
	return user
}
