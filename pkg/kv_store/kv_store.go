package kvstore

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

var (
	ErrKeyValueNotExists = errors.New("key value does not exits")
)

type KVStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string) error
	Truncate() error
}
type RedisKVStore struct {
	client *redis.Client
}

func New() *RedisKVStore {
	return &RedisKVStore{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
func (i *RedisKVStore) Get(key string) string {
	value := i.client.Get(key)
	return value.String()
}

func (i *RedisKVStore) Set(key, value string) error {
	i.client.Set(key, value, 24*time.Hour)
	return nil
}

func (i *RedisKVStore) Delete(key string) {
	i.client.Del(key)
}
