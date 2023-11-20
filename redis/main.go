package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	client *redis.Client
}

func Connect(index int) *DB {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  index,  // use default DB
    })
    return &DB{
        client: client,
    }
}

func (db *DB) Size(index int) int {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    return int(db.client.DBSize(ctx).Val())
}

func (db *DB) Set(key string, value string) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    db.client.Set(ctx, key, value, 0)
}

func (db *DB) Get(key string) string {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    return db.client.Get(ctx, key).String()
}