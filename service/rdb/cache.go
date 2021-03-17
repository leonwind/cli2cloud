package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

var Cache = redis.NewClient(&redis.Options {
	Addr: "localhost:6379",
	Password: "",
	DB: 0,
})