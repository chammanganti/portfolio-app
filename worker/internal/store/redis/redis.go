package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Redis interface
type RedisInterface interface {
	createClient() *redis.Client
	Get(key string) (string, error)
	Set(key string, value interface{}, exp time.Duration) error
	Del(keys ...string) error
	DelPattern(key string) error
}

// Redis
type Redis struct {
	Context  context.Context
	Addr     string
	Password string
	DB       int
}

// New redis
func NewRedis(context context.Context, addr, password string, db int) RedisInterface {
	return &Redis{
		Context:  context,
		Addr:     addr,
		Password: password,
		DB:       db,
	}
}

// Create a redis client
func (r Redis) createClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
}

// Gets data
func (r Redis) Get(key string) (string, error) {
	client := r.createClient()
	defer client.Close()

	return client.Get(r.Context, key).Result()
}

// Sets data
func (r Redis) Set(key string, val interface{}, exp time.Duration) error {
	client := r.createClient()
	defer client.Close()

	return client.Set(r.Context, key, val, exp).Err()
}

// Deletes data
func (r Redis) Del(keys ...string) error {
	client := r.createClient()
	defer client.Close()

	return client.Del(r.Context, keys...).Err()
}

// Deletes data using pattern
func (r Redis) DelPattern(pattern string) error {
	client := r.createClient()
	defer client.Close()

	keys := client.Scan(r.Context, 0, pattern, 0).Iterator()
	for keys.Next(r.Context) {
		client.Del(r.Context, keys.Val())
	}

	return keys.Err()
}
